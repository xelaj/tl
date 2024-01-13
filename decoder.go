// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
)

func Unmarshal(b []byte, res any) error {
	return NewDecoder(bytes.NewBuffer(b)).Decode(res)
}

type Decoder interface {
	SetRegistry(registry Registry) Decoder

	Decode(res any) error
}

type UnmarshalState interface {
	// ReadWords reads words from unmarshaler with fixed size of word.
	io.Reader
	ReadAll() ([]byte, error)
	Peek(seek, size int) ([]byte, error)
	SkipBytes(int)

	PopBool() (bool, error)
	PopInt() (int32, error)
	PopLong() (int64, error)
	PopCRC() (crc32, error)
	PopMessage() ([]byte, error)
}

// A decoder reads and decodes TL values from an input stream.
type decoder struct { //nolint:govet // false positive for fieldalignment
	r bufio.Reader

	registry  Registry
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
		registry:  DefaultRegistry(),
		endianess: binary.LittleEndian,
	}
}

func (d *decoder) SetRegistry(registry Registry) Decoder {
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

func (d *decoder) Peek(seek, size int) ([]byte, error) {
	peeked, err := d.r.Peek(seek + size)
	if err != nil {
		return nil, err //nolint:wrapcheck // must not be wrapped
	}

	res := make([]byte, size)
	if l := copy(res, peeked[seek:]); l != len(res) {
		// just for test, who knows
		panic(fmt.Sprintf("copying bytes: expected %v, got %v", size, l))
	}

	return res, nil
}

func (d *decoder) SkipBytes(n int) {
	// discarded will be only exact number or less, BUT with returned error
	if _, err := d.r.Discard(n); err != nil {
		panic(fmt.Errorf("something wrong with peeking bytes, Discard must always be ok: %w", err))
	}
}

func (d *decoder) decodeValue(value reflect.Value) error {
	if v, ok := value.Interface().(Unmarshaler); ok {
		return v.UnmarshalTL(d)
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
		val, err = d.PopLong()

	case reflect.Uint32:
		if _, ok := value.Interface().(Object); ok {
			crc, err := d.PopCRC()
			if err != nil {
				return err
			}

			return d.decodeEnum(value, crc)
		}

		val, err = d.PopUint()

	case reflect.Int32:
		val, err = d.PopInt()

	case reflect.Bool:
		val, err = d.PopBool()

	case reflect.String:
		val, err = d.popString()

	case reflect.Slice:
		switch value.Type() {
		case byteSliceTyp: // []byte
			val, err = d.PopMessage()

		default:
			return d.decodeVector(value, false)
		}

	case reflect.Array:
		if value.Type().Elem() == byteTyp { // [N]byte
			return d.decodeRaw(value)
		}

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
		if err := d.decodeInterface(value); err != nil {
			return fmt.Errorf("decoding interface: %w", err)
		}
		return nil

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

	v := reflect.ValueOf(&val).Elem().Elem()
	value.Set(v)

	return nil
}

func (d *decoder) decodeRaw(v reflect.Value) error {
	if v.Kind() != reflect.Array {
		panic("raw must be array")
	} else if v.Type().Elem() != byteTyp {
		panic("raw must be array of bytes")
	} else if n := v.Len(); n%WordLen != 0 {
		// special case: this means that we want to take exact N of bytes and pop it from reader
		// n%WordLen == 0, cause we can't read less or more than word
		return fmt.Errorf("array of bytes must be divided by %v, got %v", WordLen, n)
	}

	val, err := d.Peek(0, v.Len())
	if err != nil {
		return err
	}
	d.SkipBytes(v.Len())

	valRaw := reflect.ValueOf(&val).Elem()
	arr := reflect.New(reflect.ArrayOf(v.Len(), byteTyp)).Elem()
	for i := 0; i < v.Len(); i++ {
		arr.Index(i).Set(valRaw.Index(i))
	}
	v.Set(arr)

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
		return fmt.Errorf("read vector size: %w", err)
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

	if crc == crcGzipPacked {
		// ! special case for gzip packed objects
		//
		// serialized data MIGHT place gzipped objects (and only vectors and
		// objects!) in a random places, where it looks like the message is "too
		// big". To handle that, here we unzipping it, and replacing original
		// reader with custom one.
		//
		// Important: according to (https://core.telegram.org/api/invoking, it's
		// "recommended" to zip large messages from client side, but telegram
		// is... Telegram, so it might argue sometimes. But in any case, it
		// makes no sense to implement mtproto extension of "classic" tl
		// serialization as required component)
		gz, err := d.popZip(true)
		if err != nil {
			return err
		}

		// replacing reader with custom, which will read from gzip reader.
		old := d.r
		defer func() { d.r = old }()
		d.r = *bufio.NewReader(gz)

		crc, err = d.PopCRC()
		if err != nil {
			return errReadCRC{err}
		}
	}

	obj, ok := d.registry.ConstructObject(crc)
	if !ok {
		return ErrObjectNotRegistered(crc)
	}
	o := reflect.ValueOf(&obj).Elem()

	if o.Kind() != reflect.Interface {
		panic("unexpected object which is not interface")
	}
	if !o.Elem().IsValid() {
		panic("constructor must return non empty interface")
	}

	if o.Elem().Kind() == reflect.Ptr && o.Elem().IsNil() {
		realType := o.Elem().Type().Elem()
		o.Set(reflect.New(realType))
	}

	if unmarshaler, ok := o.Interface().(Unmarshaler); ok {
		if err := unmarshaler.UnmarshalTL(d); err != nil {
			return err
		}
		v.Set(o)

		return nil
	}

	o = o.Elem().Elem()

	if o.Kind() != reflect.Struct {
		return fmt.Errorf("object must be struct, got %v", o.Type())
	}

	err = d.decodeObject(o, true)
	if err != nil {
		return err // no need to wrap
	}

	// set value of o into v with interface conversion
	vtyp := reflect.New(v.Type()).Elem()

	if !o.Type().Implements(vtyp.Type()) {
		if !o.Addr().Type().Implements(vtyp.Type()) {
			panic("can't find interface implementation")
		}
		o = o.Addr()
	}

	v.Set(o)

	return nil
}

func (d *decoder) decodeEnum(v reflect.Value, crc crc32) error {
	enum, ok := d.registry.ConstructObject(crc)
	if !ok {
		return fmt.Errorf("enum 0x%08x not found", crc)
	}
	value := reflect.ValueOf(enum)
	if v.Type() != value.Type() {
		return fmt.Errorf("invalid type of enum: want %v, got %v", v.Type(), value.Type())
	} else if value.Kind() != reflect.Uint32 {
		panic("internal error: enum must not have fields")
	}

	v.Set(value)

	return nil
}

// decode EXACT object. means that v must always be struct.
func (d *decoder) decodeObject(v reflect.Value, ignoreCRC bool) error {
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("expected struct, got %v with %v kind", v.Type(), v.Kind()))
	}

	crc, ok := getCRCofObject(v)
	if !ok {
		return errors.New("value must implement Object interface, got: " + v.Type().String())
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

	bitflags := map[string]uint32{}
	for i := 0; i < v.NumField(); i++ {
		typ := v.Type().Field(i)
		tags, err := ParseTag(typ.Tag.Get(tagName), typ.Name)
		if err != nil {
			panic(fmt.Sprintf("invalid tag: %v", err))
		}

		if tags.Ignore() {
			continue
		}
		if tags.isBitflag() {
			bits, err := d.PopUint()
			if err != nil {
				return fmt.Errorf("getting %v flag: %w", tags.Name, err)
			}
			bitflags[tags.Name] = bits

			continue
		}
		needToDecode := true
		if tags.BitFlags != nil {
			f, ok := bitflags[tags.BitFlags.TargetField]
			if !ok {
				panic("internal error: " +
					"struct is not validated on register stage: " +
					"optional field was called BEFORE bitflags")
			}
			bitflagValue := f&(1<<tags.BitFlags.BitPosition) > 0

			if tags.isImplicit() {
				// implicit can be only boolean, so leave this initialization alone
				v.Field(i).Set(reflect.ValueOf(bitflagValue))

				continue
			}

			needToDecode = bitflagValue
		}

		if needToDecode {
			// normal field
			if v.Field(i).Kind() == reflect.Ptr {
				v.Field(i).Set(reflect.New(v.Field(i).Type().Elem()))
			}

			if err := d.decodeValue(v.Field(i)); err != nil {
				return wrapPath(err, v.Type().Field(i).Name)
			}
		}
	}

	return nil
}

func (d *decoder) PopInt() (int32, error) {
	val, err := d.Peek(0, WordLen)
	if err != nil {
		return 0, err
	}

	d.SkipBytes(WordLen)

	return int32(binary.LittleEndian.Uint32(val)), nil
}

func (d *decoder) PopLong() (int64, error) {
	val, err := d.Peek(0, LongLen)
	if err != nil {
		return 0, err
	}

	d.SkipBytes(LongLen)

	return int64(binary.LittleEndian.Uint64(val)), nil
}

// popCRC just an alias for self documenting code.
func (d *decoder) PopCRC() (crc32, error)   { return d.PopUint() }
func (d *decoder) PopUint() (uint32, error) { return convertNumErr[uint32](d.PopInt()) }

func (d *decoder) popDouble() (float64, error) {
	val, err := d.PopLong()

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

// https://core.telegram.org/type/bytes
func (d *decoder) popString() (string, error) { return convertStrErr[string](d.PopMessage()) }

func (d *decoder) PopMessage() ([]byte, error) {
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
			return nil, fmt.Errorf("reading last %v bytes of message size: %w", WordLen-1, errPeek)
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
		return nil, fmt.Errorf("reading message data with len of %v: %w", realSize, err)
	}
	d.SkipBytes(readLen + realSize + pad(readLen+realSize, WordLen))

	return buf, nil
}

// ! special case for gzip packed objects
//
// serialized data MIGHT have replacements from original objects to gzipped
// objects (and only vectors and objects!) in a random places, where it looks
// like the message is "too big". To handle that, here we unzipping it, and
// replacing original reader with custom one.
//
// Important: according to https://core.telegram.org/api/invoking, it's
// "recommended" to zip large messages from client side, but telegram
// is... Telegram, so it might argue sometimes. But in any case, it
// makes no sense to implement mtproto extension of "classic" tl
// serialization as required component).
func (d *decoder) popZip(ignoreCRC bool) (io.ReadCloser, error) {
	if !ignoreCRC {
		gotCrc, err := d.PopCRC()
		if err != nil {
			return nil, errReadCRC{err}
		}

		if gotCrc != crcGzipPacked {
			return nil, ErrInvalidCRC{Got: gotCrc, Want: crcGzipPacked}
		}
	}

	data, err := d.PopMessage()
	if err != nil {
		return nil, fmt.Errorf("reading gzip value: %w", err)
	}

	return gzip.NewReader(bytes.NewBuffer(data))
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

func (d *decoder) ReadAll() ([]byte, error) { return io.ReadAll(&d.r) }

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
