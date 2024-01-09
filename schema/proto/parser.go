// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package proto

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

//nolint:gochecknoglobals // obviously parser must be global
var parser = participle.MustBuild[Proto](participle.UseLookahead(2))

type Proto struct {
	Pos lexer.Position

	Entries []*Entry `parser:"( @@ ';'* )*"`
}

type Entry struct {
	Pos lexer.Position

	Syntax  string   `parser:"'syntax' '=' @String"`
	Package string   `parser:"| 'package' @(Ident ( '.' Ident )*)"`
	Import  string   `parser:"| 'import' @String"`
	Message *Message `parser:"| @@"`
	Service *Service `parser:"| @@"`
	Enum    *Enum    `parser:"| @@"`
	Option  *Option  `parser:"| 'option' @@"`
	Extend  *Extend  `parser:"| @@"`
}

type Option struct {
	Pos lexer.Position

	Name  string  `parser:"( '(' @Ident @( '.' Ident )* ')' | @Ident @( '.' @Ident )* )"`
	Attr  *string `parser:"( '.' @Ident ( '.' @Ident )* )?"`
	Value *Value  `parser:"'=' @@"`
}

type Value struct {
	Pos lexer.Position

	String    *string  `parser:"  @String"`
	Number    *float64 `parser:"| @Float"`
	Int       *int64   `parser:"| @Int"`
	Bool      *bool    `parser:"| (@'true' | 'false')"`
	Reference *string  `parser:"| @Ident @( '.' Ident )*"`
	Map       *Map     `parser:"| @@"`
	Array     *Array   `parser:"| @@"`
}

type Array struct {
	Pos lexer.Position

	Elements []*Value `parser:"'[' ( @@ ( ','? @@ )* )? ']'"`
}

type Map struct {
	Pos lexer.Position

	Entries []*MapEntry `parser:"'{' ( @@ ( ( ',' )? @@ )* )? '}'"`
}

type MapEntry struct {
	Pos lexer.Position

	Key   *Value `parser:"@@"`
	Value *Value `parser:"':'? @@"`
}

type Extensions struct {
	Pos lexer.Position

	Extensions []Range `parser:"'extensions' @@ ( ',' @@ )*"`
}

type Reserved struct {
	Pos lexer.Position

	Reserved []Range `parser:"'reserved' @@ ( ',' @@ )*"`
}

type Range struct {
	Ident string `parser:"  @String"`
	Start int    `parser:"| ( @Int"`
	End   *int   `parser:"  ( 'to' ( @Int"`
	Max   bool   `parser:"           | @'max' ) )? )"`
}

type Extend struct {
	Pos lexer.Position

	Reference string   `parser:"'extend' @Ident ( '.' @Ident )*"`
	Fields    []*Field `parser:"'{' ( @@ ';'? )* '}'"`
}

type Service struct {
	Pos lexer.Position

	Name  string          `parser:"'service' @Ident"`
	Entry []*ServiceEntry `parser:"'{' ( @@ ';'? )* '}'"`
}

type ServiceEntry struct {
	Pos lexer.Position

	Option *Option `parser:"  'option' @@"`
	Method *Method `parser:"| @@"`
}

type Method struct {
	Pos lexer.Position

	Name              string    `parser:"'rpc' @Ident"`
	StreamingRequest  bool      `parser:"'(' @'stream'?"`
	Request           *Type     `parser:"    @@ ')'"`
	StreamingResponse bool      `parser:"'returns' '(' @'stream'?"`
	Response          *Type     `parser:"              @@ ')'"`
	Options           []*Option `parser:"( '{' ( 'option' @@ ';' )* '}' )?"`
}

type Enum struct {
	Pos lexer.Position

	Name   string       `parser:"'enum' @Ident"`
	Values []*EnumEntry `parser:"'{' ( @@ ( ';' )* )* '}'"`
}

type EnumEntry struct {
	Pos lexer.Position

	Value  *EnumValue `parser:"  @@"`
	Option *Option    `parser:"| 'option' @@"`
}

type EnumValue struct {
	Pos lexer.Position

	Key   string `parser:"@Ident"`
	Value int    `parser:"'=' @( [ '-' ] Int )"`

	Options []*Option `parser:"( '[' @@ ( ',' @@ )* ']' )?"`
}

type Message struct {
	Pos lexer.Position

	Name    string          `parser:"'message' @Ident"`
	Entries []*MessageEntry `parser:"'{' @@* '}'"`
}

type MessageEntry struct {
	Pos lexer.Position

	Enum       *Enum       `parser:"( @@"`
	Option     *Option     `parser:" | 'option' @@"`
	Message    *Message    `parser:" | @@"`
	Oneof      *Oneof      `parser:" | @@"`
	Extend     *Extend     `parser:" | @@"`
	Reserved   *Reserved   `parser:" | @@"`
	Extensions *Extensions `parser:" | @@"`
	Field      *Field      `parser:" | @@ ) ';'*"`
}

type Oneof struct {
	Pos lexer.Position

	Name    string        `parser:"'oneof' @Ident"`
	Entries []*OneofEntry `parser:"'{' ( @@ ';'* )* '}'"`
}

type OneofEntry struct {
	Pos lexer.Position

	Field  *Field  `parser:"  @@"`
	Option *Option `parser:"| 'option' @@"`
}

type Field struct {
	Pos lexer.Position

	Optional bool `parser:"(   @'optional'"`
	Required bool `parser:"  | @'required'"`
	Repeated bool `parser:"  | @'repeated' )?"`

	Type *Type  `parser:"@@"`
	Name string `parser:"@Ident"`
	Tag  int    `parser:"'=' @Int"`

	Options []*Option `parser:"( '[' @@ ( ',' @@ )* ']' )?"`
}

type Scalar int

const (
	None Scalar = iota
	Double
	Float
	Int32
	Int64
	Uint32
	Uint64
	Sint32
	Sint64
	Fixed32
	Fixed64
	SFixed32
	SFixed64
	Bool
	String
	Bytes
)

var scalarToString = map[Scalar]string{
	None: "None", Double: "Double", Float: "Float", Int32: "Int32", Int64: "Int64", Uint32: "Uint32",
	Uint64: "Uint64", Sint32: "Sint32", Sint64: "Sint64", Fixed32: "Fixed32", Fixed64: "Fixed64",
	SFixed32: "SFixed32", SFixed64: "SFixed64", Bool: "Bool", String: "String", Bytes: "Bytes",
}

func (s Scalar) GoString() string { return scalarToString[s] }

var stringToScalar = map[string]Scalar{
	"double": Double, "float": Float, "int32": Int32, "int64": Int64, "uint32": Uint32, "uint64": Uint64,
	"sint32": Sint32, "sint64": Sint64, "fixed32": Fixed32, "fixed64": Fixed64, "sfixed32": SFixed32,
	"sfixed64": SFixed64, "bool": Bool, "string": String, "bytes": Bytes,
}

func (s *Scalar) Parse(lex *lexer.PeekingLexer) error {
	token := lex.Peek()
	v, ok := stringToScalar[token.Value]
	if !ok {
		return participle.NextMatch
	}
	lex.Next()
	*s = v
	return nil
}

type Type struct {
	Pos lexer.Position

	Scalar    Scalar   `parser:"  @@"`
	Map       *MapType `parser:"| @@"`
	Reference string   `parser:"| @(Ident ( '.' Ident )*)"`
}

type MapType struct {
	Pos lexer.Position

	Key   *Type `parser:"'map' '<' @@"`
	Value *Type `parser:"',' @@ '>'"`
}
