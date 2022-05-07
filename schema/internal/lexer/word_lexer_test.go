// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package lexer_test

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/xelaj/tl/schema/internal/lexer"
)

func TestAllLexers(t *testing.T) {
	for _, tcase := range []struct {
		source   string
		initFunc StateFunc
		want     []*Token
		wantErr  assert.ErrorAssertionFunc
	}{{
		source:   "dhjsojfj.wsorigjwr#12345678",
		initFunc: LexLcIdentFull,
		want: []*Token{{
			Type:  TokenLcIdentFull,
			Value: "dhjsojfj.wsorigjwr#12345678",
			Pos:   Position{Line: 1, Column: 1},
		}},
	}, {
		source:   "dhjsojfj..iogjs.wsorigjwr#12345678",
		initFunc: LexLcIdentFull,
		wantErr:  assert.Error,
	}, {
		source:   "wsorigjwr#1234567",
		initFunc: LexLcIdentFull,
		want: []*Token{{
			Type:  TokenLcIdentFull,
			Value: "wsorigjwr#1234567",
			Pos:   Position{Line: 1, Column: 1},
		}},
	}, {
		source:   "dhjsojfj.iogjs.wsorigjwr#12345678",
		initFunc: LexLcIdentFull,
		wantErr:  assert.Error,
	}, {
		source:   "dhjsojfj.iogjs.wsorigjwr.#12345678",
		initFunc: LexLcIdentFull,
		wantErr:  assert.Error,
	}} {
		tcase.wantErr = noErrAsDefault(tcase.wantErr)
		ctx, cancel := context.WithCancel(context.Background())

		t.Run("", func(t *testing.T) {
			t.Cleanup(cancel)
			var got []*Token

			l := New(tcase.source, tcase.initFunc)

			var wg sync.WaitGroup

			wg.Add(1)
			go func() {
				defer wg.Done()

				for {
					t, ok := l.NextToken()
					if !ok {
						return
					}

					got = append(got, t)
				}
			}()

			err := l.Start(ctx)
			wg.Wait()

			if ok := tcase.wantErr(t, err); ok || err != nil {
				return
			}

			assert.Equal(t, tcase.want, got)
		})
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}
