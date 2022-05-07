// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"math/big"
)

// Int128 is alias-like type for fixed size of big int (1024 bit value). It using only for tl objects encoding
// cause native big.Int isn't supported for en(de)coding
type Int128 struct {
	big.Int
}

// NewInt128 creates int128 with zero value
func NewInt128(value int) *Int128 {
	i := big.NewInt(int64(value))
	return &Int128{*i}
}

// NewInt128 creates int128 with random value
func RandomInt128() *Int128 {
	i := big.NewInt(0).SetBytes(randBytes(Int128Len))
	return &Int128{*i}
}

// MarshalTL implements tl marshaler from this package. Just don't use it by your hands, tl.Encoder does all
// what you need
func (i *Int128) MarshalTL(e *Encoder) error {
	e.PutRawBytes(bigIntBytes(&i.Int, Int128Len*bitsInByte))
	return nil
}

// UnmarshalTL implements tl unmarshaler from this package. Just don't use it by your hands, tl.Decoder does
// all what you need
func (i *Int128) UnmarshalTL(d *Decoder) error {
	val, err := d.peek(Int128Len)
	if err != nil {
		return err
	}
	d.success()
	v := big.NewInt(0).SetBytes(val)
	i.Int = *v
	return nil
}

// Int256 is alias-like type for fixed size of big int (2048 bit value). It using only for tl objects encoding
// cause native big.Int isn't supported for en(de)coding
type Int256 struct {
	big.Int
}

// NewInt256 creates int256 with zero value
func NewInt256(value int) *Int256 {
	i := big.NewInt(int64(value))
	return &Int256{*i}
}

// NewInt256 creates int256 with random value
func RandomInt256() *Int256 {
	i := big.NewInt(0).SetBytes(randBytes(Int256Len))
	return &Int256{*i}
}

// MarshalTL implements tl marshaler from this package. Just don't use it by your hands, tl.Encoder does all
// what you need
func (i *Int256) MarshalTL(e *Encoder) error {
	e.PutRawBytes(bigIntBytes(&i.Int, Int256Len*bitsInByte))
	return nil
}

// UnmarshalTL implements tl unmarshaler from this package. Just don't use it by your hands, tl.Decoder does
// all what you need
func (i *Int256) UnmarshalTL(d *Decoder) error {
	val, err := d.peek(Int256Len)
	if err != nil {
		return err
	}
	d.success()
	v := big.NewInt(0).SetBytes(val)
	i.Int = *v
	return nil
}
