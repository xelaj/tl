// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

//go:build !race
package tl_test

import (
	"testing"

	. "github.com/xelaj/tl"
)

func BenchmarkEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Marshal(&AccountInstallThemeParams{
			Dark:   true,
			Format: ptr("abc"),
			Theme: &InputThemeObj{
				ID:         123,
				AccessHash: 321,
			},
		})
	}
}
