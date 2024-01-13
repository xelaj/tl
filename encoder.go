// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

func Marshal(v any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := NewEncoder(buf)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// MarshalState provides set of different methods to marshal binary message into
// specific struct. It's working absolutely like fmt.State, you just need to
// write data to this abstraction.
type MarshalState interface {
	// if object will write bytes the size of which is not divided by the length
	// of the word, it will throw specific error
	io.Writer

	PutBool(bool) error
	PutInt(int32) error
	PutLong(int64) error
	PutCRC(crc32) error
	PutMessage([]byte) error
}

type RealEncoder interface {
	Encode(v any) error
}

// encoder is a type, which allows you to decode serialized message.
type encoder struct {
	w io.Writer

	endianess binary.ByteOrder
}

func NewEncoder(w io.Writer) RealEncoder {
	return &encoder{w: w, endianess: binary.LittleEndian}
}

func (e *encoder) Encode(value any) error {
	if value == nil {
		return ErrUnexpectedNil
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("value is not pointer as expected. got %v", v.Type())
	}

	return e.encodeValue(v)
}

// writeErr works like write (also throwing panic), but without count of
// written bytes.
func (e *encoder) writeErr(b []byte) error {
	if n, err := e.write(b); err != nil {
		return err
	} else if n != len(b) {
		return &ErrPartialWrite{Has: n, Want: len(b)}
	}

	return nil
}

func (e *encoder) Write(b []byte) (int, error) {
	if len(b)%WordLen != 0 {
		return 0, errors.New("value can't be divided by word length")
	}

	return e.write(b)
}

// write is private, cause this function might panic.
func (e *encoder) write(b []byte) (int, error) {
	if len(b)%WordLen != 0 { //revive:disable-line:add-constant // makes no sense
		// it's panic, because it's internal method, and we must not write in
		// any case data, which is not divided by word length
		panic("raw bytes does not divide by word size of protocol")
	}

	if len(b) == 0 {
		return 0, nil
	}

	return e.w.Write(b) //nolint:wrapcheck // write() is a wrapper
}

//nolint:cyclop // it contains only assertion and switch statement
//revive:disable:function-length // same: can't make better
func (e *encoder) encodeValue(value reflect.Value) error {
	if maybeNil(value) {
		return ErrUnexpectedNil
	}

	if marshaler, ok := value.Interface().(Marshaler); ok {
		//nolint:wrapcheck // object implements Marshaler must throw unwrapped
		//                    error
		return marshaler.MarshalTL(e)
	}

	switch k := value.Type().Kind(); k { //nolint:exhaustive // has default case
	case reflect.Uint32:
		return e.putUint(uint32(value.Uint()))

	case reflect.Int32:
		return e.putUint(uint32(value.Int()))

	case reflect.Int64:
		return e.PutLong(value.Int())

	case reflect.Float64:
		return e.putDouble(value.Float())

	case reflect.Bool:
		return e.PutBool(value.Bool())

	case reflect.String:
		return e.putString(value.String())

	case reflect.Struct:
		return e.encodeStruct(value, false)

	case reflect.Ptr, reflect.Interface:
		return e.encodeValue(value.Elem())

	// case reflect.Map:
	//	return e.encodeMap(value)

	case reflect.Slice:
		return e.encodeVector(value)
	case reflect.Array:
		if value.Type().Elem() == byteTyp { // [N]byte
			return e.encodeRaw(value)
		}

		return e.encodeVector(value)

	default:
		return ErrUnsupportedType{Type: value.Type()}
	}
}

//revive:enable

// v must be pointer to struct.
func (e *encoder) encodeStruct(v reflect.Value, ignoreCRC bool) error {
	o, ok := v.Interface().(Object)
	if !ok {
		// Trying to look implementation by pointer
		//
		// Since just struct might be non addressable, we are creating new
		// instance, and setting it to provided value.
		vcopy := reflect.New(v.Type()).Elem()
		vcopy.Set(v)

		var ok bool
		if o, ok = vcopy.Addr().Interface().(Object); !ok {
			return errors.New(v.Type().String() + " doesn't implement tl.Object interface")
		}
	}

	typ := v.Type()

	properties, bitflags, err := parseStructTags(typ)
	if err != nil {
		return fmt.Errorf("parsing struct flags: %w", err)
	}

	optFlags := make(map[int]crc32)
	for i, target := range bitflags {
		// even if we don't have any non null optional values, we still need to
		// initialize bitflags
		if _, ok := optFlags[target.FieldIndex]; !ok {
			optFlags[target.FieldIndex] = 0
		}
		if isFieldContainsData(v.Field(i)) {
			optFlags[target.FieldIndex] |= 1 << target.BitIndex
		}
	}

	// what we checked and what we know about value:
	// 1) it's not Marshaler (marshaler method if exist used already in c.encodeValue())
	// 2) implements tl.Object
	// 3) it's addressable
	// 4) definitely struct (we don't call encodeStruct() but in c.encodeValue())
	// 5) not nil (structs can't be nil, only pointers and interfaces)
	if !ignoreCRC {
		if err := e.PutCRC(o.CRC()); err != nil {
			return err
		}
	}

	for i := 0; i < v.NumField(); i++ {
		// putting bitflags, if this field is bitflag
		if flags, ok := optFlags[i]; ok {
			if err := e.putUint(flags); err != nil {
				return err
			}

			continue
		}

		tag := properties[i]
		_, fieldOptional := bitflags[i]

		// if ignore or field is unexported, then go on
		if tag.Ignore() ||
			!v.Field(i).CanSet() ||
			(fieldOptional && (!isFieldContainsData(v.Field(i)) || tag.isImplicit())) {
			continue
		}

		err := e.encodeValue(v.Field(i))
		if err != nil {
			return fmt.Errorf("encoding %v: %w", v.Type().Field(i).Name, err)
		}
	}

	return nil
}

/*
func (e *encoder) encodeMap(m reflect.Value) error {
	if m.Type().Key().Kind() != reflect.String {
		return errors.New("map keys are not string")
	}

	crc, err := getCRCFromMap(m)
	if err != nil {
		return err
	}

	definition, ok := e.registry.Tags(crc)
	if !ok {
		//nolint:goerr113 // it's an internal error
		return fmt.Errorf("crc code 0x%08x is not found in registry", crc)
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
*/

func (e *encoder) encodeRaw(v reflect.Value) error {
	if v.Kind() != reflect.Array {
		panic("raw must be array")
	} else if v.Type().Elem() != byteTyp {
		panic("raw must be array of bytes")
	} else if n := v.Len(); n%WordLen != 0 {
		// special case: this means that we want to take exact N of bytes and pop it from reader
		// n%WordLen == 0, cause we can't read less or more than word
		return fmt.Errorf("array of bytes must be divided by %v, got %v", WordLen, n)
	}

	_, err := e.Write(v.Slice(0, v.Len()).Interface().([]byte))

	return err
}

func (e *encoder) encodeVector(slice reflect.Value) error {
	if b, ok := slice.Interface().([]byte); ok {
		return e.PutMessage(b)
	}

	if err := e.PutCRC(crcVector); err != nil {
		return err
	}
	if err := e.putUint(uint32(slice.Len())); err != nil {
		return err
	}

	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		err := e.encodeValue(item)
		if err != nil {
			return fmt.Errorf("[%v]: %w", i, err)
		}
	}

	return nil
}

