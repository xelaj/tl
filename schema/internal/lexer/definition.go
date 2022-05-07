// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package lexer

import (
	"context"
	"io"

	"github.com/alecthomas/participle/v2/lexer"
)

type Definition struct{}

func NewDefinition() lexer.Definition { return Definition{} }

func (Definition) Symbols() map[string]lexer.TokenType {
	return map[string]lexer.TokenType{
		"EOF":              lexer.EOF,
		"underscore":       TokenUnderscore,
		"colon":            TokenColon,
		"semicolon":        TokenSemicolon,
		"open_par":         TokenOpenPar,
		"close_par":        TokenClosePar,
		"open_bracket":     TokenOpenBracket,
		"close_bracket":    TokenCloseBracket,
		"open_brace":       TokenOpenBrace,
		"close_brace":      TokenCloseBrace,
		"functions_decl":   TokenFunctionsDecl,
		"constraints_decl": TokenConstraintsDecl,
		"nat_const":        TokenNatConst,
		"lc_ident_full":    TokenLcIdentFull,
		"lc_ident":         TokenLcIdent,
		"lc_ident_ns":      TokenLcIdentNs,
		"uc_ident":         TokenUcIdent,
		"uc_ident_ns":      TokenUcIdentNs,
		"equals":           TokenEquals,
		"hash":             TokenHash,
		"question_mark":    TokenQuestionMark,
		"percent":          TokenPercent,
		"plus":             TokenPlus,
		"langle":           TokenLeftAngle,
		"rangle":           TokenRightAngle,
		"comma":            TokenComma,
		"dot":              TokenDot,
		"asterisk":         TokenAsterisk,
		"excl_mark":        TokenExclMark,
		"final_kw":         TokenFinalKW,
		"new_kw":           TokenNewKW,
		"empty_kw":         TokenEmptyKW,
		"ws":               TokenWS,
		"newline":          TokenNewline,
		"comment":          TokenComment,
	}
}

func (d Definition) Lex(filename string, r io.Reader) (lexer.Lexer, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err //nolint:wrapcheck // lexer inherits io errors
	}

	return createLexer(filename, string(data)), nil
}

type defLexer struct {
	file string
	l    *L
	err  error
}

func createLexer(filename, data string) *defLexer {
	l := &defLexer{
		l: New(data, LexAny(lexers...)),
	}

	go func(d *defLexer) {
		// https://github.com/alecthomas/participle/discussions/223
		err := l.l.Start(context.TODO())
		if err != nil {
			panic("OH NO! ERROR!!!" + err.Error())
			// d.err = err
			// return
		}
	}(l)
	return l
}

func (d *defLexer) Next() (lexer.Token, error) {
	if d.err != nil {
		return lexer.Token{}, d.err
	}

	t, ok := d.l.NextToken()
	if !ok {
		return lexer.EOFToken(lexer.Position{}), nil
	}
	pos := lexer.Position{
		Filename: d.file,
		Offset:   t.Pos.Offset,
		Line:     t.Pos.Line,
		Column:   t.Pos.Column,
	}

	if t.Type == TokenEOF {
		return lexer.EOFToken(pos), d.err
	}

	return lexer.Token{
		Type:  t.Type,
		Value: t.Value,
		Pos:   pos,
	}, d.err
}
