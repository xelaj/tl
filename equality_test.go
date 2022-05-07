// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/xelaj/tl"
)

// checking that serializing and deserializing again got same result
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
		got := reflect.New(reflect.TypeOf(tt.obj)).Elem().Interface()

		t.Run(tt.name, func(t *testing.T) {
			encoded, err := Marshal(tt.obj)
			require.NoError(t, err)

			err = Unmarshal(encoded, &got)
			require.NoError(t, err)

			require.Equal(t, tt.obj, got)
		})
	}
}
