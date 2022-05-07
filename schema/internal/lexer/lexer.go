package lexer

import (
	"context"
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

type StateFunc func(*L) (StateFunc, error)

type ErrLexer struct {
	Pos     Position
	Wrapped error
}

func (e ErrLexer) Error() string {
	return fmt.Sprintf("%v:%v: %v", e.Pos.Column, e.Pos.Line, e.Wrapped.Error())
}

const EOFRune = rune(-1)

const (
	TokenEOF                 = lexer.EOF
	TokenErr lexer.TokenType = (iota * -1) - 1
	TokenUnderscore
	TokenColon
	TokenSemicolon
	TokenOpenPar
	TokenClosePar
	TokenOpenBracket
	TokenCloseBracket
	TokenOpenBrace
	TokenCloseBrace
	TokenFunctionsDecl
	TokenConstraintsDecl
	TokenNatConst
	TokenLcIdentFull
	TokenLcIdent
	TokenLcIdentNs
	TokenUcIdent
	TokenUcIdentNs
	TokenEquals
	TokenHash
	TokenQuestionMark
	TokenPercent
	TokenPlus
	TokenLeftAngle
	TokenRightAngle
	TokenComma
	TokenDot
	TokenAsterisk
	TokenExclMark
	TokenFinalKW
	TokenNewKW
	TokenEmptyKW
	TokenWS
	TokenNewline
	TokenComment
)

type Token struct {
	Type  lexer.TokenType
	Value string
	Pos   Position
}

// Position describes an arbitrary source position
// including the file, line, and column location.
// A Position is valid if the line number is > 0.
type Position struct {
	Offset int // exact symbol index
	Line   int // line number, starting at 1
	Column int // column number, starting at 1 (byte count)
}

type L struct {
	source     []rune
	start, pos int
	startState StateFunc
	tokens     chan Token
}

// New creates a returns a lexer ready to parse the given source code.
func New(src string, start StateFunc) *L {
	return &L{
		source:     []rune(src),
		startState: start,
		tokens:     make(chan Token),
	}
}

// Start begins executing the Lexer in an asynchronous manner.
func (l *L) Start(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	defer close(l.tokens)

	state := l.startState
	var err error
	for state != nil {
		if err = assertCancel(ctx); err != nil {
			return err
		}

		state, err = state(l)
		if err != nil {
			return ErrLexer{Pos: l.startPos(), Wrapped: err}
		}
	}
	return nil
}

// Current returns the value being being analyzed at this moment.
func (l *L) Current() string {
	return string(l.source[l.start:l.pos])
}

// Emit will receive a token type and push a new token with the current analyzed
// value into the tokens channel.
func (l *L) Emit(t lexer.TokenType) {
	if l.start == l.pos && t != lexer.EOF {
		panic("got empty token")
	}

	tok := Token{
		Type:  t,
		Value: l.Current(),
		Pos:   l.startPos(),
	}
	l.tokens <- tok
	l.start = l.pos
}

// Ignore clears the rewind stack and then sets the current beginning position
// to the current position in the source which effectively ignores the section
// of the source being analyzed.
func (l *L) Ignore() {
	l.start = l.pos
}

// Peek performs a Next operation immediately followed by a Rewind returning the
// peeked rune.
func (l *L) Peek() rune {
	r := l.Next()
	l.Rewind()
	return r
}

// Rewind will take the last rune read (if any) and rewind back. Rewinds can
// occur more than once per call to Next but you can never rewind past the
// last point a token was emitted.
func (l *L) Rewind() {
	l.pos--
	if l.pos < l.start {
		l.pos = l.start
	}
}

// RewindTotally rewinds all read runes
func (l *L) RewindTotally() {
	l.pos = l.start
}

// Next pulls the next rune from the Lexer and returns it, moving the position
// forward in the source.
func (l *L) Next() rune {
	if l.pos >= len(l.source) {
		return EOFRune
	}

	r := l.source[l.pos]
	l.pos++
	return r
}

func (l *L) IsNextEOF() bool {
	return l.pos >= len(l.source)
}

// Take receives a string containing all acceptable strings and will contine
// over each consecutive character in the source until a token not in the given
// string is encountered. This should be used to quickly pull token parts.
// returns count of taken tokens.
func (l *L) Take(f func(r rune) bool) int {
	var i int
	for ; l.pos+i < len(l.source); i++ {
		r := l.source[l.pos+i]
		if !f(r) {
			break
		}
	}

	l.pos += i
	return i
}

func (l *L) TakeExact(s string) bool {
	if !(len(l.source)-l.pos >= len(s) && string(l.source[l.pos:l.pos+len(s)]) == s) {
		return false
	}

	l.pos += len(s)
	return true
}

// NextToken returns the next token from the lexer and a value to denote whether
// or not the token is finished.
func (l *L) NextToken() (token *Token, next bool) {
	if tok, ok := <-l.tokens; ok {
		return &tok, true
	}

	return nil, false
}

// todo: cache current line and column
func (l *L) startPos() Position {
	line, column := 1, 1
	for i := 0; i < l.start; i++ {
		if l.source[i] == '\n' {
			line++
			column = 1
		} else {
			column++
		}
	}

	return Position{Line: line, Column: column}
}

func assertCancel(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return context.Canceled
	default:
		return nil
	}
}
