// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE.md for details.

package main

import (
	"strings"
	"unicode/utf8"

	"github.com/xelaj/tl/schema/internal/lexer"
)

func IsTypeLangDefinition(raw []byte, limit uint32) bool {
	if !utf8.Valid(raw) {
		return false
	}

	str := string(raw)

	// someEnum#5508ec75 = CoolEnumerate;
	var line string
	for len(str) > 0 {
		line, str = readLineWithoutComment(str)
		if line != "" {
			break
		}
	}

	lexer.New(line).Start()

	return false
}

// returns line without comment, and without prefixed or suffixed spaces
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
