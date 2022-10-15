// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

import (
	"context"
	"errors"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/xelaj/tl/schema/internal/lexer"
)

func IsTypeLangDefinition(raw []byte, limit uint32) (matched bool) {
	if !utf8.Valid(raw) {
		return false
	}

	str := string(raw)

	// expecting `someEnum#5508ec75 = CoolEnumerate;`
	var line string
	for len(str) > 0 {
		line, str = readLineWithoutComment(str)
		if line != "" {
			break
		}
	}

	l := lexer.New(line, lexer.LexAny())
	var wg sync.WaitGroup
	var err error
	ctx, cancel := context.WithCancel(context.TODO())
	defer func() {
		cancel()
		wg.Done()
		matched = matched && err != nil && !errors.Is(err, context.Canceled)
	}()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		err = l.Start(ctx)
	}(ctx)

	var t *lexer.Token
	var ok bool
	if t, ok = l.NextToken(); !ok || t == nil ||
		!(t.Type == lexer.TokenLcIdent || t.Type == lexer.TokenLcIdentFull) {
		return false
	}
	if t, ok = l.NextToken(); !ok || t == nil ||
		t.Type != lexer.TokenWS {
		return false
	}
	// TODO: parse full line to make sure that this is typelang spec.

	return true
}

// returns line without comment, and without prefixed or suffixed spaces.
func readLineWithoutComment(str string) (line, left string) {
	i := strings.IndexRune(str, '\n')
	if i >= 0 {
		line, left = str[:i], str[i:]
	} else {
		line, left = str, ""
	}

	commentStart := strings.Index(line, "//")
	return strings.TrimSpace(line[:commentStart]), left
}
