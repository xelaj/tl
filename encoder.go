// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

func Marshal(v any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := NewEncoder(buf)
	err := encoder.encodeValue(reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Encoder) encodeValue(value reflect.Value) error {
	if value.Type().Implements(marshalerTyp) {
		return value.Interface().(Marshaler).MarshalTL(c)
	}

	switch k := value.Type().Kind(); k { //nolint:exhaustive has default case
	case reflect.Uint32:
		return c.PutUint(uint32(value.Uint()))

	case reflect.Int32:
		return c.PutUint(uint32(value.Int()))

	case reflect.Int64:
		return c.PutLong(value.Int())

	case reflect.Float64:
		return c.PutDouble(value.Float())

	case reflect.Bool:
		return c.PutBool(value.Bool())

	case reflect.String:
		return c.PutString(value.String())

	case reflect.Struct:
		return c.encodeStruct(value.Addr())

	case reflect.Ptr, reflect.Interface:
		if value.IsNil() {
			return ErrUnexpectedNil{}
		}

		return c.encodeValue(value.Elem())

	case reflect.Map:
		if value.Type().Key().Kind() != reflect.String {
			return errors.New("map keys are not string")
		}

		return c.encodeMap(value)

	case reflect.Slice:
		if b, ok := value.Interface().([]byte); ok {
			return c.PutMessage(b)
		}

		return c.encodeVector(value)

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

// v must be pointer to struct
func (c *Encoder) encodeStruct(v reflect.Value) error {
	o, ok := v.Interface().(Object)
	if !ok {
		return errors.New(v.Type().String() + " doesn't implement tl.Object interface")
	}

	v = reflect.Indirect(v)

	properties, ok := c.registry.structFields[o.CRC()]
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
	if err := c.PutCRC(o.CRC()); err != nil {
		return err
	}

	for i := 0; i < v.NumField(); i++ {
		// putting bitflags, if this field is bitflag
		if flags, ok := optFlags[i]; ok {
			if err := c.PutUint(flags); err != nil {
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

		err := c.encodeValue(v.Field(i))
		if err != nil {
			return errors.Wrapf(err, "encoding %v", v.Type().Field(i).Name)
		}
	}

	return nil
}

var uintType = reflect.TypeOf(uint32(0))

func getCRCFromMap(m reflect.Value) (uint32, error) {
	crcVal := m.MapIndex(reflect.ValueOf(MapCrcKey))
	if !crcVal.IsValid() {
		return 0, errors.New("key " + MapCrcKey + " not exist in map")
	}
	if !crcVal.Type().ConvertibleTo(uintType) {
		return 0, errors.New(MapCrcKey + " is not convertible to uint32")
	}
	return crcVal.Convert(uintType).Interface().(uint32), nil
}

func (c *Encoder) encodeMap(m reflect.Value) error {
	crc, err := getCRCFromMap(m)
	if err != nil {
		return err
	}

	definition, ok := c.registry.orphans[crc]
	if !ok {
		return fmt.Errorf("crc code %08x is not found in registry", crc)
	}

	// init bitlflags
	// TODO: need to cache encoded non empty objects in slice, then write everything after we will be sure
	//       that cached bitflag will not be changed (means that we checked already most right optional field
	//       for this bitflag)
	bitflags := map[uint8]crc32{}
	for i, field := range definition.fields {
		if _, ok := field.fType.(fieldBitflag); ok {
			bitflags[uint8(i)] = 0
			continue
		}

		ok := m.MapIndex(reflect.ValueOf(field.name)).IsValid()
		if field.optional && (ok || field.NoEncode) {
			bitflags[field.FlagTrigger] |= crc32(1 << field.BitTrigger)
		} else if !field.optional && !ok {
			return fmt.Errorf("field %q is required for this type", field.name)
		}
	}

	if err := c.PutCRC(crc); err != nil {
		return err
	}

	for i, field := range definition.fields {
		// putting bitflags, if this field is bitflag
		if flags, ok := bitflags[uint8(i)]; ok {
			if err := c.PutUint(flags); err != nil {
				return err
			}
			continue
		}

		val := m.MapIndex(reflect.ValueOf(field.name))
		ok := val.IsValid()
		if !ok || field.NoEncode {
			continue
		}

		if err := c.encodeValue(val); err != nil {
			return errors.Wrapf(err, "encoding %q", field.name)
		}
	}

	return nil
}

func (c *Encoder) encodeVector(slice reflect.Value) error {
	if err := c.PutCRC(CrcVector); err != nil {
		return err
	}
	if err := c.PutUint(uint32(slice.Len())); err != nil {
		return err
	}

	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		err := c.encodeValue(item)
		if err != nil {
			return errors.Wrapf(err, "[%v]", i)
		}
	}

	return nil
}


// v is bool, pointer, interface or slice
func isFieldContainsData(v reflect.Value) bool {
	// special cases for enums and objects
	if v.Type().Implements(enumTyp) {
		return v.Interface().(Enum).CRC() != 0
	}

	switch k := v.Kind(); k { // nolint:exhaustive // have default
	case reflect.Bool:
		return v.Bool()

	case reflect.Pointer, reflect.Interface, reflect.Slice:
		return !v.IsNil()

	default:
		panic(
			fmt.Sprintf(
				"something is wrong with registry, optional field must be pointer or bool, got %v",
				k,
			),
		)
	}
}
