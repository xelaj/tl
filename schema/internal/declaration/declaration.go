// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package declaration

import (
	"strings"
)

type Program struct {
	Constraints []ProgramEntry `parser:"@@*"`
	Methods     []ProgramEntry `parser:"(functions_decl newline @@*)?"`
}

type ProgramEntry struct {
	Newline bool            `parser:"@newline |"`
	Decl    *CombinatorDecl `parser:"@@ ws? semicolon ( newline | EOF ) |"`
	Comment *string         `parser:"@comment ( newline | EOF ) "`
}

type CombinatorDecl struct {
	ID     string     `parser:"@lc_ident_full ws"`
	Args   []Argument `parser:"(@@ (ws @@)* ws)? equals ws?"`
	Result ResultType `parser:"@@"`
}

type ResultType struct {
	Simple *string         `parser:"( @uc_ident_ns | @uc_ident )"`
	Expr   *IdentExtension `parser:"@@?"`
}

type Argument struct {
	Ident       VarIdentOpt     `parser:"@@ colon"`
	Conditional *ConditionalDef `parser:"@@?"`
	Term        Ident           `parser:"@@"`
}

type Ident struct {
	Simple    SimpleIdent     `parser:"@@"`
	Extension *IdentExtension `parser:"@@?"`
}

func (i *Ident) String() string {
	res := i.Simple.String()
	if i.Extension != nil {
		res += i.Extension.String()
	}
	return res
}

type SimpleIdent struct {
	Var  *VarIdent  `parser:"@@ |"`
	Type *TypeIdent `parser:"@@"`
}

func (s *SimpleIdent) String() string {
	switch {
	case s.Var != nil:
		return s.Var.Value
	case s.Type != nil:
		return s.Type.String()
	default:
		panic("impossible situation")
	}
}

type VarIdent struct {
	Value string `parser:"@lc_ident | @uc_ident"`
}

type TypeIdent struct {
	Ident *string `parser:"@uc_ident_ns | @lc_ident_ns |"`
	Empty bool    `parser:"@hash"`
}

func (t *TypeIdent) String() string {
	if t.Empty {
		return "#"
	}

	return *t.Ident
}

type IdentExtension struct {
	Inner []SimpleIdent `parser:"langle @@ (ws @@)* rangle"`
}

func (i *IdentExtension) String() string {
	items := make([]string, len(i.Inner))
	for i, item := range i.Inner {
		items[i] = item.String()
	}

	return "<" + strings.Join(items, " ") + ">"
}

type VarIdentOpt struct {
	Ident *VarIdent `parser:"@@ |"`
	Empty bool      `parser:"@underscore"`
}

func (i *VarIdentOpt) String() string {
	if i.Empty {
		return "_"
	}

	return i.Ident.Value
}

type ConditionalDef struct {
	Ident VarIdent `parser:"@@"`
	Index int      `parser:"dot @nat_const question_mark"`
}