func (e *encoder) putUint(v uint32) error    { return e.writeErr(u32b(e.endianess, v)) }
func (e *encoder) PutLong(v int64) error     { return e.writeErr(u64b(e.endianess, uint64(v))) }
func (e *encoder) putDouble(v float64) error { return e.writeErr(f64b(e.endianess, v)) }
func (e *encoder) PutCRC(v uint32) error     { return e.putUint(v) } // for selfdoc code
func (e *encoder) PutInt(v int32) error      { return e.putUint(uint32(v)) }
func (e *encoder) PutBool(v bool) error      { return e.putUint(boolToCRC(v)) }
func (e *encoder) putString(v string) error  { return e.PutMessage([]byte(v)) }

func (e *encoder) PutMessage(msg []byte) error {
	// 3 left bytes of word, which is 4 bytes
	const maxLen = 1 << ((WordLen - 1) * bitsInByte)
	if len(msg) > maxLen {
		//nolint:goerr113 // it's an internal error
		return fmt.Errorf("message entity too large: expect less than %v, got %v", maxLen, len(msg))
	}

	var lenBytes []byte

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
		lenBytes = []byte{byte(len(msg))}
	} else {
		lenBytes = append([]byte{fuckingMagicNumber}, littleUint24Bytes(len(msg))...)
	}

	return e.writeErr(appendMany(
		lenBytes,
		msg,
		make([]byte, pad(len(lenBytes)+len(msg), WordLen)),
	))
}

// m only map.
func getCRCFromMap(m reflect.Value) (uint32, error) {
	crcVal := m.MapIndex(reflect.ValueOf(MapCrcKey))
	if !crcVal.IsValid() {
		return 0, errors.New("key " + MapCrcKey + " not exist in map")
	}
	if !crcVal.Type().ConvertibleTo(uint32Typ) {
		return 0, errors.New(MapCrcKey + " is not convertible to uint32")
	}

	return uint32(crcVal.Convert(uint32Typ).Uint()), nil
}
