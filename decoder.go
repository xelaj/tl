// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

func Unmarshal(b []byte, res any) error {
	return NewDecoder(bytes.NewBuffer(b)).Decode(res)
}

type Decoder interface {
	SetRegistry(registry *ObjectRegistry) Decoder

	Decode(res any) error
}

type UnmarshalState interface {
	// ReadWords reads words from unmarshaler with fixed size of word.
	io.Reader
	Peek(seek, size int) ([]byte, error)
	SkipBytes(int)

	PopBool() (bool, error)
	PopCRC() (crc32, error)
}

// A decoder reads and decodes TL values from an input stream.
type decoder struct { //nolint:govet // false positive for fieldalignment
	r bufio.Reader

	registry  *ObjectRegistry
	endianess binary.ByteOrder
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) Decoder { return NewDecoderWithSize(r, 0) }

// NewDecoderWithSize works absolutely like NewDecoder, but it sets buffer with
// size that you want. It could be useful, when you want to debug serialized
// data. You can combine this constructor with DumpWithoutRead method, which
// returns all data, until cache of bufio will be full, or until reader reach
// end of file.
//
// To make correct cache, in most of cases, you will get length of all
// serialized message. pass this length and few more bytes (50-60 bytes will be
// enough). Note that DumpWithoutRead can't guarantee to be stable, so use it
// only for debugging.
func NewDecoderWithSize(r io.Reader, bufSize int) Decoder {
	// bufio doesn't export this constant, so, we must set it manually
	const bufioDefaultBufSize = 4096

	if bufSize <= 0 {
		bufSize = bufioDefaultBufSize
	}

	return &decoder{
		r:         *bufio.NewReaderSize(r, bufSize),
		registry:  defaultRegistry,
		endianess: binary.LittleEndian,
	}
}

func (d *decoder) SetRegistry(registry *ObjectRegistry) Decoder {
	d.registry = registry

	return d
}

func (d *decoder) Decode(res any) error {
	if res == nil {
		return ErrUnexpectedNil
	}

	v := reflect.ValueOf(res)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("res value is not pointer as expected. got %v", v.Type())
	}

	// decoding elem cause we are taking pointer in res, but we'll take
	// additional step to call decodeValue again. Who needs that?
	return d.decodeValue(v.Elem())
}

// DecodeUnknown works like Decode, but it tries to get object stored in data stream.
func (d *decoder) DecodeUnknown() (any, error) {
	crc, err := d.PopCRC()
	if err != nil {
		return nil, errReadCRC{err}
	}

	switch crc {
	case crcVector:
		// TODO: last time problem was exact in this case, we couldn't parse
		//       unknown vector from telegram responses. So, we added cache for
		//       types "prediction", which parser used before. If we can't
		//       understand which type written, we can't parse **any** received
		//       message from server.
		//
		//       it's really important problem for typelang: dynamic messages
		//       are not so dynamic as telegram devs tell to you.
		// hashtag for quick search in code: #explanation_of_vector_problem
		return nil, errors.New("got vector, not allowed to decode it manually")
	case crcFalse:
		return false, nil
	case crcTrue:
		return true, nil
	case crcNull:
		return nil, nil
	}

	enum, ok := d.registry.enums[crc]
	if ok {
		return enum, nil
	}

	object, err := d.registry.spawnObject(crc)
	if err != nil {
		return nil, err
	}

	err = d.decodeObject(object, true)
	if err != nil {
		return nil, err
	}

	return object.Interface(), nil
}

func (d *decoder) peekPadding(seek, msgSize int) (err error) {
	if p := pad(msgSize, WordLen); p > 0 {
		_, err = d.Peek(seek, p)
	}

	return err
}

func (d *decoder) Peek(seek, size int) ([]byte, error) {
	peeked, err := d.r.Peek(seek + size)
	if err != nil {
		return nil, err //nolint:wrapcheck // must not be wrapped
	}

	return peeked[seek:], nil
}

func (d *decoder) SkipBytes(n int) {
	// discarded will be only exact number or less, BUT with returned error
	if _, err := d.r.Discard(n); err != nil {
		panic(errors.Wrap(err, "something wrong with peeking bytes, Discard must always be ok"))
	}
}

