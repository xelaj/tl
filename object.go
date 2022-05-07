// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"fmt"
	"reflect"
)

type crc32 = uint32

var (
	byteSliceTyp   = reflect.TypeOf((*[]byte)(nil)).Elem()
	marshalerTyp   = reflect.TypeOf((*Marshaler)(nil)).Elem()
	unmarshalerTyp = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
	objectTyp      = reflect.TypeOf((*Object)(nil)).Elem()
	enumTyp        = reflect.TypeOf((*Enum)(nil)).Elem()
	uint32Type     = reflect.TypeOf((*uint32)(nil)).Elem()
)

// Object is default interface, which ANY struct must implement to decode it in tl format.
type Object interface {
	CRC() uint32
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
	MarshalTL(*Encoder) error
}

type Unmarshaler interface {
	UnmarshalTL(*Decoder) error
}

