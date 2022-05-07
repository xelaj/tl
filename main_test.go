// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	. "github.com/xelaj/tl"
)

func tearup() {
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
		r, ok := readAndCheck(reader)
		if !ok {
			break
		}
		switch r {
		case ' ', '\n', '\t':
			continue
		}

		if r == '/' {
			r, ok := readAndCheck(reader)
			if !ok {
				panic("expected comment")
			}
			if r != '/' {
				panic("expected comment")
			}
			skipComment(reader)
			continue
		}

		buf = append(buf, r)
	}

	res, err := hex.DecodeString(string(buf))

	check(err)
	return res
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

func check(err error) {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}