func (d *decoder) decodeValue(value reflect.Value) error {
	if unmarshaler, ok := value.Interface().(Unmarshaler); ok {
		return unmarshaler.UnmarshalTL(d) //nolint:wrapcheck // makes no sense
	}

	// extra case
	if value.Type().Implements(enumTyp) {
		crc, err := d.PopCRC()
		if err != nil {
			return err
		}

		return d.decodeEnum(value, crc)
	}

	var val any
	var err error

	switch kind := value.Kind(); kind { //nolint:exhaustive
	// pointer always diving into value
	case reflect.Ptr:
		// need to init firstly
		if value.IsZero() {
			value.Set(reflect.New(value.Type().Elem()))
		}

		return d.decodeValue(value.Elem())

	// simple types
	case reflect.Float64:
		val, err = d.popDouble()

	case reflect.Int64:
		val, err = d.popLong()

	case reflect.Uint32:
		val, err = d.PopUint()

	case reflect.Int32:
		val, err = d.popInt()

	case reflect.Bool:
		val, err = d.PopBool()

	case reflect.String:
		val, err = d.popString()

	case reflect.Slice:
		switch value.Type() {
		case byteSliceTyp: // []byte
			val, err = d.popMessage()

		default:
			return d.decodeVector(value, false)
		}

	case reflect.Array:
		return d.decodeVector(value, false)

	// complex types
	case reflect.Struct:
		return d.decodeObject(value, false)

	// struct but it's map. Maps are not defined by TL, so we use them as type
	// values.
	case reflect.Map:
		// TODO: must implement map decoding
		return errors.New("map is not ordered object (must order like structs)")

	// interfaces in terms of TL is a set of allowed structs.
	case reflect.Interface:
		return d.decodeInterface(value)

	default:
		// unsupported at all
		// Chan, Func, Uintptr, UnsafePointer
		//
		// supported, but TL doesn't support 8 and 16 bit numbers
		// Int, Int8, Int16, Uint, Uint8, Uint16, Uint64
		//
		// same: supported, but not in TL, so we can't understand, how much bytes
		// we need to scan.
		// Float32, Complex64, Complex128
		return ErrUnsupportedType{Type: value.Type()}
	}

	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(val))

	return nil
}

// allowed values are only slice and array.
func (d *decoder) decodeVector(v reflect.Value, ignoreCRC bool) error {
	if !ignoreCRC {
		crc, err := d.PopCRC()
		if err != nil {
			return errReadCRC{err}
		}

		if crc != crcVector {
			return fmt.Errorf("not a vector: 0x%08x, want: 0x%08x", crc, crcVector)
		}
	}

	size, err := d.PopUint()
	if err != nil {
		return errors.Wrap(err, "read vector size")
	}

	switch v.Kind() { //nolint:exhaustive // default unreachable
	case reflect.Array:
		if v.Len() < int(size) {
			return fmt.Errorf("array size is smaller than message: got %v, want %v", v.Len(), size)
		}
	case reflect.Slice:
		v.Set(reflect.MakeSlice(v.Type(), int(size), int(size)))

	default:
		unreachable()
	}

	for i := 0; i < int(size); i++ {
		if err := d.decodeValue(v.Index(i)); err != nil {
			return wrapPath(err, "["+strconv.Itoa(i)+"]")
		}
	}

	return nil
}

func (d *decoder) decodeInterface(v reflect.Value) error {
	crc, err := d.PopCRC()
	if err != nil {
		return errReadCRC{err}
	}

	o, err := d.registry.spawnObject(crc)
	if err != nil {
		return err
	}

	err = d.decodeObject(o, true)
	if err != nil {
		return err // no need to wrap
	}
	v.Set(o)

	return nil
}

func (d *decoder) decodeEnum(v reflect.Value, crc crc32) error {
	enum, ok := d.registry.enums[crc]
	if !ok {
		return fmt.Errorf("enum 0x%08x not found", crc)
	}
	value := reflect.ValueOf(enum)
	if v.Type() != value.Type() {
		return fmt.Errorf("invalid type of enum: want %v, got %v", v.Type(), value.Type())
	}

	v.Set(value)

	return nil
}

// decode EXACT object. means that v must always be struct.
func (d *decoder) decodeObject(v reflect.Value, ignoreCRC bool) error {
	if v.Kind() == reflect.Interface && ignoreCRC {
		return errors.New("can't decode value to interface without reading crc code")
	}

	crc, ok := getCRCofObject(v)
	if !ok {
		return errors.New("value must implement Object interface, got: " + v.Type().String())
	}

	fields, ok := d.registry.structFields[crc]
	if !ok {
		return fmt.Errorf("object 0x%08x is not registered", crc)
	}

	if !ignoreCRC {
		gotCrc, err := d.PopCRC()
		if err != nil {
			return errReadCRC{err}
		}

		if gotCrc != crc {
			return ErrInvalidCRC{Got: gotCrc, Want: crc}
		}
	}

	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct, got %v with %v kind", v.Type(), v.Kind())
	}

	bitflags := map[int]uint32{}

	for i := 0; i < v.NumField(); i++ {
		tags := fields.tags[i]
		if tags.Ignore() {
			continue
		}
		if tags.IsBitflag {
			bits, err := d.PopUint()
			if err != nil {
				return errors.Wrapf(err, "getting %v flag", tags.Name)
			}
			bitflags[i] = bits

			continue
		}
		needToDecode := true
		if tags.BitFlags != nil {
			f, ok := bitflags[fields.bitflags[i].fieldIndex]
			if !ok {
				panic("internal error: " +
					"struct is not validated on register stage: " +
					"optional field was called BEFORE bitflags")
			}
			bitflagValue := f&(1<<tags.BitFlags.BitPosition) > 0

			if tags.Implicit {
				// implicit can be only boolean, so leave this initialization alone
				v.Field(i).Set(reflect.ValueOf(bitflagValue))

				continue
			}

			needToDecode = bitflagValue
		}

		if needToDecode {
			// normal field
			if err := d.decodeValue(v.Field(i)); err != nil {
				return wrapPath(err, tags.Name)
			}
		}
	}

	return nil
}

