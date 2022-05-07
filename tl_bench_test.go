// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

//go:build !race
// +build !race

//revive:disable:add-constant It's a test file, we can't avoid magic constants

package tl_test

import (
	"testing"

	. "github.com/xelaj/tl"
)

func BenchmarkEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkFunc(Marshal(&AccountInstallThemeParams{
			Dark:   true,
			Format: ptr("abc"),
			Theme: &InputThemeObj{
				ID:         123,
				AccessHash: 321,
			},
		}))
	}
}
