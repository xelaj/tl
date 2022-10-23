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

func unreachable()   { panic("Unreachable") }   //nolint:deadcode // why not?
func unimplemented() { panic("Unimplemented") } //nolint:deadcode // why not?

func randBytes(size int) []byte {
	b := make([]byte, size)
	rand.Read(b)

	return b
}

func littleUint24Bytes(v int) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),  //nolint:gomnd // r u kiddin
		byte(v >> 16), //nolint:gomnd // r u kiddin
	}
}

func pad(l, padSize int) int {
	mod := l % padSize

	return (padSize - mod) * (1 + (mod-padSize)/padSize)
}

// v is bool, pointer, interface or slice.
func isFieldContainsData(v reflect.Value) bool {
	// special cases for enums
	if v, ok := v.Interface().(Enum); ok {
		return v.CRC() != 0
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

func convertNumErr[V, T convertibleNum](res T, err error) (V, error) { return V(res), err }

type convertibleStr interface {
	~string | ~[]byte
}

func convertStrErr[V, T convertibleStr](res T, err error) (V, error) { return V(res), err }

type errorConst string

func (e errorConst) Error() string { return string(e) }

func ptr[T any](value T) *T { return &value }
func val[T any](value *T) T { return *value }

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
