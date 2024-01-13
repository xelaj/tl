// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

//revive:disable:add-constant    // It's a test file, we can't avoid magic constants
//revive:disable:function-length // who cares for test files

package tl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/xelaj/tl"
)

type TestCase interface {
	Name() string
	Run(t *testing.T)
}

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

type TcaseDecode[T any] struct {
	name     string
	data     []byte
	expected T
	wantErr  assert.ErrorAssertionFunc
}

var _ TestCase = (*TcaseDecode[any])(nil)

func (tt TcaseDecode[T]) Name() string { return tt.name }

func (tt TcaseDecode[T]) Run(t *testing.T) {
	t.Parallel()

	tt.wantErr = noErrAsDefault(tt.wantErr)

	var got T
	err := Unmarshal(tt.data, &got)

	if !tt.wantErr(t, err) || err != nil {
		t.Logf("%+v", err)
		return
	}

	assert.Equal(t, tt.expected, got)
}

func TestDecode(t *testing.T) {
	t.Parallel()

	for _, tt := range []TestCase{
		TcaseDecode[AuthSentCode]{
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
			expected: AuthSentCode{
				Type: &AuthSentCodeTypeApp{
					Length: 5,
				},
				PhoneCodeHash: "1f76da0d1551159636",
				NextType:      AuthCodeTypeSms,
			},
		}, TcaseDecode[PollResults]{
			name: "poll-results",
			data: Hexed("a3c1dcba1e00000015c4b51c02000000d2da6d3b010000000301020302000000d2da6d3b" +
				"0000000003040506060000000c00000015c4b51c02000000050000000600000005616c616c610000" +
				"15c4b51c00000000"),
			expected: PollResults{
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
		}, TcaseDecode[Poll]{
			name: "issue_59", // https://github.com/xelaj/mtproto/issues/59
			data: Hexed(`
			6181e186                                         // crc
			100000006115f84a                                 // id
			14000000                                         // flag
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
							    0000                         // padding
								    00000000                 // ClosePeriod`),
			expected: Poll{
				ID:             5402091259386920976,
				MultipleChoice: true,
				Question:       "Достаточно?",

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
				ClosePeriod: ptr(int32(0)),
			},
		}, TcaseDecode[DHParamsOk]{
			name: "long byte sequence",
			data: Hexed(`
            5c07e8d0                         // crc
            0000000000000000000000000000da09 // nonce
            00000000000000000000000000007c7a // server nonce
            fe                               // extension
              500200                         // real size (uint24 little endian)
            3ac422e37b570c2bddfddc0cac1f7d18 // data
            d3c6facad6d8966e08c59045957e3522
            541b54e80d848719ecfeb63e0ce90ae6
            a20b3bedbdb20bcb43666af9e1f53a2f
            7892bbaff2635e0ede2111003675214a
            01d5b2eb7249120c1d1534b521d49091
            b2cd33e51f1665406889df357ae2d8fe
            9b9bc544cdfa96fc4b07575ce0a42d84
            25ff354c1f89b1e2ac4210b6798791f9
            8f3016a8bdbbd6f840e54e339c340225
            23e8b7e0d2e16b6ae84fcbbfe1eb7e2a
            469c3af24e4d7950af2012de147e2eb1
            c118aaedef80c1534c605831bc35c01d
            fb606d89773abf92a0e984816e97b839
            d53ed72ccfd47a366c8434b732869b58
            81c67898a7dfaa6a141930ecf368f927
            ac07cae9887712defdd07203ba85cafb
            b72cb5f91aa81a9ebcebd3a7c5d3b67a
            72f334343795092ddaa2ab81a35807f1
            946f6d981adb24bf6740b39b4650b48f
            5546526f805629904ef648d1d19d6673
            eead1431f6c6937e52946a360356ae9e
            1ce64199ba0a2e679337eba31b49b12e
            ad1eddc98a47efe45f366006a5250735
            6f9b5b8f879116c90bcd1313f28fb186
            5ddfbc2400082b1f4ea54b84e41dfb48
            75ef606eeb81d4a53b5fada701116a5c
            db3877fa5b148e1afeb3e753997aab97
            bdc1c3f2fed829498bf523ca28c51ae1
            bce69c6a9fde708c7c6be906f5cdbd0e
            a977cd401227d94ac51713f64673f25c
            f3eb726dc46f333bc11a6d39b01a1610
            dbe28e428dd8627e36b81ba5477f3135
            0cb7583e238b9434de69dc030a9be749
            fccbc232b99db5e9f90c2d732c3930ae
            2ddec73c7a38ef29f9470343a7e316a0
            0366b11fdd95e4351ce2b59c9eca3d76
            `),
			expected: DHParamsOk{
				Nonce:       NewInt128(55817),
				ServerNonce: NewInt128(31866),
				EncryptedAnswer: []byte{
					0x3a, 0xc4, 0x22, 0xe3, 0x7b, 0x57, 0x0c, 0x2b, 0xdd, 0xfd, 0xdc, 0x0c, 0xac, 0x1f, 0x7d, 0x18,
					0xd3, 0xc6, 0xfa, 0xca, 0xd6, 0xd8, 0x96, 0x6e, 0x08, 0xc5, 0x90, 0x45, 0x95, 0x7e, 0x35, 0x22,
					0x54, 0x1b, 0x54, 0xe8, 0x0d, 0x84, 0x87, 0x19, 0xec, 0xfe, 0xb6, 0x3e, 0x0c, 0xe9, 0x0a, 0xe6,
					0xa2, 0x0b, 0x3b, 0xed, 0xbd, 0xb2, 0x0b, 0xcb, 0x43, 0x66, 0x6a, 0xf9, 0xe1, 0xf5, 0x3a, 0x2f,
					0x78, 0x92, 0xbb, 0xaf, 0xf2, 0x63, 0x5e, 0x0e, 0xde, 0x21, 0x11, 0x00, 0x36, 0x75, 0x21, 0x4a,
					0x01, 0xd5, 0xb2, 0xeb, 0x72, 0x49, 0x12, 0x0c, 0x1d, 0x15, 0x34, 0xb5, 0x21, 0xd4, 0x90, 0x91,
					0xb2, 0xcd, 0x33, 0xe5, 0x1f, 0x16, 0x65, 0x40, 0x68, 0x89, 0xdf, 0x35, 0x7a, 0xe2, 0xd8, 0xfe,
					0x9b, 0x9b, 0xc5, 0x44, 0xcd, 0xfa, 0x96, 0xfc, 0x4b, 0x07, 0x57, 0x5c, 0xe0, 0xa4, 0x2d, 0x84,
					0x25, 0xff, 0x35, 0x4c, 0x1f, 0x89, 0xb1, 0xe2, 0xac, 0x42, 0x10, 0xb6, 0x79, 0x87, 0x91, 0xf9,
					0x8f, 0x30, 0x16, 0xa8, 0xbd, 0xbb, 0xd6, 0xf8, 0x40, 0xe5, 0x4e, 0x33, 0x9c, 0x34, 0x02, 0x25,
					0x23, 0xe8, 0xb7, 0xe0, 0xd2, 0xe1, 0x6b, 0x6a, 0xe8, 0x4f, 0xcb, 0xbf, 0xe1, 0xeb, 0x7e, 0x2a,
					0x46, 0x9c, 0x3a, 0xf2, 0x4e, 0x4d, 0x79, 0x50, 0xaf, 0x20, 0x12, 0xde, 0x14, 0x7e, 0x2e, 0xb1,
					0xc1, 0x18, 0xaa, 0xed, 0xef, 0x80, 0xc1, 0x53, 0x4c, 0x60, 0x58, 0x31, 0xbc, 0x35, 0xc0, 0x1d,
					0xfb, 0x60, 0x6d, 0x89, 0x77, 0x3a, 0xbf, 0x92, 0xa0, 0xe9, 0x84, 0x81, 0x6e, 0x97, 0xb8, 0x39,
					0xd5, 0x3e, 0xd7, 0x2c, 0xcf, 0xd4, 0x7a, 0x36, 0x6c, 0x84, 0x34, 0xb7, 0x32, 0x86, 0x9b, 0x58,
					0x81, 0xc6, 0x78, 0x98, 0xa7, 0xdf, 0xaa, 0x6a, 0x14, 0x19, 0x30, 0xec, 0xf3, 0x68, 0xf9, 0x27,
					0xac, 0x07, 0xca, 0xe9, 0x88, 0x77, 0x12, 0xde, 0xfd, 0xd0, 0x72, 0x03, 0xba, 0x85, 0xca, 0xfb,
					0xb7, 0x2c, 0xb5, 0xf9, 0x1a, 0xa8, 0x1a, 0x9e, 0xbc, 0xeb, 0xd3, 0xa7, 0xc5, 0xd3, 0xb6, 0x7a,
					0x72, 0xf3, 0x34, 0x34, 0x37, 0x95, 0x09, 0x2d, 0xda, 0xa2, 0xab, 0x81, 0xa3, 0x58, 0x07, 0xf1,
					0x94, 0x6f, 0x6d, 0x98, 0x1a, 0xdb, 0x24, 0xbf, 0x67, 0x40, 0xb3, 0x9b, 0x46, 0x50, 0xb4, 0x8f,
					0x55, 0x46, 0x52, 0x6f, 0x80, 0x56, 0x29, 0x90, 0x4e, 0xf6, 0x48, 0xd1, 0xd1, 0x9d, 0x66, 0x73,
					0xee, 0xad, 0x14, 0x31, 0xf6, 0xc6, 0x93, 0x7e, 0x52, 0x94, 0x6a, 0x36, 0x03, 0x56, 0xae, 0x9e,
					0x1c, 0xe6, 0x41, 0x99, 0xba, 0x0a, 0x2e, 0x67, 0x93, 0x37, 0xeb, 0xa3, 0x1b, 0x49, 0xb1, 0x2e,
					0xad, 0x1e, 0xdd, 0xc9, 0x8a, 0x47, 0xef, 0xe4, 0x5f, 0x36, 0x60, 0x06, 0xa5, 0x25, 0x07, 0x35,
					0x6f, 0x9b, 0x5b, 0x8f, 0x87, 0x91, 0x16, 0xc9, 0x0b, 0xcd, 0x13, 0x13, 0xf2, 0x8f, 0xb1, 0x86,
					0x5d, 0xdf, 0xbc, 0x24, 0x00, 0x08, 0x2b, 0x1f, 0x4e, 0xa5, 0x4b, 0x84, 0xe4, 0x1d, 0xfb, 0x48,
					0x75, 0xef, 0x60, 0x6e, 0xeb, 0x81, 0xd4, 0xa5, 0x3b, 0x5f, 0xad, 0xa7, 0x01, 0x11, 0x6a, 0x5c,
					0xdb, 0x38, 0x77, 0xfa, 0x5b, 0x14, 0x8e, 0x1a, 0xfe, 0xb3, 0xe7, 0x53, 0x99, 0x7a, 0xab, 0x97,
					0xbd, 0xc1, 0xc3, 0xf2, 0xfe, 0xd8, 0x29, 0x49, 0x8b, 0xf5, 0x23, 0xca, 0x28, 0xc5, 0x1a, 0xe1,
					0xbc, 0xe6, 0x9c, 0x6a, 0x9f, 0xde, 0x70, 0x8c, 0x7c, 0x6b, 0xe9, 0x06, 0xf5, 0xcd, 0xbd, 0x0e,
					0xa9, 0x77, 0xcd, 0x40, 0x12, 0x27, 0xd9, 0x4a, 0xc5, 0x17, 0x13, 0xf6, 0x46, 0x73, 0xf2, 0x5c,
					0xf3, 0xeb, 0x72, 0x6d, 0xc4, 0x6f, 0x33, 0x3b, 0xc1, 0x1a, 0x6d, 0x39, 0xb0, 0x1a, 0x16, 0x10,
					0xdb, 0xe2, 0x8e, 0x42, 0x8d, 0xd8, 0x62, 0x7e, 0x36, 0xb8, 0x1b, 0xa5, 0x47, 0x7f, 0x31, 0x35,
					0x0c, 0xb7, 0x58, 0x3e, 0x23, 0x8b, 0x94, 0x34, 0xde, 0x69, 0xdc, 0x03, 0x0a, 0x9b, 0xe7, 0x49,
					0xfc, 0xcb, 0xc2, 0x32, 0xb9, 0x9d, 0xb5, 0xe9, 0xf9, 0x0c, 0x2d, 0x73, 0x2c, 0x39, 0x30, 0xae,
					0x2d, 0xde, 0xc7, 0x3c, 0x7a, 0x38, 0xef, 0x29, 0xf9, 0x47, 0x03, 0x43, 0xa7, 0xe3, 0x16, 0xa0,
					0x03, 0x66, 0xb1, 0x1f, 0xdd, 0x95, 0xe4, 0x35, 0x1c, 0xe2, 0xb5, 0x9c, 0x9e, 0xca, 0x3d, 0x76,
				},
			},
		},
	} {
		t.Run(tt.Name(), tt.Run)
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
				t.Logf("%+v", err)
				return
			}

			assert.Equal(t, tt.expected, got)
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
		name: "respq_raw",
		obj: &ResPQRaw{
			Nonce:        [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7b},
			ServerNonce:  [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x41},
			Pq:           []byte{1, 2, 3},
			Fingerprints: []int64{322, 1337},
		},
		want: Hexed("642416050000000000000000000000000000007B00000000000000000000" +
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

type TcaseEquality[T any] struct {
	name string
	obj  T
}

var _ TestCase = (*TcaseEquality[any])(nil)

func (tt TcaseEquality[T]) Name() string { return tt.name }

func (tt TcaseEquality[T]) Run(t *testing.T) {
	t.Parallel()

	encoded, err := Marshal(&tt.obj)
	require.NoError(t, err)

	var got T
	err = Unmarshal(encoded, &got)
	require.NoError(t, err, "type %T", tt.obj)

	require.Equal(t, tt.obj, got)
}

// checking that serializing and deserializing again got same result.
func TestEquality(t *testing.T) {
	for _, tt := range []TestCase{
		TcaseEquality[MultipleChats]{
			name: "MultipleChats",
			obj: MultipleChats{
				Chats: []Chat{{
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
				}},
			},
		}, TcaseEquality[InvokeWithLayerParams]{
			name: "InitConnectionParams",
			obj: InvokeWithLayerParams{
				Layer: 322,
				Query: &InitConnectionParams{
					APIID:          1337,
					DeviceModel:    "abc",
					SystemVersion:  "def",
					AppVersion:     "123",
					SystemLangCode: "en",
					LangCode:       "en",
					Proxy:          &SomeNullStruct{},
					Query: &Rights{
						DeleteMessages: true,
						BanUsers:       true,
					},
				},
			},
		}, TcaseEquality[GenericRequest[[]int64]]{
			name: "complex response of Vector<long>",
			obj: GenericRequest[[]int64]{
				MsgID:          1234,
				IsSpecificType: true,
				Msg:            []int64{1, 3, 5, 7},
			},
		}, TcaseEquality[ResPQRaw]{
			name: "respq_raw",
			obj: ResPQRaw{
				Nonce:        [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7b},
				ServerNonce:  [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x41},
				Pq:           []byte{1, 2, 3},
				Fingerprints: []int64{322, 1337},
			},
		},
	} {
		t.Run(tt.Name(), tt.Run)
	}
}

// checking that serializing and deserializing again got same result.
func TestParseTag(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name      string
		tag       string
		fieldName string
		want      StructTag
		wantErr   assert.ErrorAssertionFunc
	}{{
		tag:       "",
		fieldName: "SomeField",
		want: StructTag{
			Name: "SomeField",
		},
	}, {
		tag:       "some_field",
		fieldName: "SomeField",
		want: StructTag{
			Name: "some_field",
		},
	}, {
		tag:       ",omitempty:bitflag:30",
		fieldName: "SomeField",
		want: StructTag{
			Name: "SomeField",
			BitFlags: &Bitflag{
				TargetField: "bitflag",
				BitPosition: 30,
			},
		},
	}, {
		tag:       ",omitempty:bitflag:30,implicit",
		fieldName: "SomeField",
		want: StructTag{
			Name: "SomeField",
			BitFlags: &Bitflag{
				TargetField: "bitflag",
				BitPosition: 30,
			},
			Type: ImplicitBoolType,
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
		want: StructTag{
			Name: "SomeField",
			Type: BitflagType,
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
		want: StructTag{
			Name: "subflags",
			BitFlags: &Bitflag{
				TargetField: "global_bitflags",
				BitPosition: 0,
			},
			Type: BitflagType,
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

type TcaseParseStructTags[T Object] struct {
	name         string
	wantTags     []StructTag
	wantBitflags map[int]BitflagBit
	wantErr      assert.ErrorAssertionFunc
}

var _ TestCase = (*TcaseParseStructTags[Object])(nil)

func (tt TcaseParseStructTags[T]) Name() string { return tt.name }

func (tt TcaseParseStructTags[T]) Run(t *testing.T) {
	tt.wantErr = noErrAsDefault(tt.wantErr)
	t.Parallel()

	var obj T
	gotTags, gotFlags, err := ParseStructTags(obj)
	if !tt.wantErr(t, err) || err != nil {
		return
	}

	require.Equal(t, tt.wantTags, gotTags)
	require.Equal(t, tt.wantBitflags, gotFlags)
}

func TestParseStructTags(t *testing.T) {
	for _, tt := range []TestCase{
		TcaseParseStructTags[*Poll]{
			name: "Poll",
			wantTags: []StructTag{
				{Name: "ID"},
				{Name: "flag", Type: BitflagType},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 0}, Name: "Closed", Type: ImplicitBoolType},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 1}, Name: "PublicVoters", Type: ImplicitBoolType},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 2}, Name: "MultipleChoice", Type: ImplicitBoolType},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 3}, Name: "Quiz", Type: ImplicitBoolType},
				{Name: "Question"},
				{Name: "Answers"},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 4}, Name: "ClosePeriod"},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 5}, Name: "CloseDate"},
			},
			wantBitflags: map[int]BitflagBit{
				2: {FieldIndex: 1, BitIndex: 0},
				3: {FieldIndex: 1, BitIndex: 1},
				4: {FieldIndex: 1, BitIndex: 2},
				5: {FieldIndex: 1, BitIndex: 3},
				8: {FieldIndex: 1, BitIndex: 4},
				9: {FieldIndex: 1, BitIndex: 5},
			},
		},
	} {
		t.Run(tt.Name(), tt.Run)
	}
}

// type Poll struct {
// 	ID int64
// 	//nolint:revive // tl works with unexported tags
// 	_              null `tl:"flag,bitflag"`
// 	Closed         bool `tl:",omitempty:flag:0,implicit"`
// 	PublicVoters   bool `tl:",omitempty:flag:1,implicit"`
// 	MultipleChoice bool `tl:",omitempty:flag:2,implicit"`
// 	Quiz           bool `tl:",omitempty:flag:3,implicit"`
// 	Question       string
// 	Answers        []*PollAnswer
// 	ClosePeriod    *int32 `tl:",omitempty:flag:4"`
// 	CloseDate      *int32 `tl:",omitempty:flag:5"`
// }
