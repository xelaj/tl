// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"crypto/rand"
)

type null = struct{}

func randBytes(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}
