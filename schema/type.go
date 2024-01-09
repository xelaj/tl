package schema

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type Type interface {
	_type()

	fmt.Stringer
}

var (
	_ Type = TypeCommon(ObjName{})
	_ Type = TypeVector(ObjName{})
)

type TypeCommon ObjName

func (_ TypeCommon) _type()         {}
func (t TypeCommon) String() string { return ObjName(t).String() }

type TypeVector ObjName

func (_ TypeVector) _type()         {}
func (t TypeVector) String() string { return "Vector<" + ObjName(t).String() + ">" }

func isFirstRuneUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}
