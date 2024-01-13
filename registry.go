// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"cmp"
	"fmt"
	"reflect"

	"github.com/quenbyako/ext/slices"
)

// key is crc code, value is name of constructor.
type (
	enumNames = map[crc32]string
	typeName  = string
)

type Registry interface {
	ConstructObject(code crc32) (Object, bool)
}

var _defaultRegistry = val(NewRegistry())

func DefaultRegistry() *ObjectRegistry { return &_defaultRegistry }

func RegisterObjectDefault[T Object]()         { RegisterObject[T](&_defaultRegistry) }
func RegisterEnumDefault[T Object](enums ...T) { RegisterEnum[T](&_defaultRegistry, enums...) }
func RegisterCustomDefault[T Object](constructor func(uint32) T, crcs ...uint32) {
	RegisterCustom[T](&_defaultRegistry, constructor, crcs...)
}

// ObjectRegistry is a type, which handles code generated schema, and could be
// useful for spawning TL objects. Unlike RawSchemaRegistry, it can work only
// with predefined go types.
//
// If you are not able to use codegen, use RawSchemaRegistry, it could be
// slower, but more flexible.
type ObjectRegistry struct {
	// in objects it's allowed to store ONLY structs and uint32.
	objects map[crc32]func(uint32) Object
}

var _ Registry = (*ObjectRegistry)(nil)

func NewRegistry() *ObjectRegistry {
	return &ObjectRegistry{
		objects: make(map[crc32]func(uint32) Object),
	}
}

// ConstructObject spawns new object by crc code from registry.
func (r *ObjectRegistry) ConstructObject(crc crc32) (Object, bool) {
	if obj, ok := r.objects[crc]; ok {
		return obj(crc), true
	}

	return nil, false
}

type field struct {
	fType       fieldType
	name        string
	flagTrigger uint8 // index of field in list of fields in orphanType
	bitTrigger  uint8
	noEncode    bool // don't encode value, if value is true. works ONLY for boolean types
	optional    bool
}

type structFields struct {
	// key is an index of optional field in list of fields in object
	// value is bit, which you need to trigger
	bitflags map[int]BitflagBit
	tags     []StructTag
}

func (s *structFields) isFieldOptional(fieldIndex int) bool {
	_, ok := s.bitflags[fieldIndex]

	return ok
}

type BitflagBit struct {
	FieldIndex int // в какое поле пихать бит что поле существует
	BitIndex   int // собственно в какой бит пихать флаг что все ок
}

func RegisterObject[T Object](r *ObjectRegistry) {
	r.registerObject(asObject(new[T]))
}

func RegisterEnum[T Object](r *ObjectRegistry, enums ...T) {
	if t := new[T](); reflect.TypeOf(t).Kind() != reflect.Uint32 {
		panic("enums must be uint32")
	}

	if len(enums) == 0 {
		panic("no enums provided")
	}

	enums = slices.SortFunc(enums, func(a, b T) int { return cmp.Compare(a.CRC(), b.CRC()) })

	f := func(crc uint32) (res T, ok bool) {
		if i, ok := slices.BinarySearchFunc(enums, crc, func(enum T, crc uint32) int {
			return cmp.Compare(enum.CRC(), crc)
		}); ok {
			return enums[i], true
		}

		return res, false
	}

	r.registerCustom(slices.Remap(enums, func(enum T) uint32 { return enum.CRC() }), asEnum(f))
}

func RegisterCustom[T Object](r *ObjectRegistry, constructor func(uint32) T, crcs ...uint32) {
	if len(crcs) == 0 {
		panic("no enums provided")
	}
	r.registerCustom(crcs, func(u uint32) Object { return constructor(u) })
}

func (r *ObjectRegistry) registerObject(c func(uint32) Object) {
	typ := reflect.TypeOf(c(0))
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		panic("accepted only structs")
	}

	if r.objects == nil {
		r.objects = make(map[crc32]func(uint32) Object)
	}
	crc := c(0).CRC()
	if _, ok := r.objects[crc]; ok {
		panic(fmt.Sprintf("object with code %x already registered", crc))
	}

	r.objects[crc] = c
}

func (r *ObjectRegistry) registerCustom(valid []uint32, c func(uint32) Object) {
	if r.objects == nil {
		r.objects = make(map[crc32]func(uint32) Object)
	}
	for _, enum := range valid {
		r.objects[enum] = c
	}
}
