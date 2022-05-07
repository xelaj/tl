// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/xelaj/tl"
)

var (
	True = true // for pointer
)

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
					Option: []byte{
						0x01, 0x02, 0x03,
					},
					Voters: 2,
				},
				{
					Chosen:  false,
					Correct: false,
					Option: []byte{
						0x04, 0x05, 0x06,
					},
					Voters: 6,
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
		tt.wantErr = noErrAsDefault(tt.wantErr)
		got := reflect.New(reflect.TypeOf(tt.expected)).Elem().Interface()

		t.Run(tt.name, func(t *testing.T) {
			err := Unmarshal(tt.data, &got)

			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

// specific problem: if you try to create zero boolean value, you won't able to push it to decoder, cause
// reflect generates empty interface which contains bool. Decoder doesn't support any empty interfaces, cause
// it has strict rules about types.
func TestDecodeBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		data     []byte
		expected bool
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "nakedBooleanObject#00",
			data:     Hexed("b5757299"),
			expected: true,
		},
		{
			name:     "nakedBooleanObject#01",
			data:     Hexed("379779bc"),
			expected: false,
		},
	}
	for _, tt := range tests {
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
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
					Option: []byte{
						0x01, 0x02, 0x03,
					},
					Voters: 2,
				},
				{
					Chosen:  false,
					Correct: false,
					Option: []byte{
						0x04, 0x05, 0x06,
					},
					Voters: 6,
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
	},
		{
			name:     "nakedBooleanObject#00",
			data:     Hexed("b5757299"),
			expected: true,
		},
		{
			name:     "nakedBooleanObject#01",
			data:     Hexed("379779bc"),
			expected: false,
		},
		{
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
				Answers: []*PollAnswer{
					{ // don't mind on these texts, i'm too lazy to edit them
						Text: "достаточно только расписки",
						Option: []uint8{
							0x30,
						},
					},
					{ // don't mind on these texts, i'm too lazy to edit them
						Text: "помимо расписки нужны свидетельские показания",
						Option: []uint8{
							0x31,
						},
					},
					{ // don't mind on these texts, i'm too lazy to edit them
						Text: "кроме расписки кредитору нужно доказать наличие такой суммы на момент заключения договора",
						Option: []uint8{
							0x32,
						},
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr = noErrAsDefault(tt.wantErr)

			res, err := UnmarshalUnknown(tt.data)
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.expected, res)
		})
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}
