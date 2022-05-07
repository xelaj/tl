// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

//nolint:nilnil // that's the main mechanic of lexers
package lexer

import (
	"unicode"

	"github.com/pkg/errors"
	"golang.org/x/text/unicode/rangetable"
)

const newline = '\n'

//nolint:gochecknoglobals // otherwise we can't create unicode tables
var (
	lcLetter   = singleRange16('a', 'z', true)
	ucLetter   = singleRange16('A', 'Z', true)
	digit      = singleRange16('0', '9', true) // unicode.Digit covers too many digits
	hexDigit   = unicode.ASCII_Hex_Digit       // extended hex digits with upper symbols
	underscore = rangetable.New('_')
	letter     = rangetable.Merge(lcLetter, ucLetter)
	identChar  = rangetable.Merge(letter, digit, underscore)
	divider    = rangetable.New(' ', '\t')
)

var lexers = []StateFunc{
	LexLcIdentFull,
	LexLcIdentNs,
	LexUcIdentNs,
	LexLcIdent,
	LexUcIdent,
	LexNatConst,
	LexWS,
	LexComment,
	LexNewline,
	LexFunctionsDecl,
	LexConstraintDecl,
	LexExact,
	LexEOF,
}

type ErrWrongToken struct{}

func (ErrWrongToken) Error() string {
	return "wrong token"
}

func LexAny(funcs ...StateFunc) StateFunc {
	return func(l *L) (StateFunc, error) {
		for _, f := range funcs {
			_, err := LexEOF(l)
			if err == nil {
				return nil, nil
			}

			s, err := f(l)
			if s == nil {
				s = LexAny(funcs...)
			}

			if err == nil {
				return s, nil
			}

			if errors.Is(err, ErrWrongToken{}) {
				l.RewindTotally() // just in case, if lexer won't rewind input

				continue
			}

			return nil, err
		}

		return nil, ErrWrongToken{}
	}
}

func LexWS(l *L) (StateFunc, error) {
	if !unicode.Is(divider, l.Peek()) {
		return nil, ErrWrongToken{}
	}
	l.Take(func(r rune) bool { return unicode.Is(divider, r) })
	l.Emit(TokenWS)
	return nil, nil
}

func LexComment(l *L) (StateFunc, error) {
	if !l.TakeExact("//") {
		return nil, ErrWrongToken{}
	}

	l.Take(func(r rune) bool { return r != newline })

	l.Emit(TokenComment)
	return nil, nil
}

func LexFunctionsDecl(l *L) (StateFunc, error) {
	if !l.TakeExact("---") {
		return nil, ErrWrongToken{}
	}

	l.Take(func(r rune) bool { return r == ' ' })
	if !l.TakeExact("functions") {
		return nil, ErrWrongToken{}
	}
	l.Take(func(r rune) bool { return r == ' ' })

	if !l.TakeExact("---") {
		return nil, ErrWrongToken{}
	}

	l.Emit(TokenFunctionsDecl)
	return nil, nil
}

func LexConstraintDecl(l *L) (StateFunc, error) {
	if !l.TakeExact("---") {
		return nil, ErrWrongToken{}
	}

	l.Take(func(r rune) bool { return r == ' ' })
	if !l.TakeExact("types") {
		return nil, ErrWrongToken{}
	}
	l.Take(func(r rune) bool { return r == ' ' })

	if !l.TakeExact("---") {
		return nil, ErrWrongToken{}
	}

	l.Emit(TokenConstraintsDecl)
	return nil, nil
}

func LexExact(l *L) (StateFunc, error) {
	switch {
	case l.TakeExact("_"):
		l.Emit(TokenUnderscore)

	case l.TakeExact(":"):
		l.Emit(TokenColon)

	case l.TakeExact(";"):
		l.Emit(TokenSemicolon)

	case l.TakeExact("("):
		l.Emit(TokenOpenPar)

	case l.TakeExact(")"):
		l.Emit(TokenClosePar)

	case l.TakeExact("["):
		l.Emit(TokenOpenBracket)

	case l.TakeExact("]"):
		l.Emit(TokenCloseBracket)

	case l.TakeExact("{"):
		l.Emit(TokenOpenBrace)

	case l.TakeExact("}"):
		l.Emit(TokenCloseBrace)

	case l.TakeExact("="):
		l.Emit(TokenEquals)

	case l.TakeExact("#"):
		l.Emit(TokenHash)

	case l.TakeExact("?"):
		l.Emit(TokenQuestionMark)

	case l.TakeExact("%"):
		l.Emit(TokenPercent)

	case l.TakeExact("+"):
		l.Emit(TokenPlus)

	case l.TakeExact("<"):
		l.Emit(TokenLeftAngle)

	case l.TakeExact(">"):
		l.Emit(TokenRightAngle)

	case l.TakeExact(","):
		l.Emit(TokenComma)

	case l.TakeExact("."):
		l.Emit(TokenDot)

	case l.TakeExact("*"):
		l.Emit(TokenAsterisk)

	case l.TakeExact("!"):
		l.Emit(TokenExclMark)

	case l.TakeExact("Final"), l.TakeExact("final"):
		l.Emit(TokenFinalKW)

	case l.TakeExact("New"), l.TakeExact("new"):
		l.Emit(TokenEmptyKW)

	case l.TakeExact("Empty"), l.TakeExact("empty"):
		l.Emit(TokenEmptyKW)

	default:
		return nil, ErrWrongToken{}
	}
	return nil, nil
}

