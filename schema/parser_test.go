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
			TypeOrder: []ObjName{{Name: "CoolEnumerate"}},
			Objects:   map[ObjName]TypeObjects{},
			Enums: map[ObjName]EnumObjects{
				{Name: "CoolEnumerate"}: {
					Objects: []Object{{

						Name:   ObjName{Name: "someEnum"},
						CRC:    0x5508ec75,
						Fields: []Parameter{},
						Type:   TypeCommon(ObjName{Name: "CoolEnumerate"}),
					}},
				},
			},
			MethodGroupOrder: []string{"", "auth"},
			MethodsGroups: map[string][]Object{
				"": {{
					Name:   ObjName{Name: "someFunc"},
					CRC:    0x7da07ec9,
					Fields: []Parameter{},
					Type:   TypeCommon(ObjName{Name: "CoolEnumerate"}),
				}},
				"auth": {{
					Name:   ObjName{Group: "auth", Name: "someFunc"},
					CRC:    0x7da07ec9,
					Fields: []Parameter{},
					Type:   TypeCommon(ObjName{Name: "CoolEnumerate"}),
				}},
			},
		},
	}, {
		file: "internal/testdata/many_flags.tl",
		expected: &Schema{
			TypeOrder: []ObjName{{Name: "ChatFull"}},
			Objects: map[ObjName]TypeObjects{
				{Name: "ChatFull"}: {
					Objects: []Object{{
						Name: ObjName{Name: "a"},
						CRC:  0xf2355507,
						Fields: []Parameter{BitflagParameter{
							Name: "flags",
						}, TriggerParameter{
							Name:        "opt_prop",
							FlagTrigger: "flags",
							BitTrigger:  3,
						}, BitflagParameter{
							Name: "flags2",
						}, OptionalParameter{
							Name:        "opt2_prop",
							Type:        TypeCommon(ObjName{Name: "double"}),
							FlagTrigger: "flags2",
							BitTrigger:  9,
						}, RequiredParameter{
							Name: "id",
							Type: TypeCommon(ObjName{Name: "long"}),
						}},
						Type: TypeCommon(ObjName{Name: "ChatFull"}),
					}},
				},
			},
			Enums:            map[ObjName]EnumObjects{},
			MethodGroupOrder: []string{},
			MethodsGroups:    map[string][]Object{},
		},
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := testdata.ReadFile(tt.file)
			require.NoError(t, err)

			got, err := ParseString(tt.file, string(data))
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestEquality(t *testing.T) {
	for _, tt := range []struct {
		name string
		file string
	}{{
		name: "simplest",
		file: "internal/testdata/simplest.tl",
	}, {
		name: "many_flags",
		file: "internal/testdata/many_flags.tl",
	}, {
		name: "with_comments",
		file: "internal/testdata/with_comments.tl",
	}} {
		tt := tt // for parallel tests

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := testdata.ReadFile(tt.file)
			require.NoError(t, err)

			got, err := ParseString(tt.file, string(data))
			require.NoError(t, err)

			assert.Equal(t, string(data), got.String())
		})
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}
