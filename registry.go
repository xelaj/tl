// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"fmt"
	"reflect"
)

type enumNames = map[crc32]string

type ObjectRegistry struct {
	// in objects it's allowed to store ONLY structs
	objects      map[crc32]reflect.Type
	structFields map[crc32]structFields
	// enumStrings are specific case unlike objects: we can already predict enum
	// type (cause enum is an interface, which are implemented only in objects
	// without arguments). It's guaranteed that we know enum type, so no need to
	enums     map[crc32]Enum
	enumNames enumNames

	orphans map[crc32]orphanType
}

func (r *ObjectRegistry) pushObject(crc crc32, typ reflect.Type) {
	if typ.Kind() != reflect.Struct {
		panic("accepted only structs")
	}

	if r.objects == nil {
		r.objects = make(map[crc32]reflect.Type)
	}
	r.objects[crc] = typ
}

// spawnObject spawns new object by crc code from registry.
func (r *ObjectRegistry) spawnObject(crc crc32) (reflect.Value, error) {
	_type, ok := r.objects[crc]
	if !ok {
		return reflect.Value{}, fmt.Errorf("object with crc 0x%08x not found", crc)
	}

	v := reflect.New(_type).Elem()
	if !v.Type().Implements(objectTyp) {
		v = v.Addr()
	}
	if !v.Type().Implements(objectTyp) {
		panic("internal error: " +
			"object neither implements object as raw struct, nor pointer " + v.Type().String())
	}

	return v, nil
}

// orphanType is a type definition, which doesn't have any struct type to it.
type orphanType struct {
	implements string
	fields     []field
}

// collectBitflags collects bitflags from map.
func (o *orphanType) collectBitflags(m reflect.Value) (map[uint8]crc32, error) {
	if m.Kind() != reflect.Map {
		panic("provided non map value to method")
	}

	f := map[uint8]crc32{}

	for i, field := range o.fields {
		if _, ok := field.fType.(fieldBitflag); ok {
			f[uint8(i)] = 0

			continue
		}

		ok := m.MapIndex(reflect.ValueOf(field.name)).IsValid()
		if field.optional && (ok || field.noEncode) {
			f[field.flagTrigger] |= crc32(1 << field.bitTrigger)
		} else if !field.optional && !ok {
			//nolint:goerr113 // it's an internal error
			return nil, fmt.Errorf("field %q is required for this type", field.name)
		}
	}

	return f, nil
}

type field struct {
	fType       fieldType
	name        string
	flagTrigger uint8 // index of field in list of fields in orphanType
	bitTrigger  uint8
	noEncode    bool // don't encode value, if value is true. works ONLY for boolean types
	optional    bool
}

type fieldType interface {
	_isFieldType()
}

type fieldBitflag null

func (fieldBitflag) _isFieldType() {}

type fieldObject crc32

func (fieldObject) _isFieldType() {}

type fieldInterface string

func (fieldInterface) _isFieldType() {}

//nolint:gochecknoglobals // obvious reason to fo that.
var defaultRegistry = &ObjectRegistry{}

func RegisterObjects(objects ...Object) { defaultRegistry.RegisterObjects(objects...) }
func RegisterEnums(enums ...Enum)       { defaultRegistry.RegisterEnums(enums...) }

type structFields struct {
	// key is an index of optional field in list of fields in object
	// value is bit, which you need to trigger
	bitflags map[int]bitflagBit
	tags     []structTag
}

func (s *structFields) isFieldOptional(fieldIndex int) bool {
	_, ok := s.bitflags[fieldIndex]

	return ok
}

type bitflagBit struct {
	fieldIndex int // в какое поле пихать бит что поле существует
	bitIndex   int // собственно в какой бит пихать флаг что все ок
}

func (r *ObjectRegistry) RegisterObjects(objects ...Object) {
	for _, object := range objects {
		r.registerObject(object)
	}
}

func (r *ObjectRegistry) registerObject(o Object) {
	if o == nil {
		panic("object is nil")
	}

	typ := reflect.TypeOf(o)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	typData := structFields{
		tags:     make([]structTag, typ.NumField()),
		bitflags: make(map[int]bitflagBit),
	}

	tagNamesIndexes := make(map[string]int, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fTyp := typ.Field(i)
		typName := typ.Name() + "." + fTyp.Name

		tag, err := parseTag(fTyp.Tag.Get(tagName), fTyp.Name)
		if err != nil {
			panic(fmt.Sprintf("parsing tag of %v: %v", typName, err.Error()))
		}
		tagNamesIndexes[tag.Name] = i

		if tag.BitFlags != nil {
			targetField, ok := tagNamesIndexes[tag.BitFlags.TargetField]
			if !ok {
				panic(fmt.Sprintf("%v: field is optional, but bitflags indicating that this "+
					"field became after exact field", typName))
			}

			switch fTyp.Type.Kind() { //nolint:exhaustive // passes 4 kinds, otherwise panics
			case reflect.Bool, reflect.Ptr, reflect.Interface, reflect.Slice:
			default:
				// enum is specific case, 0 value is always null
				if fTyp.Type.Implements(enumTyp) {
					break
				}
				panic(fmt.Sprintf("%v: field tagged as omitempty, but kind is not pointer to "+
					"value or bool", typName))
			}

			typData.bitflags[i] = bitflagBit{
				fieldIndex: targetField,
				bitIndex:   int(tag.BitFlags.BitPosition),
			}

			if tag.Implicit() && fTyp.Type.Kind() != reflect.Bool {
				panic(fmt.Sprintf("%v: %q tag works only for bool fields", typName, implicitFlag))
			}
		}

		typData.tags[i] = *tag
	}

	if r.structFields == nil {
		r.structFields = make(map[crc32]structFields)
	}
	r.structFields[o.CRC()] = typData
	r.pushObject(o.CRC(), typ)
}

func (r *ObjectRegistry) RegisterEnums(enums ...Enum) {
	for _, enum := range enums {
		r.registerEnum(enum)
	}
}

func (r *ObjectRegistry) registerEnum(enum Enum) {
	if enum == nil {
		panic("enum is nil")
	}

	if r.enums == nil {
		r.enums = make(map[uint32]Enum, 1)
	}
	r.enums[enum.CRC()] = enum

	if r.enumNames == nil {
		r.enumNames = make(map[uint32]string, 1)
	}
	r.enumNames[enum.CRC()] = enum.String()
}
