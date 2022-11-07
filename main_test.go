// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

//revive:disable:function-length // who cares for test files

package tl_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/xelaj/tl"
)

func tearup() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("TEARUP ERROR: %v", err)
			os.Exit(2)
		}
	}()
	RegisterObjects(
		&MultipleChats{},
		&Chat{},
		&AuthSentCode{},
		&SomeNullStruct{},
		&AuthSentCodeTypeApp{},
		&Rights{},
		&PollResults{},
		&PollAnswerVoters{},
		&AccountInstallThemeParams{},
		&InputThemeObj{},
		&AccountUnregisterDeviceParams{},
		&InvokeWithLayerParams{},
		&InitConnectionParams{},
		&ResPQ{},
		&AnyStructWithAnyType{},
		&AnyStructWithAnyObject{},
		&Poll{},
		&PollAnswer{},
	)

	RegisterEnums(
		AuthCodeTypeSms,
		AuthCodeTypeCall,
		AuthCodeTypeFlashCall,
	)
}

func teardown() {}

func TestMain(m *testing.M) {
	tearup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func Hexed(in string) []byte {
	reader := bytes.NewReader([]byte(in))
	buf := []rune{}
	for {
		r, ok := readByte(reader)
		if !ok {
			break
		}
		if r != 0 {
			buf = append(buf, r)
		}
	}

	return checkFunc(hex.DecodeString(string(buf)))
}

func readByte(reader *bytes.Reader) (rune, bool) {
	r, ok := readAndCheck(reader)
	if !ok {
		return 0, false
	}
	switch r {
	case ' ', '\n', '\t':
		return 0, true

	case '/':
		if r, ok := readAndCheck(reader); !ok || r != '/' {
			panic("expected comment")
		}
		skipComment(reader)

		return 0, true

	default:
		return r, true
	}
}

func skipComment(reader *bytes.Reader) {
	for {
		r, ok := readAndCheck(reader)
		if !ok || r == '\n' {
			break
		}
	}
}

func readAndCheck(reader *bytes.Reader) (r rune, ok bool) {
	r, _, err := reader.ReadRune()
	if err == io.EOF {
		return 0, false
	}
	check(err)

	return r, true
}

func checkFunc[T any](res T, err error) T {
	check(err)

	return res
}

func check(err error) {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}

func ptr[T any](value T) *T { return &value }

type null = struct{}
