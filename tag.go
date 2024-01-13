// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// tag name of struct tags.
const (
	tagName = "tl"

	implicitFlag    = "implicit"
	isBitflagFlag   = "bitflag"
	omitemptyPrefix = "omitempty"

	maxBitflagIndex = 32
)

type StructTag struct {
	BitFlags *Bitflag
	Name     string
	Type     fieldType
}

// ParseTag is a function which parses struct field tag for structes, defined by
// you.
//
//   - `tag` parameter is a text representation of struct field tag.
//   - `defaultName` is a name of struct field, if it's not defined in tag.
//
// how tag could look:
//
//	// field named somefield, it could be empty, value is boolean, so it can
//	// be encoded implicitly, value stored in some_flag value
//	`tl:"somefield,omitempty:some_flag:3,implicit"`
//
//	// value is abcd, it's required to exist in serialized
//	`tl:"abcd"`
func ParseTag(tag, defaultName string) (StructTag, error) {
	return parseTag(tag, defaultName, nil)
}

// ft is optional, you can add related type later.
func parseTag(tag, defaultName string, ft reflect.Type) (t StructTag, err error) {
	parts := strings.Split(tag, ",")
	name := parts[0]
	if name == "" {
		name = defaultName
	}

	t.Name = name
	for _, option := range parts[1:] {
		switch {
		case option == implicitFlag:
			if t.Type != nil {
				return StructTag{}, ErrImplicitBitflag
			}
			t.Type = fieldImplicitBool{}

		case option == isBitflagFlag:
			if t.Type != nil {
				return StructTag{}, ErrImplicitBitflag
			}
			t.Type = fieldBitflags{}

		case strings.HasPrefix(option, omitemptyPrefix):
			if t.BitFlags, err = parseOmitemptyTag(option); err != nil {
				return StructTag{}, err
			}

		default:
			return StructTag{}, fmt.Errorf("%v: %w", option, ErrInvalidTagOption)
		}
	}
	if t.Type == nil && ft != nil {
		if t.Type = GetFieldTypeFromGoType(ft); t.Type == nil {
			return StructTag{}, ErrUnsupportedType{Type: ft}
		}
	}

	if err := t.valid(); err != nil {
		return StructTag{}, err
	}

	return t, nil
}

func ParseStructTags(s Object) ([]StructTag, map[int]BitflagBit, error) {
	t := indirectType(reflect.TypeOf(s))
	if t.Kind() != reflect.Struct {
		return nil, nil, errors.New("value is not a struct")
	}

	return parseStructTags(t)
}

func parseStructTags(t reflect.Type) ([]StructTag, map[int]BitflagBit, error) {
	tags := make([]StructTag, t.NumField())
	optional := map[int]BitflagBit{}
	tagNamesIndexes := make(map[string]int, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		typName := t.Name() + "." + ft.Name

		var err error
		if tags[i], err = ParseTag(ft.Tag.Get(tagName), ft.Name); err != nil {
			return nil, nil, fmt.Errorf("parsing tag of %v: %w", typName, err)
		}

		tagNamesIndexes[tags[i].Name] = i

		bitflags := tags[i].BitFlags
		if bitflags == nil {
			continue
		}

		targetField, ok := tagNamesIndexes[bitflags.TargetField]
		if !ok {
			return nil, nil, fmt.Errorf("tag %v as bitflag not found in struct", tags[i].BitFlags.TargetField)
		}

		switch ft.Type.Kind() { //nolint:exhaustive // passes 5 kinds, otherwise panics
		case reflect.Ptr, reflect.Interface, reflect.Slice:
		case reflect.Uint32:
			if ft.Type.ConvertibleTo(objectTyp) {
				break
			}

			fallthrough
		case reflect.Bool:
			if _, ok := tags[i].Type.(fieldImplicitBool); ok {
				break
			}

			fallthrough
		default:
			panic(fmt.Sprintf("%v: field tagged as omitempty, but kind is neither pointer to "+
				"value, nor enum, nor bool", typName))
		}

		optional[i] = BitflagBit{
			FieldIndex: targetField,
			BitIndex:   int(bitflags.BitPosition),
		}
	}

	return tags, optional, nil
}

func indirectType(t reflect.Type) reflect.Type { //cover:ignore
	if t.Kind() != reflect.Pointer {
		return t
	}
	return t.Elem()
}

func (t StructTag) Ignore() bool { return t.Name == "-" } //cover:ignore

func (t StructTag) String() string { //cover:ignore
	res := t.Name
	if t.BitFlags != nil {
		res += "," + t.BitFlags.String()
	}
	switch t.Type.(type) {
	case fieldImplicitBool:
		res += "," + implicitFlag
	case fieldBitflags:
		res += "," + isBitflagFlag
	}

	return res
}

func (t StructTag) valid() error {
	switch {
	case t.Name == "":
		return ErrTagNameEmpty
	case t.isImplicit() && (t.BitFlags == nil || t.BitFlags.TargetField == ""):
		return ErrImplicitNoTarget
	case t.BitFlags != nil && t.BitFlags.BitPosition > maxBitflagIndex:
		return ErrBitflagTooHigh
	default:
		return nil
	}
}

