// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type ErrOutIsNil struct{}

func (e ErrOutIsNil) Error() string {
	return "can't unmarshal to nil value"
}

func Unmarshal(b []byte, res any) error {
	return NewDecoder(bytes.NewBuffer(b)).Decode(res)
}

func UnmarshalUnknown(b []byte) (any, error) {
	return NewDecoder(bytes.NewBuffer(b)).DecodeUnknown()
}

func (d *Decoder) Decode(res any) error {
	if res == nil {
		return ErrOutIsNil{}
	}
	v := reflect.ValueOf(res)

	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("res value is not pointer as expected. got %v", v.Type())
	}

	// decoding elem cause we are taking pointer in res, but we'll take additional step to call decodeValue
	// again. Who needs that?
	return d.decodeValue(v.Elem())
}

// DecodeUnknown works like Decode, but it tries to get object stored in data stream
func (d *Decoder) DecodeUnknown() (any, error) {
	crc, err := d.PopCRC()
	if err != nil {
		return nil, errors.Wrap(err, "getting crc code of decoded object")
	}

	switch crc {
	case CrcVector:
		return nil, errors.New("got vector, not allowed to decode it manually")
	case CrcFalse:
		return false, nil
	case CrcTrue:
		return true, nil
	case CrcNull:
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

func (d *Decoder) decodeValue(value reflect.Value) error {
	if value.Type().Implements(unmarshalerTyp) {
		return value.Interface().(Unmarshaler).UnmarshalTL(d)
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

	switch kind := value.Kind(); kind {
	// unsupported
	case reflect.Chan, reflect.Func, reflect.Uintptr, reflect.UnsafePointer:
		return errors.New(value.Kind().String() + " isn't supported")

	// supported, but TL is not supported 8 and 16 bit numbers
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64:
		// as discussed in encoder: we can't be sure that we decode 4 bytes or 8, or even 16. so throwing
		// an error
		return fmt.Errorf("int kind: %v (must converted to int32, int64 or uint32 explicitly)", kind)

	// same: supported, but not in TL, so we can't understand, how much bytes we need to scan.
	case reflect.Float32, reflect.Complex64, reflect.Complex128:
		return fmt.Errorf("float kind: %s (must be converted to float64 explicitly)", kind)

	// pointer always diving into value
	case reflect.Ptr:
		// need to init firstly
		if value.IsZero() {
			value.Set(reflect.New(value.Type().Elem()))
		}

		return d.decodeValue(value.Elem())

	// simple types
	case reflect.Float64:
		val, err = d.PopDouble()

	case reflect.Int64:
		val, err = d.PopLong()

	case reflect.Uint32: // no need to handle enums explicitly, they are decoding here
		val, err = d.PopUint()

	case reflect.Int32:
		val, err = d.PopInt()

	case reflect.Bool:
		val, err = d.PopBool()

	case reflect.String:
		val, err = d.PopString()

	case reflect.Slice:
		switch value.Type() {
		case byteSliceTyp: // []byte
			val, err = d.PopMessage()

		default:
			return d.decodeVector(value, false)
		}

	case reflect.Array:
		return d.decodeVector(value, false)

	// complex types
	case reflect.Struct:
		return d.decodeObject(value, false)

	// struct but it's map. Maps are not defined by TL, so we use them as type values.
	case reflect.Map:
		// TODO: must implement map decoding
		return errors.New("map is not ordered object (must order like structs)")

	// interfaces in terms of TL is a set of allowed structs.
	case reflect.Interface:
		return d.decodeInterface(value)

	default:
		panic("unknown thingie: " + value.Type().String())
	}

	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(val))
	return nil
}

// allowed values are only slice and array.
func (d *Decoder) decodeVector(v reflect.Value, ignoreCRC bool) error {
	if !ignoreCRC {
		crc, err := d.PopCRC()
		if err != nil {
			return errors.Wrap(err, "read crc")
		}

		if crc != CrcVector {
			return fmt.Errorf("not a vector: 0x%08x, want: 0x%08x", crc, CrcVector)
		}
	}

	size, err := d.PopUint()
	if err != nil {
		return errors.Wrap(err, "read vector size")
	}

	switch v.Kind() {
	case reflect.Array:
		if v.Len() < int(size) {
			return fmt.Errorf("array size is smaller than message: got %v, want %v", v.Len(), size)
		}
	case reflect.Slice:
		v.Set(reflect.MakeSlice(v.Type(), int(size), int(size)))
	}

	for i := 0; i < int(size); i++ {
		if err := d.decodeValue(v.Index(i)); err != nil {
			return wrapPath(err, "["+strconv.Itoa(i)+"]")
		}
	}

	return nil
}

func (d *Decoder) decodeInterface(v reflect.Value) error {
	crc, err := d.PopCRC()
	if err != nil {
		return errors.Wrap(err, "read crc")
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

func (d *Decoder) decodeEnum(v reflect.Value, crc crc32) error {
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

// decode EXACT object. means that v must always be struct
func (d *Decoder) decodeObject(v reflect.Value, ignoreCRC bool) error {
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
			return errors.Wrap(err, "read crc")
		}

		if gotCrc != crc {
			return fmt.Errorf("invalid crc code: got 0x%08x; want 0x%08x", gotCrc, crc)
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
				panic("internal error: struct is not validated on register stage: optional field called BEFORE bitflags")
			}
			bitflagValue := f&(1<<tags.BitFlags.BitPosition) > 0

			if tags.Implicit() {
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

/*
	crcV := v.MapIndex(reflect.ValueOf(MapCrcKey))
	if !crcV.IsValid() {
		return errors.New("can't find " + MapCrcKey + " key")
	}
	if !crcV.Type().ConvertibleTo(uint32Type) {
		return errors.New(MapCrcKey + ": can't convert to uint32")
	}
	crc = crcV.Convert(uint32Type).Interface().(uint32)
*/

func getCRCofObject(v reflect.Value) (crc32, bool) {
	if !v.Type().Implements(objectTyp) {
		v = v.Addr()
	}
	if !v.Type().Implements(objectTyp) {
		return 0, false
	}

	return v.Interface().(Object).CRC(), true
}
