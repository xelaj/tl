// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema_test

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/xelaj/tl/schema"
)

//go:embed internal/testdata
var testdata embed.FS

func TestParseFile(t *testing.T) {
	for _, tt := range []struct {
		name     string
		file     string
		expected *Schema
		wantErr  assert.ErrorAssertionFunc
	}{{
		file: "internal/testdata/simplest.tl",
		expected: &Schema{
			Objects: []*Object{{
				Name:   "someEnum",
				CRC:    0x5508ec75,
				Fields: []Parameter{},
				Type:   TypeCommon("CoolEnumerate"),
			}, Method(&Object{
				Name:   "someFunc",
				CRC:    0x7da07ec9,
				Fields: []Parameter{},
				Type:   TypeCommon("CoolEnumerate"),
			})},
			TypeComments: map[string]string{},
		},
	}, {
		file: "internal/testdata/many_flags.tl",
		expected: &Schema{
			Objects: []*Object{{
				Name: "a",
				CRC:  0xf2355507,
				Fields: []Parameter{RequiredParameter{
					Name: "flags",
					Type: TypeCommon("#"),
				}, TriggerParameter{
					Name:        "opt_prop",
					FlagTrigger: "flags",
					BitTrigger:  3,
				}, RequiredParameter{
					Name: "flags2",
					Type: TypeCommon("#"),
				}, OptionalParameter{
					Name:        "opt2_prop",
					Type:        TypeCommon("double"),
					FlagTrigger: "flags2",
					BitTrigger:  9,
				}, RequiredParameter{
					Name: "id",
					Type: TypeCommon("long"),
				}},
				Type: TypeCommon("ChatFull"),
			}},
			TypeComments: map[string]string{},
		},
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := testdata.ReadFile(tt.file)
			require.NoError(t, err)

			got, err := ParseFile(tt.file, string(data))
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}
