// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE.md for details.

package main

import (

	"io"
	"strings"
	"unicode/utf8"

	"github.com/gabriel-vasile/mimetype"
	"github.com/urfave/cli/v2"

	"github.com/xelaj/tl/schema"
	"github.com/xelaj/tl/schema/internal/lexer"
)

func Action(c *cli.Context) error {
	return nil
}

func DetectAndParseSchema(predictedMime string, r io.Reader) (*schema.Schema, error) {

}

// for both of these file formats, there is no official mime type, so using
// most common mime variants.
const (
	mimeTypeLang = "application/x-typelang"
	mimeProtobuf = "application/x-protofile"
)

func predictMime(r io.Reader) (string, error) {
	mimetype.Extend(typelangDetector, mimeTypeLang, ".tl")

}
