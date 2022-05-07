// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"fmt"
	"math"
	"math/big"
)

// allowed bitsize is 8, 16, 32, 64, 128, 256, 512, 1024 and 2048 bits
func bigIntBytes(v *big.Int, bitsize int) []byte {
	bitsizeSquaredFloat := math.Log2(float64(bitsize))
	if !isInt(bitsizeSquaredFloat) {
		panic(fmt.Errorf("bitsize not squaring by 2: bitsize %v", bitsize))
	}
	if bitsizeSquaredFloat > 11 {
		panic(fmt.Errorf("bitsize is larger than 2048 bit: bitsize %v", bitsize))
	}

	vbytes := v.Bytes()
	vbytesLen := len(vbytes)

	offset := bitsize/8 - vbytesLen
	if offset < 0 {
		panic(fmt.Errorf("bitsize too small: have %v, want at least %v", bitsize, vbytes))
	}

	return append(make([]byte, offset), vbytes...)
}

func isInt(value float64) bool {
	_, f := math.Modf(value)
	return f == 0
}

func littleUint24Bytes(v int) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
	}
}

// https://go.dev/play/p/uU0LiPVCPnI
func pad(lengthOfSizeLength, wordLen, msgLen int) int {
	wMinMod := wordLen - ((lengthOfSizeLength + msgLen) % wordLen)
	return (1 - wMinMod/wordLen) * wMinMod
}
