// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

type null = struct{}

func unreachable()   { panic("Unreachable") }   //cover:ignore
func unimplemented() { panic("Unimplemented") } //cover:ignore

func randBytes(size int) []byte { //cover:ignore
	b := make([]byte, size)
	rand.Read(b)

	return b
}

func littleUint24Bytes(v int) []byte { //cover:ignore
	return []byte{
		byte(v),
		byte(v >> 8),  //nolint:gomnd // r u kiddin
		byte(v >> 16), //nolint:gomnd // r u kiddin
	}
}

// pad returns size of padding for message.
func pad(l, padSize int) int {
	mod := l % padSize

	return (padSize - mod) * (1 + (mod-padSize)/padSize)
}

// v is bool, pointer, interface or slice.
func isFieldContainsData(v reflect.Value) bool {
	// special cases for enums
	if enum, ok := v.Interface().(Object); ok && v.Kind() == reflect.Uint32 {
		return enum != nil && enum.CRC() != 0
	}

	switch k := v.Kind(); k { //nolint:exhaustive // have default
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

func appendMany[T any](b ...[]T) []T { //cover:ignore
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

func as118[T error](err error) (T, bool) { //cover:ignore
	var converted T

	return converted, errors.As(err, &converted)
}

type convertibleNum interface {
	constraints.Integer | constraints.Float
}

func convertNumErr[V, T convertibleNum](res T, err error) (V, error) { return V(res), err } //cover:ignore

type convertibleStr interface {
	~string | ~[]byte
}

func convertStrErr[V, T convertibleStr](res T, err error) (V, error) { return V(res), err } //cover:ignore

func ptr[T any](value T) *T { return &value }     //cover:ignore
func val[T any](value *T) T { return *value }     //cover:ignore
func new[T any]() T         { var t T; return t } //cover:ignore

func u32b(order binary.ByteOrder, v uint32) []byte { //cover:ignore
	b := make([]byte, WordLen)
	order.PutUint32(b, v)

	return b
}

func u64b(order binary.ByteOrder, v uint64) []byte { //cover:ignore
	b := make([]byte, LongLen)
	order.PutUint64(b, v)

	return b
}

func f64b(order binary.ByteOrder, v float64) []byte { //cover:ignore
	b := make([]byte, LongLen)
	order.PutUint64(b, math.Float64bits(v))

	return b
}

// allowed bitsize is 8, 16, 32, 64, 128, 256, 512, 1024 and 2048 bits.
func bigIntBytes(v *big.Int, bitsize int) []byte {
	bitsizeSquaredFloat := math.Log2(float64(bitsize))
	if _, f := math.Modf(bitsizeSquaredFloat); f != 0 {
		panic(fmt.Sprintf("bitsize not squaring by 2: bitsize %v", bitsize))
	}
	if bitsizeSquaredFloat > 11 {
		panic(fmt.Sprintf("bitsize is larger than 2048 bit: bitsize %v", bitsize))
	}

	vbytes := v.Bytes()
	vbytesLen := len(vbytes)

	offset := bitsize/8 - vbytesLen
	if offset < 0 {
		panic(fmt.Sprintf("bitsize too small: have %v, want at least %v", bitsize, vbytes))
	}

	return append(make([]byte, offset), vbytes...)
}

func maybeNil(v reflect.Value) bool {
	switch v.Kind() { //nolint:exhaustive // has default statement
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
