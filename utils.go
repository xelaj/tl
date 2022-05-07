// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

type null = struct{}

func littleUint24Bytes(v int) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),  //nolint:gomnd // r u kiddin
		byte(v >> 16), //nolint:gomnd // r u kiddin
	}
}

// TODO: пофиксить этот ужас, можно сделать проще, без знания размера длины

// https://go.dev/play/p/uU0LiPVCPnI
func pad(lengthOfSizeLength, wordLen, msgLen int) int {
	wMinMod := wordLen - ((lengthOfSizeLength + msgLen) % wordLen)

	return (1 - wMinMod/wordLen) * wMinMod
}

// v is bool, pointer, interface or slice.
func isFieldContainsData(v reflect.Value) bool {
	// special cases for enums
	if v, ok := v.Interface().(Enum); ok {
		return v.CRC() != 0
	}

	switch k := v.Kind(); k { // nolint:exhaustive // have default
	case reflect.Bool:
		return v.Bool()

	case reflect.Pointer, reflect.Interface, reflect.Slice:
		return !v.IsNil()

	default:
		panic(
			fmt.Sprintf(
				"something is wrong with registry, optional field must be pointer or bool, got %v",
				k,
			),
		)
	}
}

func appendMany[T any](b ...[]T) []T {
	var size int
	for _, s := range b {
		size += len(s)
	}
	tmp := make([]T, size)
	var i int
	for _, s := range b {
		i += copy(tmp[i:], s)
	}

	return tmp
}

func as118[T error](err error) (T, bool) {
	var converted T
	return converted, errors.As(err, &converted)
}

type convertibleNum interface {
	constraints.Integer | constraints.Float
}

func convertNumErr[V, T convertibleNum](res T, err error) (V, error) {
	return V(res), err
}

type convertibleStr interface {
	~string | ~[]byte
}

func convertStrErr[V, T convertibleStr](res T, err error) (V, error) {
	return V(res), err
}

func ptr[T any](value T) *T { return &value }

func u32b(order binary.ByteOrder, v uint32) []byte {
	b := make([]byte, WordLen)
	order.PutUint32(b, v)

	return b
}

func u64b(order binary.ByteOrder, v uint64) []byte {
	b := make([]byte, LongLen)
	order.PutUint64(b, v)

	return b
}

func f64b(order binary.ByteOrder, v float64) []byte {
	b := make([]byte, LongLen)
	order.PutUint64(b, math.Float64bits(v))

	return b
}
