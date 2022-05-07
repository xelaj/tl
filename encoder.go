// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"

	"github.com/pkg/errors"
)

func Marshal(v any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := NewEncoder(buf)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type MarshalProvider interface {
	// if object will write bytes the size of which is not divided by the length
	// of the word, it will throw specific error
	io.Writer
}

type Encoder struct {
	w        io.Writer
	registry *Registry

	endianess binary.ByteOrder
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w, registry: defaultRegistry, endianess: binary.LittleEndian}
}

func (e *Encoder) SetRegistry(registry *Registry) *Encoder {
	e.registry = registry

	return e
}

func (e *Encoder) Encode(v any) error {
	return e.encodeValue(reflect.ValueOf(v))
}

// writeErr works like write (also throwing panic), but without count of
// written bytes.
func (e *Encoder) writeErr(b []byte) error {
	if n, err := e.write(b); err != nil {
		return err
	} else if n != len(b) {
		return &ErrorPartialWrite{Has: n, Want: len(b)}
	}

	return nil
}

// write is private, cause this function might panic.
func (e *Encoder) write(b []byte) (int, error) {
	if len(b)%WordLen != 0 { //revive:disable:add-constant // makes no sense
		// it's panic, because it's internal method, and we must not write in
		// any case data, which is not divided by word length
		panic("raw bytes does not divide by word size of protocol")
	}

	if len(b) == 0 {
		return 0, nil
	}

	return e.w.Write(b) //nolint:wrapcheck // write() is a wrapper
}

//nolint:gocyclo,cyclop // it contains only assertion and switch statement
//revive:disable:function-length // same: can't make better
func (e *Encoder) encodeValue(value reflect.Value) error {
	if maybeNil(value) {
		return ErrUnexpectedNil{}
	}

	if marshaler, ok := value.Interface().(Marshaler); ok {
		//nolint:wrapcheck // object implements Marshaler must throw unwrapped
		//                    error
		return marshaler.MarshalTL(e.provider())
	}

	switch k := value.Type().Kind(); k { //nolint:exhaustive // has default case
	case reflect.Uint32:
		return e.putUint(uint32(value.Uint()))

	case reflect.Int32:
		return e.putUint(uint32(value.Int()))

	case reflect.Int64:
		return e.putLong(value.Int())

	case reflect.Float64:
		return e.putDouble(value.Float())

	case reflect.Bool:
		return e.putBool(value.Bool())

	case reflect.String:
		return e.putString(value.String())

	case reflect.Struct:
		return e.encodeStruct(value.Addr())

	case reflect.Ptr, reflect.Interface:
		return e.encodeValue(value.Elem())

	case reflect.Map:
		return e.encodeMap(value)

	case reflect.Slice:
		return e.encodeVector(value)

	case reflect.Int, reflect.Uint:
		return ErrImplicitInteger{}

	case reflect.Int8, reflect.Int16, reflect.Uint8, reflect.Uint16, reflect.Uint64:
		return ErrUnsupportedInt{Kind: k}

	case reflect.Float32, reflect.Complex64, reflect.Complex128:
		return ErrUnsupportedFloat{Kind: k}

	default:
		return ErrUnsupportedType{Type: value.Type()}
	}
}

// v must be pointer to struct.
func (e *Encoder) encodeStruct(v reflect.Value) error {
	o, ok := v.Interface().(Object)
	if !ok {
		return errors.New(v.Type().String() + " doesn't implement tl.Object interface")
	}

	v = reflect.Indirect(v)

	properties, ok := e.registry.structFields[o.CRC()]
	if !ok {
		return errors.New(v.Type().String() + " is not registered type")
	}

	optFlags := make(map[int]crc32)
	for i, target := range properties.bitflags {
		// even if we don't have any non null optional values, we still need to initialize bitflags
		if _, ok := optFlags[target.fieldIndex]; !ok {
			optFlags[target.fieldIndex] = 0
		}
		if isFieldContainsData(v.Field(i)) {
			optFlags[target.fieldIndex] |= 1 << target.bitIndex
		}
	}

	// what we checked and what we know about value:
	// 1) it's not Marshaler (marshaler method if exist used already in c.encodeValue())
	// 2) implements tl.Object
	// 3) definitely struct (we don't call encodeStruct(), only in c.encodeValue())
	// 4) not nil (structs can't be nil, only pointers and interfaces)
	if err := e.putCRC(o.CRC()); err != nil {
		return err
	}

	for i := 0; i < v.NumField(); i++ {
		// putting bitflags, if this field is bitflag
		if flags, ok := optFlags[i]; ok {
			if err := e.putUint(flags); err != nil {
				return err
			}

			continue
		}

		tag := properties.tags[i]

		// if ignore or field is unexported, then go on
		if tag.Ignore() ||
			!v.Field(i).CanSet() ||
			(properties.isFieldOptional(i) && !isFieldContainsData(v.Field(i))) {
			continue
		}

		if _, ok := properties.bitflags[i]; ok && tag.Implicit() {
			continue
		}

		err := e.encodeValue(v.Field(i))
		if err != nil {
			return errors.Wrapf(err, "encoding %v", v.Type().Field(i).Name)
		}
	}

	return nil
}