func (t StructTag) isImplicit() bool {
	_, ok := t.Type.(fieldImplicitBool)
	return ok
}

func (t StructTag) isBitflag() bool {
	_, ok := t.Type.(fieldBitflags)
	return ok
}

type Bitflag struct {
	TargetField string
	BitPosition uint8
}

func (t *Bitflag) String() string { //cover:ignore
	return strings.Join([]string{
		omitemptyPrefix, t.TargetField, strconv.Itoa(int(t.BitPosition)),
	}, ":")
}

const omitemptyParts = 3

func parseOmitemptyTag(opt string) (*Bitflag, error) {
	parts := strings.Split(opt, ":")

	if len(parts) != omitemptyParts {
		return nil, fmt.Errorf("%v: %w", omitemptyPrefix, ErrInvalidTagFormat)
	}

	pos, err := parseUintMax32(parts[2])
	if err != nil {
		return nil, fmt.Errorf("%v: %w", omitemptyPrefix, err)
	}

	return &Bitflag{
		TargetField: parts[1],
		BitPosition: pos,
	}, nil
}

const (
	bit32       = 5  // 5 bits to make 32 different variants
	defaultBase = 10 // base 10 of numbers
)

func parseUintMax32(s string) (uint8, error) {
	if pos, err := strconv.ParseUint(s, defaultBase, bit32); err == nil {
		return uint8(pos), nil
	}

	return 0, ErrInvalidTagFormat
}

func GetFieldTypeFromGoType(t reflect.Type) fieldType {
	if t.Implements(unmarshalerTyp) {
		return fieldCustom{}
	}

	switch t.Kind() {
	case reflect.Ptr:
		return GetFieldTypeFromGoType(t.Elem())

	case reflect.Uint32, reflect.Int32:
		return fieldInt{}

	case reflect.Int64:
		return fieldLong{}

	case reflect.Float64:
		return fieldDouble{}

	case reflect.Bool:
		return fieldBool{}

	case reflect.String:
		return fieldString{}

	case reflect.Struct:
		if crcGetter, ok := reflect.New(t).Interface().(Object); ok {
			return fieldObject(crcGetter.CRC())
		}

		// we can't support other types but Object.
		return nil

	case reflect.Interface:
		if t.Implements(objectTyp) {
			return fieldInterface{Type: t}
		}

		return nil

	case reflect.Map:
		return fieldInterface{Type: objectTyp} // any object we have

	case reflect.Slice:
		if inner := GetFieldTypeFromGoType(t.Elem()); inner != nil {
			return fieldVector{Inner: inner}
		}

		return nil

	default:
		return nil
	}
}

type fieldType interface{ _fieldType() }

var (
	BitflagType      = fieldBitflags{}
	ImplicitBoolType = fieldImplicitBool{}
)

type fieldBitflags null

func (fieldBitflags) _fieldType() {}

type fieldBool null

func (fieldBool) _fieldType() {}

type fieldImplicitBool null

func (fieldImplicitBool) _fieldType() {}

type fieldInt null

func (fieldInt) _fieldType() {}

type fieldLong null

func (fieldLong) _fieldType() {}

type fieldDouble null

func (fieldDouble) _fieldType() {}

type fieldInt128 null

func (fieldInt128) _fieldType() {}

type fieldInt256 null

func (fieldInt256) _fieldType() {}

type fieldBytes null

func (fieldBytes) _fieldType() {}

type fieldString null

func (fieldString) _fieldType() {}

type fieldVector struct {
	Inner fieldType
}

func (fieldVector) _fieldType() {}

type fieldObject crc32

func (fieldObject) _fieldType() {}

type fieldInterface struct {
	Type reflect.Type
}

func (fieldInterface) _fieldType() {}

type fieldCustom struct {
	Type reflect.Type
}

func (fieldCustom) _fieldType() {}

func typStructFields(typ reflect.Type) structFields {
	typData := structFields{
		tags:     make([]StructTag, typ.NumField()),
		bitflags: make(map[int]BitflagBit),
	}

	tagNamesIndexes := make(map[string]int, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fTyp := typ.Field(i)
		typName := typ.Name() + "." + fTyp.Name

		tag, err := ParseTag(fTyp.Tag.Get(tagName), fTyp.Name)
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
				if fTyp.Type.Implements(objectTyp) && fTyp.Type.Kind() == reflect.Uint32 {
					break
				}
				panic(fmt.Sprintf("%v: field tagged as omitempty, but kind is not pointer to "+
					"value or bool", typName))
			}

			typData.bitflags[i] = BitflagBit{
				FieldIndex: targetField,
				BitIndex:   int(tag.BitFlags.BitPosition),
			}

			if tag.isImplicit() && fTyp.Type.Kind() != reflect.Bool {
				panic(fmt.Sprintf("%v: %q tag works only for bool fields", typName, implicitFlag))
			}
		}

		typData.tags[i] = tag
	}

	return typData
}
