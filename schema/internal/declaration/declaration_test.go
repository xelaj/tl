// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package declaration_test

import (
	"reflect"
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"

	. "github.com/xelaj/tl/schema/internal/declaration"
	"github.com/xelaj/tl/schema/internal/lexer"
)

func TestParseIt(t *testing.T) {
	for _, tt := range []struct {
		name    string
		in      string
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "args",
			in:   "user_id:long",
			want: &Argument{
				Ident: VarIdentOpt{
					Ident: &VarIdent{Value: "user_id"},
				},
				Term: Ident{
					Simple: SimpleIdent{Var: &VarIdent{Value: "long"}},
				},
			},
		},
		{
			name: "declaration",
			in:   "a#00000000 flags:# b_x:c_x d:e pipka:flags.2?Vector<popka> = F",
			want: &CombinatorDecl{
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
		{
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
			want: &Program{
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
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			parser := participle.MustBuild(
				tt.want,
				participle.Lexer(lexer.NewDefinition()),
			)

			got := reflect.New(reflect.TypeOf(tt.want).Elem()).Interface()

			err := parser.ParseString("test.tl", tt.in, got)

			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}

func stringPtr(s string) *string { return &s }
