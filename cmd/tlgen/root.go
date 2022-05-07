// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package main

import (
	"bufio"
	"io"
	"sync"

	"github.com/gabriel-vasile/mimetype"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/xelaj/tl/schema"
)

func Action(_ *cli.Context) error {
	return nil
}

func DetectAndParseSchema(filename, predictedMime string, r io.Reader) (*schema.Schema, error) {
	if predictedMime == "" {
		buf := bufio.NewReader(r)
		r = buf
		b, err := buf.Peek(mimeBufferMax)
		if err != nil {
			return nil, errors.Wrap(err, "predicting data format")
		}
		predictedMime = predictMime(b)
	}

	switch predictedMime {
	case mimeTypeLang:
		data, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}

		return schema.ParseFile(filename, string(data))

	case mimeProtobuf:
		return nil, errors.New("protofiles are not supported yet")
	default:
		return nil, errors.New(predictedMime + " is not supported")
	}
}

// for both of these file formats, there is no official mime type, so using
// most common mime variants.
const (
	mimeTypeLang = "application/x-typelang"
	mimeProtobuf = "application/x-protofile"

	mimeBufferMax = 512
)

var onceExtend sync.Once

func predictMime(b []byte) string {
	onceExtend.Do(func() {
		mimetype.Extend(schema.IsTypeLangDefinition, mimeTypeLang, ".tl")
	})

	mimetype.SetLimit(mimeBufferMax)
	return mimetype.Detect(b).String()
}
