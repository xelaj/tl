// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"fmt"
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

// Object is default interface, which ANY struct must implement to decode it in tl format.
type Object interface {
	CRC() crc32
}

const MapCrcKey = "_crc"

// Enum is an interface which implementations are ONLY objects without any fields.
//
// If enum is not set (null in TL terms), it MUST return zero value in CRC() method.
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
	bitsInByte = 8 // cause we don't want store magic numbers

	WordLen   = 32 / bitsInByte // length of word in tl is 32 bits
	LongLen   = 64 / bitsInByte // int64 size in bytes
	DoubleLen = 64 / bitsInByte // float64 size in bytes

	// according to typelang specification, there is no able to encode value in
	// 1 byte more than 253 (max is 252). why it's 253? Idk. Why not 254 and 255
	// is value extension? I. DON'T. KNOW. So here we are, it's a fucking magic
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

func IsZeroValue[T comparable](t T) bool { return t == *new(T) }
func IsZeroValueReflect(t any) bool      { return reflect.ValueOf(&t).Elem().IsZero() }
