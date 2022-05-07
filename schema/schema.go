// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

import (
	"fmt"
	"hash/crc32"
	"strings"
)

type Schema struct { //nolint:govet // it's ok
	Objects      []Object
	Methods      []Object
	TypeComments map[string]string
}

type Object struct { //nolint:govet // it's ok
	Comment string

	Name       string
	CRC        uint32
	Parameters []Parameter
	Type       Type
}

func (o *Object) GenerateCRC() uint32 {
	return crc32.ChecksumIEEE([]byte(o.StringNoCRC()))
}

func (o *Object) StringNoCRC() string {
	params := make([]string, len(o.Parameters))
	for i, item := range o.Parameters {
		params[i] = item.String()
	}

	return fmt.Sprintf("%v %v = %v;", o.Name, strings.Join(params, " "), o.Type)
}

func (o *Object) String() string {
	params := make([]string, len(o.Parameters))
	for i, item := range o.Parameters {
		params[i] = item.String()
	}

	crc := o.CRC
	if crc == 0 {
		crc = o.GenerateCRC()
	}

	return fmt.Sprintf("%v#%08x %v = %v;", o.Name, crc, strings.Join(params, " "), o.Type)
}

type Parameter interface {
	fmt.Stringer
	GetComment() string

	_isParameter()
}

type RequiredParameter struct {
	Type    Type
	Comment string
	Name    string
}

func (RequiredParameter) _isParameter() {}

func (t RequiredParameter) String() string {
	return t.Name + ":" + t.Type.String()
}

func (t RequiredParameter) GetComment() string {
	return t.Comment
}

type OptionalParameter struct {
	Comment     string
	Name        string
	Type        Type
	FlagTrigger string
	BitTrigger  int
}

func (OptionalParameter) _isParameter() {}

func (t OptionalParameter) String() string {
	return fmt.Sprintf("%v:%v.%v?%v", t.Name, t.FlagTrigger, t.BitTrigger, t.Type.String())
}

func (t OptionalParameter) GetComment() string {
	return t.Comment
}

type TriggerParameter struct {
	Comment     string
	Name        string
	FlagTrigger string
	BitTrigger  int
}

func (TriggerParameter) _isParameter() {}

func (t TriggerParameter) String() string {
	return fmt.Sprintf("%v:%v.%v?true", t.Name, t.FlagTrigger, t.BitTrigger)
}

func (t TriggerParameter) GetComment() string {
	return t.Comment
}

type Type interface {
	fmt.Stringer

	_isType()
}

type TypeCommon string

func (TypeCommon) _isType() {}

func (t TypeCommon) String() string {
	return string(t)
}

type TypeVector string

func (TypeVector) _isType() {}

func (t TypeVector) String() string {
	return "Vector<" + string(t) + ">"
}
