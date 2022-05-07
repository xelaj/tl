// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"fmt"
	"reflect"
)

type enumNames = map[crc32]string

type Registry struct {
	// in objects it's allowed to store ONLY structs
	objects      map[crc32]reflect.Type
	structFields map[crc32]structFields
	// enumStrings are specific case unlike objects: we can already predict enum type (cause enum is an interface,
	// which are implemented only in objects without arguments). It's guaranteed that we know enum type, so no need to
	enums     map[crc32]Enum
	enumNames enumNames

	orphans map[crc32]orphanType
}

func (r *Registry) pushObject(crc crc32, typ reflect.Type) {
	if typ.Kind() != reflect.Struct {
		panic("accepted only structs")
	}

	if r.objects == nil {
		r.objects = make(map[crc32]reflect.Type)
	}
	r.objects[crc] = typ
}

// spawnObject spawns new object by crc code from registry
func (r *Registry) spawnObject(crc crc32) (reflect.Value, error) {
	_type, ok := r.objects[crc]
	if !ok {
		return reflect.Value{}, fmt.Errorf("object with crc 0x%08x not found", crc)
	}

	v := reflect.New(_type).Elem()
	if !v.Type().Implements(objectTyp) {
		v = v.Addr()
	}
	if !v.Type().Implements(objectTyp) {
		panic("internal error: object neither implements object as raw struct, nor pointer " + v.Type().String())
	}
	return v, nil
}

type interfaceImplementations map[string]map[crc32]null

// orphanType is a type definition, which doesn't have any struct type to it.
type orphanType struct {
	fields     []field
	implements string
}

type field struct {
	name        string
	fType       fieldType
	optional    bool
	FlagTrigger uint8 // index of field in list of fields in orphanType
	BitTrigger  uint8
	NoEncode    bool // don't encode value, if value is true. works ONLY for boolean types
}

type fieldType interface {
	_isFieldType()
}

type fieldBitflag struct{}

func (fieldBitflag) _isFieldType() {}

type fieldObject crc32

func (fieldObject) _isFieldType() {}

type fieldInterface string

func (fieldInterface) _isFieldType() {}

//nolint:gochecknoglobals // obvious reason to fo that.
var defaultRegistry = &Registry{}

// зачем регистрация объектов:
// для энкдоинга это не так важно, здесь проблем никаких нет. Регистр хоть и используется, но лишь для оптимизации, что бы не парсить 300 раз теги
//
// c декодированием обратная история: каждый объект имеет crc код, которым определяется его тип (примерно так
// же как устроены интерфейсы в golang), конечно можно было бы использовать рефлексию, что бы предугадывать
// тип, но в этом случае все равно бы пришлось делать отдельный регистр для каждого интерфейса, что менее удобно.
// gj'nj

func RegisterObjects(objects ...Object) { defaultRegistry.RegisterObjects(objects...) }
func RegisterEnums(enums ...Enum)       { defaultRegistry.RegisterEnums(enums...) }

type structFields struct {
	tags     []structTag
	bitflags map[int]bitflagBit // ключ это индекс опционального поля, значение это во флаг под каким индексом его пихать
}

func (s *structFields) isFieldOptional(fieldIndex int) bool {
	_, ok := s.bitflags[fieldIndex]
	return ok
}

type bitflagBit struct {
	fieldIndex int // в какое поле пихать бит что поле существует
	bitIndex   int // собственно в какой бит пихать флаг что все ок
}

func (r *Registry) RegisterObjects(objects ...Object) {
	for _, object := range objects {
		r.registerObject(object)
	}
}

func (r *Registry) registerObject(o Object) {
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
				panic(fmt.Sprintf("%v: field is optional, but bitflags indicating that this field became after exact field", typName))
			}

			switch fTyp.Type.Kind() {
			case reflect.Bool, reflect.Ptr, reflect.Interface, reflect.Slice:
			default:
				// enum is specific case, 0 value is always null
				if fTyp.Type.Implements(enumTyp) {
					break
				}
				panic(fmt.Sprintf("%v: field tagged as omitempty, but kind is not pointer to value or bool", typName))
			}

			typData.bitflags[i] = bitflagBit{
				fieldIndex: targetField,
				bitIndex:   int(tag.BitFlags.BitPosition),
			}

			if tag.Implicit() && fTyp.Type.Kind() != reflect.Bool {
				panic(fmt.Sprintf("%v: %q tag works only for boolean fields", typName, implicitFlag))
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

func (r *Registry) RegisterEnums(enums ...Enum) {
	for _, enum := range enums {
		r.registerEnum(enum)
	}
}

func (r *Registry) registerEnum(enum Enum) {
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
