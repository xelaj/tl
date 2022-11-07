// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"fmt"
	"math/big"
	"reflect"
)

type crc32 = uint32

//nolint:gochecknoglobals // types must be global
var (
	byteSliceTyp = reflect.TypeOf((*[]byte)(nil)).Elem()
	objectTyp    = reflect.TypeOf((*Object)(nil)).Elem()
	enumTyp      = reflect.TypeOf((*Enum)(nil)).Elem()
	uint32Typ    = reflect.TypeOf((*uint32)(nil)).Elem()
)

// Object is default interface, which ANY struct must implement to decode it in
// tl format.
type Object interface {
	CRC() crc32
}

// Enum is an interface which implementations are ONLY objects without any
// fields. If enum is not set (null in TL terms), it MUST return zero value in
// CRC() method.
type Enum interface {
	Object
	fmt.Stringer
}

type Marshaler interface {
	MarshalTL(MarshalState) error
}

type Unmarshaler interface {
	UnmarshalTL(UnmarshalState) error
}

const (
	MapCrcKey = "_crc"

	bitsInByte = 8 // cause we don't want store magic numbers

	WordLen   = 32 / bitsInByte  // length of word in tl is 32 bits
	LongLen   = 64 / bitsInByte  // int64 size in bytes
	DoubleLen = 64 / bitsInByte  // float64 size in bytes
	Int128Len = 128 / bitsInByte // int128 size in bytes
	Int256Len = 256 / bitsInByte // int256 size in bytes

	// according to typelang specification, there is no able to encode value in
	// 1 byte more than 254 (max is 253). why it's 254? Idk. Why not 255 is
	// value extension? I. DON'T. KNOW. So here we are, it's a fucking magic
	// number.
	fuckingMagicNumber = 0xfe

	// https://core.telegram.org/schema/mtproto
	crcVector = 0x1cb5c415
	crcFalse  = 0xbc799737
	crcTrue   = 0x997275b5
	crcNull   = 0x56730bcc
)

func boolToCRC(v bool) crc32 { //revive:disable:flag-parameter // no, it's not
	if v {
		return crcTrue
	}

	return crcFalse
}

// Int128 is alias-like type for fixed size of big int (1024 bit value). It
// using only for tl objects encoding cause native big.Int isn't supported for
// en(de)coding.
type Int128 struct{ big.Int }

var _ interface {
	Marshaler
	Unmarshaler
} = (*Int128)(nil)

// NewInt128 creates int128 with zero value.
func NewInt128(value int) *Int128 { return &Int128{val(big.NewInt(int64(value)))} }

// RandomInt128 creates int128 with random value.
func RandomInt128() *Int128 { return &Int128{val(big.NewInt(0).SetBytes(randBytes(Int128Len)))} }

// MarshalTL implements tl marshaler from this package. Just don't use it by
// your hands, tl.Encoder does all what you need.
func (i *Int128) MarshalTL(s MarshalState) error {
	_, err := s.Write(bigIntBytes(&i.Int, Int128Len*bitsInByte))

	return err
}

// UnmarshalTL implements tl unmarshaler from this package. Just don't use it by
// your hands, tl.Decoder does all what you need.
func (i *Int128) UnmarshalTL(s UnmarshalState) error {
	v, err := s.Peek(0, Int128Len)
	if err != nil {
		return err
	}
	s.SkipBytes(Int128Len)
	i.Int = val(big.NewInt(0).SetBytes(v))

	return nil
}

// Int256 is alias-like type for fixed size of big int (2048 bit value). It
// using only for tl objects encoding cause native big.Int isn't supported for
// en(de)coding.
type Int256 struct{ big.Int }

var _ interface {
	Marshaler
	Unmarshaler
} = (*Int256)(nil)

// NewInt256 creates int256 with zero value.
func NewInt256(value int) *Int256 { return &Int256{val(big.NewInt(int64(value)))} }

// RandomInt256 creates int256 with random value.
func RandomInt256() *Int256 { return &Int256{val(big.NewInt(0).SetBytes(randBytes(Int256Len)))} }

// MarshalTL implements tl marshaler from this package. Just don't use it by
// your hands, tl.Encoder does all what you need.
func (i *Int256) MarshalTL(s MarshalState) error {
	_, err := s.Write(bigIntBytes(&i.Int, Int256Len*bitsInByte))

	return err
}

// UnmarshalTL implements tl unmarshaler from this package. Just don't use it by
// your hands, tl.Decoder does all what you need.
func (i *Int256) UnmarshalTL(s UnmarshalState) error {
	v, err := s.Peek(0, Int256Len)
	if err != nil {
		return err
	}
	s.SkipBytes(Int256Len)
	i.Int = val(big.NewInt(0).SetBytes(v))

	return nil
}

// AnyConstructor это специальный тип для декодера, который позволяет ограничить
// доступный набор конструкторов.
type AnyConstructor[T Object] struct {
	Constructor func(crc32) T
	Result      T
}

func (b *AnyConstructor[T]) UnmarshalTL(r UnmarshalState) error {
	crc, err := r.PopCRC()
	if err != nil {
		return errReadCRC{err}
	}

	obj := b.Constructor(crc)
	if Object(obj) == nil {
		return ErrObjectNotRegistered(crc)
	}

	// r.Unmarshal(obj)
	return nil
}