func (d *decoder) popInt() (int32, error) {
	val, err := d.Peek(0, WordLen)
	if err != nil {
		return 0, err
	}

	d.SkipBytes(WordLen)

	return int32(binary.LittleEndian.Uint32(val)), nil
}

func (d *decoder) popLong() (int64, error) {
	val, err := d.Peek(0, LongLen)
	if err != nil {
		return 0, err
	}

	d.SkipBytes(LongLen)

	return int64(binary.LittleEndian.Uint64(val)), nil
}

// popRequiredCRC reding crc from input, and compares with expected crc code.
func (d *decoder) popRequiredCRC(expected crc32) error {
	return nil
}

// popCRC just an alias for self documenting code.
func (d *decoder) PopCRC() (crc32, error)   { return d.PopUint() }
func (d *decoder) PopUint() (uint32, error) { return convertNumErr[uint32](d.popInt()) }

func (d *decoder) popDouble() (float64, error) {
	val, err := d.popLong()

	return math.Float64frombits(uint64(val)), err
}

func (d *decoder) PopBool() (bool, error) {
	crc, err := d.PopUint()
	if err != nil {
		return false, err
	}

	switch crc {
	case crcTrue:
		return true, nil
	case crcFalse:
		return false, nil
	default:
		return false, ErrInvalidBoolCRC(crc)
	}
}

func (d *decoder) popString() (string, error) { return convertStrErr[string](d.popMessage()) }
func (d *decoder) popMessage() ([]byte, error) {
	readLen := 1
	buf, err := d.Peek(0, 1)
	if err != nil {
		return nil, err
	}

	// how exact bytes there is a message
	var realSize int

	if firstByte := buf[0]; firstByte != fuckingMagicNumber {
		// it's a tiny message, so real size is exact a first byte value
		realSize = int(firstByte)
	} else {
		// otherwise it's a large message, next three bytes are size of message
		readLen = WordLen
		lenBuf, errPeek := d.Peek(1, WordLen-1)
		if errPeek != nil {
			return nil, errors.Wrapf(errPeek, "reading last %v bytes of message size", WordLen-1)
		}

		switch d.endianess {
		case binary.LittleEndian:
			lenBuf = append(lenBuf, 0)
		case binary.BigEndian:
			lenBuf = append([]byte{0}, lenBuf...)
		default:
			panic("wait, what?")
		}

		realSize = int(d.endianess.Uint32(lenBuf))
	}

	// this buf wil be real message
	buf, err = d.Peek(readLen, realSize)
	if err != nil {
		return nil, errors.Wrapf(err, "reading message data with len of %v", realSize)
	}
	d.SkipBytes(readLen + realSize + pad(readLen+realSize, WordLen))

	return buf, nil
}

func (d *decoder) Read(b []byte) (int, error) {
	if len(b)%WordLen != 0 {
		return 0, errors.New("value can't be divided by word length")
	}

	read, err := d.Peek(0, len(b))
	if err != nil {
		return 0, err
	}
	d.SkipBytes(len(b))

	return copy(b, read), nil
}

/*
	crcV := v.MapIndex(reflect.ValueOf(MapCrcKey))
	if !crcV.IsValid() {
		return errors.New("can't find " + MapCrcKey + " key")
	}
	if !crcV.Type().ConvertibleTo(uint32Typ) {
		return errors.New(MapCrcKey + ": can't convert to uint32")
	}
	crc = crcV.Convert(uint32Typ).Interface().(uint32)
*/

func getCRCofObject(v reflect.Value) (crc32, bool) {
	if _, ok := v.Interface().(Object); !ok {
		v = v.Addr()
	}

	if v, ok := v.Interface().(Object); ok {
		return v.CRC(), true
	}

	return 0, false
}
