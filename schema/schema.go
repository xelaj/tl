// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

import (
	"fmt"
	"strings"

	"github.com/quenbyako/ext/slices"
)

type Schema struct {
	TypeOrder []ObjName
	// key is type name
	Objects map[ObjName]TypeObjects
	// key is enum name
	Enums map[ObjName]EnumObjects

	MethodGroupOrder []string
	MethodsGroups    map[string][]Object // methods must be sorted by name
}

func (s *Schema) String() string {
	var parts []string
	for _, typ := range s.TypeOrder {
		if obj, ok := s.Objects[typ]; ok {
			parts = append(parts, obj.String())
		} else if enum, ok := s.Enums[typ]; ok {
			parts = append(parts, enum.String())
		} else {
			panic(fmt.Sprintf("missed type %#v", typ))
		}
	}

	if len(s.MethodGroupOrder) == 0 {
		return strings.Join(parts, "\n\n") + "\n"
	}

	parts = append(parts, "---functions---")

	for _, group := range s.MethodGroupOrder {
		obj, ok := s.MethodsGroups[group]
		if !ok {
			panic(fmt.Sprintf("missed group %#v", group))
		}

		parts = append(parts, methodsString(obj, group))
	}

	return strings.Join(parts, "\n\n") + "\n"
}

type CRCIndex struct {
	Type        ObjName
	ObjectIndex int
}

func (s *Schema) MakeCRCIndex() map[uint32]CRCIndex {
	res := make(map[uint32]CRCIndex, len(s.Objects))
	for typ, obj := range s.Objects {
		for i, o := range obj.Objects {
			res[o.CRC] = CRCIndex{
				Type:        typ,
				ObjectIndex: i,
			}
		}
	}

	return res
}

type TypeObjects struct {
	Comment string
	Objects []Object // must be sorted by name
}

func (s TypeObjects) String() string {
	var parts []string
	if s.Comment != "" {
		parts = append(parts, "// @type "+s.Comment)
	}

	for _, obj := range slices.SortFunc(s.Objects, sortObject) {
		parts = append(parts, obj.Comments(ObjectTypeConstructor)...)
		parts = append(parts, obj.String())
	}

	return strings.Join(parts, "\n")
}

type EnumObjects struct {
	Comment string
	Objects []Object // must be sorted by name
}

func (s EnumObjects) String() (res string) {
	var parts []string
	if s.Comment != "" {
		parts = append(parts, "// @type "+s.Comment)
	}

	for _, obj := range slices.SortFunc(s.Objects, sortObject) {
		parts = append(parts, obj.Comments(ObjectTypeEnum)...)
		parts = append(parts, obj.String())
	}

	return strings.Join(parts, "\n")
}

func methodsString(methods []Object, group string) (res string) {
	var parts []string

	if group != "" {
		group += "."
	}

	for _, obj := range slices.SortFunc(methods, sortObject) {
		parts = append(parts, obj.Comments(ObjectTypeMethod)...)
		parts = append(parts, obj.String())
	}

	return strings.Join(parts, "\n")
}
