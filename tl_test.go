// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

//revive:disable:add-constant    // It's a test file, we can't avoid magic constants
//revive:disable:function-length // who cares for test files

package tl_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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

func TestDecode(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		data     []byte
		expected any
		wantErr  assert.ErrorAssertionFunc
	}{{
		name: "authSentCode",
		data: Hexed(`
			0225005e                                 // crc
			02000000                                 // flag
			        8659bb3d                         // crc AuthSentCodeTypeApp
					05000000                         // int32
			12                                       // length of text
			  316637366461306431353531313539363336   // text
			                                      00 // padding
			8c15a372                                 // next type`),
		expected: &AuthSentCode{
			Type: &AuthSentCodeTypeApp{
				Length: 5,
			},
			PhoneCodeHash: "1f76da0d1551159636",
			NextType:      AuthCodeTypeSms,
		},
	}, {
		name: "poll-results",
		data: Hexed("a3c1dcba1e00000015c4b51c02000000d2da6d3b010000000301020302000000d2da6d3b" +
			"0000000003040506060000000c00000015c4b51c02000000050000000600000005616c616c610000" +
			"15c4b51c00000000"),
		expected: &PollResults{
			Min: false,
			Results: []*PollAnswerVoters{
				{
					Chosen:  true,
					Correct: false,
					Option:  Hexed("010203"),
					Voters:  2,
				},
				{
					Chosen:  false,
					Correct: false,
					Option:  Hexed("040506"),
					Voters:  6,
				},
			},
			TotalVoters: ptr[int32](12),
			RecentVoters: []int32{
				5,
				6,
			},
			Solution:         ptr("alala"),
			SolutionEntities: []MessageEntity{},
		},
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)
		got := reflect.New(reflect.TypeOf(tt.expected)).Elem().Interface()

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := Unmarshal(tt.data, &got)

			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

// specific problem: if you try to create zero boolean value, you won't able to
// push it to decoder, cause reflect generates empty interface which contains
// bool. Decoder doesn't support any empty interfaces, cause it has strict rules
// about types.
func TestDecodeBool(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		data     []byte
		expected bool
		wantErr  assert.ErrorAssertionFunc
	}{{
		name:     "nakedBooleanObject#00",
		data:     Hexed("b5757299"),
		expected: true,
	}, {
		name:     "nakedBooleanObject#01",
		data:     Hexed("379779bc"),
		expected: false,
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got bool
			err := Unmarshal(tt.data, &got)

			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestDecodeUnknown(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name            string
		data            []byte
		hintsForDecoder []reflect.Type
		expected        any
		wantErr         assert.ErrorAssertionFunc
	}{{
		name: "authSentCode",
		data: Hexed(`
			0225005e                                 // crc
		    02000000                                 // flag
			        8659bb3d                         // crc AuthSentCodeTypeApp
					05000000                         // int32
			12                                       // length of text
			  316637366461306431353531313539363336   // text
			                                      00 // padding
			8c15a372                                 // next type`),
		expected: &AuthSentCode{
			Type: &AuthSentCodeTypeApp{
				Length: 5,
			},
			PhoneCodeHash: "1f76da0d1551159636",
			NextType:      AuthCodeTypeSms,
		},
	}, {
		name: "poll-results",
		data: Hexed("a3c1dcba1e00000015c4b51c02000000d2da6d3b010000000301020302000000d2da6d3b" +
			"0000000003040506060000000c00000015c4b51c02000000050000000600000005616c616c610000" +
			"15c4b51c00000000"),
		expected: &PollResults{
			Min: false,
			Results: []*PollAnswerVoters{
				{
					Chosen:  true,
					Correct: false,
					Option:  Hexed("010203"),
					Voters:  2,
				},
				{
					Chosen:  false,
					Correct: false,
					Option:  Hexed("040506"),
					Voters:  6,
				},
			},
			TotalVoters: ptr[int32](12),
			RecentVoters: []int32{
				5,
				6,
			},
			Solution:         ptr("alala"),
			SolutionEntities: []MessageEntity{},
		},
	}, {
		name:     "nakedBooleanObject#00",
		data:     Hexed("b5757299"),
		expected: true,
	}, {
		name:     "nakedBooleanObject#01",
		data:     Hexed("379779bc"),
		expected: false,
	}, {
		name: "issue_59", // https://github.com/xelaj/mtproto/issues/59
		data: Hexed(`
			6181e186                                         // crc
			100000006115f84a                                 // id
			00000000                                         // flag
			15                                               // length of text
			  d094d0bed181d182d0b0d182d0bed187d0bdd0be3f     // text
			                                            0000 // padding
            15c4b51c                                         // slice
                    03000000                                 // length of slice
				            e9c2a96c                         // crc
							32                               // length of text
							  d0b4d0bed181d182d0b0d182d0bed1 // text
							87d0bdd0be20d182d0bed0bbd18cd0ba
							d0be20d180d0b0d181d0bfd0b8d181d0
							bad0b8
						          00                         // padding
							01                               // length of bytes
							  30                             // bytes
							    0000                         // padding
							e9c2a96c                         // crc
							56                               // length of text
							  d0bfd0bed0bcd0b8d0bcd0be20d180 // text
							d0b0d181d0bfd0b8d181d0bad0b820d0
							bdd183d0b6d0bdd18b20d181d0b2d0b8
							d0b4d0b5d182d0b5d0bbd18cd181d0ba
							d0b8d0b520d0bfd0bed0bad0b0d0b7d0
							b0d0bdd0b8d18f
							              00                 // padding
                            01                               // length of bytes
							  31                             // bytes
							    0000                         // padding
							e9c2a96c                         // crc
							a7                               // length of text
				              d0bad180d0bed0bcd0b520d180d0b0 // text
							d181d0bfd0b8d181d0bad0b820d0bad1
							80d0b5d0b4d0b8d182d0bed180d18320
							d0bdd183d0b6d0bdd0be20d0b4d0bed0
							bad0b0d0b7d0b0d182d18c20d0bdd0b0
							d0bbd0b8d187d0b8d0b520d182d0b0d0
							bad0bed0b920d181d183d0bcd0bcd18b
							20d0bdd0b020d0bcd0bed0bcd0b5d0bd
							d18220d0b7d0b0d0bad0bbd18ed187d0
							b5d0bdd0b8d18f20d0b4d0bed0b3d0be
							d0b2d0bed180d0b0
							01                               // length of bytes
							  32                             // bytes
							    0000                         // padding`),
		expected: &Poll{
			ID:       5402091259386920976,
			Question: "Достаточно?",

			// don't mind on these texts, i'm too lazy to edit them
			Answers: []*PollAnswer{{
				Text:   "достаточно только расписки",
				Option: Hexed("30"),
			}, {
				Text:   "помимо расписки нужны свидетельские показания",
				Option: Hexed("31"),
			}, {
				Text:   "кроме расписки кредитору нужно доказать наличие такой суммы на момент заключения договора", //nolint:lll // mne pohuy, eto testy
				Option: Hexed("32"),
			}},
		},
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res, err := UnmarshalUnknown(tt.data)
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		obj     any
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{{
		name: "Rights",
		obj: &Rights{
			DeleteMessages: true,
			BanUsers:       true,
		},
		want:    Hexed("D524B25F18000000"),
		wantErr: assert.NoError,
	}, {
		name: "AccountInstallThemeParams",
		obj: &AccountInstallThemeParams{
			Dark:   true,
			Format: ptr("abc"),
			Theme: &InputThemeObj{
				ID:         123,
				AccessHash: 321,
			},
		},
		want:    Hexed("3737E47A0300000003616263E993563C7B000000000000004101000000000000"),
		wantErr: assert.NoError,
	}, {
		name: "AccountUnregisterDeviceParams",
		obj: &AccountUnregisterDeviceParams{
			TokenType: 1,
			Token:     "foo",
			OtherUids: []int32{
				1337, 228, 322,
			},
		},
		want:    Hexed("BFC476300100000003666F6F15C4B51C0300000039050000E400000042010000"),
		wantErr: assert.NoError,
	}, {
		name: "respq",
		obj: &ResPQ{
			Nonce:        NewInt128(123),
			ServerNonce:  NewInt128(321),
			Pq:           []byte{1, 2, 3},
			Fingerprints: []int64{322, 1337},
		},
		want: Hexed("632416050000000000000000000000000000007B00000000000000000000" +
			"0000000001410301020315C4B51C0200000042010000000000003905000000000000"),
		wantErr: assert.NoError,
	}, {
		name: "InitConnectionParams",
		obj: &InvokeWithLayerParams{
			Layer: 322,
			Query: &InitConnectionParams{
				APIID:          1337,
				DeviceModel:    "abc",
				SystemVersion:  "def",
				AppVersion:     "123",
				SystemLangCode: "en",
				LangCode:       "en",
				Query:          &SomeNullStruct{},
			},
		},
		want: Hexed(`
			0d0d9bda         // crc
			42010000         // Layer
			a95ecdc1         // crc InitConnectionParams
			        00000000 // flag
			        39050000 // api id
			        03       // length of string
			          616263 // data
			        03       // length of string
			          646566 // data
		            03       // length of string
			          313233 // data
			        02       // length of string
			          656e   // data
			              00 // padding
					00000000 // empty required string
			        02       // length of string
			          656e   // data
			              00 // padding
			        6b18f9c4 // null struct crc`),
		wantErr: assert.NoError,
	}} {
		tt := tt // for parallel tests

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Marshal(tt.obj)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

// checking that serializing and deserializing again got same result.
func TestEquality(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		obj  any
	}{{
		name: "MessagesChatsObj",
		obj: &MultipleChats{
			Chats: []any{
				&Chat{
					Creator:           false,
					Kicked:            false,
					Left:              false,
					Deactivated:       true,
					ID:                123,
					Title:             "abcdef",
					Photo:             "pikcha.png",
					ParticipantsCount: 123,
					Date:              1,
					Version:           1,
					AdminRights: &Rights{
						DeleteMessages: true,
						BanUsers:       true,
					},
					BannedRights: &Rights{
						DeleteMessages: false,
						BanUsers:       false,
					},
				},
			},
		},
	}} {
		tt := tt // for parallel tests
		got := reflect.New(reflect.TypeOf(tt.obj)).Elem().Interface()

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			encoded, err := Marshal(tt.obj)
			require.NoError(t, err)

			err = Unmarshal(encoded, &got)
			require.NoError(t, err)

			require.Equal(t, tt.obj, got)
		})
	}
}

// checking that serializing and deserializing again got same result.
func TestParseTag(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name      string
		tag       string
		fieldName string
		want      *StructTag
		wantErr   assert.ErrorAssertionFunc
	}{{
		tag:       "",
		fieldName: "SomeField",
		want: &StructTag{
			Name: "SomeField",
		},
	}, {
		tag:       "some_field",
		fieldName: "SomeField",
		want: &StructTag{
			Name: "some_field",
		},
	}, {
		tag:       ",omitempty:bitflag:30",
		fieldName: "SomeField",
		want: &StructTag{
			Name: "SomeField",
			BitFlags: &Bitflag{
				TargetField: "bitflag",
				BitPosition: 30,
			},
		},
	}, {
		tag:       ",omitempty:bitflag:30,implicit",
		fieldName: "SomeField",
		want: &StructTag{
			Name: "SomeField",
			BitFlags: &Bitflag{
				TargetField: "bitflag",
				BitPosition: 30,
			},
			Implicit: true,
		},
	}, {
		tag:       ",omitempty:otherflag",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",omitempty:otherflag:1000",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",omitempty:otherflag:-1",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",implicit",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",bitflag",
		fieldName: "SomeField",
		want: &StructTag{
			Name:      "SomeField",
			IsBitflag: true,
		},
	}, {
		tag:       "some_field,abracadabre",
		fieldName: "SomeField",
		wantErr:   assert.Error,
	}, {
		tag:       ",omitempty:bitflags:0,implicit,bitflag",
		fieldName: "SomeField",
		wantErr:   assert.Error,
	}, {
		tag:       ",implicit",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",omitempty:global_bitflags:0,bitflag",
		fieldName: "subflags",
		want: &StructTag{
			Name: "subflags",
			BitFlags: &Bitflag{
				TargetField: "global_bitflags",
				BitPosition: 0,
			},
			IsBitflag: true,
		},
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseTag(tt.tag, tt.fieldName)
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			require.Equal(t, tt.want, got)
		})
	}
}
