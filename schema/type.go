package schema

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type Type interface {
	_isType()

	fmt.Stringer
	IsInterface() bool
}

var (
	_ Type = TypeCommon("")
	_ Type = TypeVector("")
)

type TypeCommon string

func (_ TypeCommon) _isType()          {}
func (t TypeCommon) String() string    { return string(t) }
func (t TypeCommon) IsInterface() bool { return !isFirstRuneUpper(string(t)) }

type TypeVector string

func (_ TypeVector) _isType()          {}
func (t TypeVector) String() string    { return "Vector<" + string(t) + ">" }
func (t TypeVector) IsInterface() bool { return true }

func isFirstRuneUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}