func (e *Encoder) encodeMap(m reflect.Value) error {
	if m.Type().Key().Kind() != reflect.String {
		return errors.New("map keys are not string")
	}

	crc, err := getCRCFromMap(m)
	if err != nil {
		return err
	}

	definition, ok := e.registry.orphans[crc]
	if !ok {
		//nolint:goerr113 // it's an internal error
		return fmt.Errorf("crc code %08x is not found in registry", crc)
	}

	// TODO: need to cache encoded non empty objects in slice, then write
	//       everything after we will be sure that cached bitflag will not be
	//       changed (means that we checked already most right optional field
	//       for this bitflag)
	bitflags, err := definition.collectBitflags(m)
	if err != nil {
		return err
	}

	if err := e.putCRC(crc); err != nil {
		return err
	}

	for i, field := range definition.fields {
		// putting bitflags, if this field is bitflag
		if flags, ok := bitflags[uint8(i)]; ok {
			if err := e.putUint(flags); err != nil {
				return err
			}

			continue
		}

		val := m.MapIndex(reflect.ValueOf(field.name))
		if !val.IsValid() || field.noEncode {
			continue
		}

		if err := e.encodeValue(val); err != nil {
			return errors.Wrapf(err, "encoding %q", field.name)
		}
	}

	return nil
}

// func (e *Encoder) collectBitflags()

func (e *Encoder) encodeVector(slice reflect.Value) error {
	if b, ok := slice.Interface().([]byte); ok {
		return e.putMessage(b)
	}

	if err := e.putCRC(CrcVector); err != nil {
		return err
	}
	if err := e.putUint(uint32(slice.Len())); err != nil {
		return err
	}

	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		err := e.encodeValue(item)
		if err != nil {
			return errors.Wrapf(err, "[%v]", i)
		}
	}

	return nil
}

func (e *Encoder) putUint(v uint32) error    { return e.writeErr(u32b(e.endianess, v)) }
func (e *Encoder) putLong(v int64) error     { return e.writeErr(u64b(e.endianess, uint64(v))) }
func (e *Encoder) putDouble(v float64) error { return e.writeErr(f64b(e.endianess, v)) }
func (e *Encoder) putCRC(v uint32) error     { return e.putUint(v) } // for selfdoc code
func (e *Encoder) putInt(v int32) error      { return e.putUint(uint32(v)) }
func (e *Encoder) putBool(v bool) error      { return e.putUint(boolToCRC(v)) }
func (e *Encoder) putString(v string) error  { return e.putMessage([]byte(v)) }

func (e *Encoder) putMessage(msg []byte) error {
	// 3 left bytes of word, which is 4 bytes
	const maxLen = 1 << ((WordLen - 1) * bitsInByte)
	if len(msg) > maxLen {
		//nolint:goerr113 // it's an internal error
		return fmt.Errorf("message entity too large: expect less than %v, got %v", maxLen, len(msg))
	}

	// how does it works:
	// any object can be putted to byte set ONLY with length, without modula
	// after dividing to word length. e.g. bytes 'Hi!' can be written as:
	//            | 0x03 0x48 0x6A 0x21 |
	// Divides by 32 bits? Yes, so it's good.
	//
	// BUT! bytes 'Hello!' MUST be written as
	//            | 0x06 0x48 0x65 0x6C | 0x6C 0x6F 0x21 0x00 |
	// See? We added extra empty byte to pad message to length of word. That is
	// most important part of putting bytes to buffer.
	//
	// So we must to create a buffer with length mod to 32 == 0. To not add
	// extra bytes manually. They could be random, but who needs that, right?
	if len(msg) < fuckingMagicNumber {
		// tiny message
		const byteStreamLengthSize = 1

		msg = appendMany(
			[]byte{byte(len(msg))},
			msg,
			make([]byte, pad(byteStreamLengthSize, WordLen, len(msg))),
		)
	} else {
		// large message
		const byteStreamLengthSize = 4

		msg = appendMany(
			[]byte{fuckingMagicNumber},
			littleUint24Bytes(len(msg)),
			msg,
			make([]byte, pad(byteStreamLengthSize, WordLen, len(msg))),
		)
	}

	return e.writeErr(msg)
}

func (e *Encoder) provider() MarshalProvider { return ptr(marshaler(*e)) }

type marshaler Encoder

var _ MarshalProvider = (*marshaler)(nil)

func (m *marshaler) e() *Encoder { return (*Encoder)(m) }

func (m *marshaler) Write(b []byte) (int, error) {
	if len(b)%WordLen != 0 {
		return 0, errors.New("value can't be divided by word length")
	}

	return m.e().write(b)
}

// m only map.
func getCRCFromMap(m reflect.Value) (uint32, error) {
	crcVal := m.MapIndex(reflect.ValueOf(MapCrcKey))
	if !crcVal.IsValid() {
		return 0, errors.New("key " + MapCrcKey + " not exist in map")
	}
	if !crcVal.Type().ConvertibleTo(uint32Type) {
		return 0, errors.New(MapCrcKey + " is not convertible to uint32")
	}

	return uint32(crcVal.Convert(uint32Type).Uint()), nil
}

func maybeNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
