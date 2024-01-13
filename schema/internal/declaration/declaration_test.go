// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package declaration_test

import (
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"

	. "github.com/xelaj/tl/schema/internal/declaration"
	"github.com/xelaj/tl/schema/internal/lexer"
)

type TestCase interface {
	Name() string
	Run(t *testing.T)
}

type TcaseParseIt[T any] struct {
	name    string
	in      string
	want    T
	wantErr assert.ErrorAssertionFunc
}

func (tt TcaseParseIt[T]) Name() string { return tt.name }

func (tt TcaseParseIt[T]) Run(t *testing.T) {
	tt.wantErr = noErrAsDefault(tt.wantErr)

	parser := participle.MustBuild[T](
		participle.Lexer(lexer.NewDefinition()),
	)

	got, err := parser.ParseString("test.tl", tt.in)
	if !tt.wantErr(t, err) || err != nil {
		return
	}

	assert.Equal(t, tt.want, *got)
}

func TestParseIt(t *testing.T) {
	for _, tt := range []TestCase{
		TcaseParseIt[Argument]{
			name: "args",
			in:   "user_id:long",
			want: Argument{
				Ident: VarIdentOpt{
					Ident: &VarIdent{Value: "user_id"},
				},
				Term: Ident{
					Simple: SimpleIdent{Var: &VarIdent{Value: "long"}},
				},
			},
		},
		TcaseParseIt[CombinatorDecl]{
			name: "declaration",
			in:   "a#00000000 flags:# b_x:c_x d:e pipka:flags.2?Vector<popka> = F",
			want: CombinatorDecl{
				ID: "a#00000000",
				Args: []Argument{
					{
						Ident: VarIdentOpt{
							Ident: &VarIdent{Value: "flags"},
						},
						Term: Ident{
							Simple: SimpleIdent{Type: &TypeIdent{Empty: true}},
						},
					},
					{
						Ident: VarIdentOpt{
							Ident: &VarIdent{Value: "b_x"},
						},
						Term: Ident{
							Simple: SimpleIdent{Var: &VarIdent{Value: "c_x"}},
						},
					},
					{
						Ident: VarIdentOpt{
							Ident: &VarIdent{Value: "d"},
						},
						Term: Ident{
							Simple: SimpleIdent{Var: &VarIdent{Value: "e"}},
						},
					},
					{
						Ident: VarIdentOpt{
							Ident: &VarIdent{Value: "pipka"},
						},
						Conditional: &ConditionalDef{
							Ident: VarIdent{Value: "flags"},
							Index: 2,
						},
						Term: Ident{
							Simple: SimpleIdent{Var: &VarIdent{Value: "Vector"}},
							Extension: &IdentExtension{
								Inner: []SimpleIdent{{Var: &VarIdent{Value: "popka"}}},
							},
						},
					},
				},
				Result: ResultType{
					Simple: stringPtr("F"),
				},
			},
		},
		TcaseParseIt[Program]{
			name: "program",
			in: `
a#123 = B;
c#456 = D;
---functions---

// another comment

// comment
e#789 = F;
g#012 = H;
`,
			want: Program{
				Constraints: []ProgramEntry{
					{Newline: true},
					{Decl: &CombinatorDecl{
						ID:     "a#123",
						Result: ResultType{Simple: stringPtr("B")},
					}},
					{Decl: &CombinatorDecl{
						ID:     "c#456",
						Result: ResultType{Simple: stringPtr("D")},
					}},
				},
				Methods: []ProgramEntry{
					{Newline: true},
					{Comment: stringPtr("// another comment")},
					{Newline: true},
					{Comment: stringPtr("// comment")},
					{Decl: &CombinatorDecl{
						ID:     "e#789",
						Result: ResultType{Simple: stringPtr("F")},
					}},
					{Decl: &CombinatorDecl{
						ID:     "g#012",
						Result: ResultType{Simple: stringPtr("H")},
					}},
				},
			},
		},
	} {
		t.Run(tt.Name(), tt.Run)
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}

func stringPtr(s string) *string { return &s }
