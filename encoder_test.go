// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/xelaj/tl"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name    string
		obj     any
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Rights",
			obj: &Rights{
				DeleteMessages: true,
				BanUsers:       true,
			},
			want:    Hexed("D524B25F18000000"),
			wantErr: assert.NoError,
		},
		{
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
		},
		{
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
		},
		{
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
		},
		{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.obj)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