func LexNatConst(l *L) (StateFunc, error) {
	if !unicode.Is(digit, l.Peek()) {
		return nil, ErrWrongToken{}
	}
	l.Take(func(r rune) bool { return unicode.Is(digit, r) })
	l.Emit(TokenNatConst)
	return nil, nil
}

func LexLcIdent(l *L) (StateFunc, error) {
	ok := takeLcIdent(l)
	if !ok {
		return nil, ErrWrongToken{}
	}

	l.Emit(TokenLcIdent)
	return nil, nil
}

func LexUcIdent(l *L) (StateFunc, error) {
	if !takeUcIdent(l) {
		return nil, ErrWrongToken{}
	}

	l.Emit(TokenUcIdent)
	return nil, nil
}

// [ lc-ident '.' ] lc-ident '#' hex-digit *8.
func LexLcIdentFull(l *L) (StateFunc, error) {
	if takeLcIdent(l) && l.Peek() == '.' {
		l.Next()
	} else {
		l.RewindTotally()
	}

	if !takeLcIdent(l) {
		return nil, ErrWrongToken{}
	}

	if l.Next() != '#' {
		return nil, ErrWrongToken{}
	}

	for i := 0; i < 8; i++ {
		h := l.Peek()
		// спецификация обязывает указывать 8 цифр, но в tl схеме телеги их ВНЕЗАПНО не всегда 8
		if !unicode.Is(hexDigit, h) {
			break
		}
		l.Next()
	}

	l.Emit(TokenLcIdentFull)
	return nil, nil
}

func LexLcIdentNs(l *L) (StateFunc, error) {
	ok := takeLcIdent(l)
	if !ok || l.Peek() != '.' {
		return nil, ErrWrongToken{}
	}
	l.Next()

	ok = takeLcIdent(l)
	if !ok {
		return nil, ErrWrongToken{}
	}
	l.Emit(TokenLcIdentNs)

	return nil, nil
}

// lc-ident '.' lc-ident.
func LexUcIdentNs(l *L) (StateFunc, error) {
	if !takeLcIdent(l) || l.Peek() != '.' {
		return nil, ErrWrongToken{}
	}
	l.Next()

	if !takeUcIdent(l) {
		return nil, ErrWrongToken{}
	}
	l.Emit(TokenUcIdentNs)

	return nil, nil
}

func takeLcIdent(l *L) bool {
	if !unicode.Is(lcLetter, l.Peek()) {
		return false
	}
	l.Next() // already picked

	l.Take(func(r rune) bool { return unicode.Is(identChar, r) })
	return true
}

func takeUcIdent(l *L) bool {
	if !unicode.In(l.Peek(), ucLetter) {
		return false
	}

	l.Next() // already picked
	l.Take(func(r rune) bool { return unicode.Is(identChar, r) })
	return true
}

func LexEOF(l *L) (StateFunc, error) {
	if l.IsNextEOF() {
		l.Emit(TokenEOF)
		return nil, nil
	}
	return nil, ErrWrongToken{}
}

func LexNewline(l *L) (StateFunc, error) {
	if l.Next() != newline {
		return nil, ErrWrongToken{}
	}
	l.Emit(TokenNewline)
	return nil, nil
}

//revive:disable:flag-parameter
func singleRange16(from, to rune, isLatin bool) *unicode.RangeTable {
	var offset int
	if isLatin {
		offset = 1
	}

	return &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: uint16(from), Hi: uint16(to), Stride: 1},
		},
		LatinOffset: offset,
	}
}

//revive:enable
