// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"errors"
	"fmt"
	"reflect"
)

type errorConst string

func (e errorConst) Error() string { return string(e) } //cover:ignore

const (
	ErrUnexpectedNil = errorConst("unexpected nil value")
	ErrImplicitInt   = errorConst("value must be converted to int32, int64 or uint32 explicitly")

	ErrBitflagTooHigh   = errorConst("trigger bit is more than 32")
	ErrImplicitNoTarget = errorConst(implicitFlag + " defined without target field to trigger")
	ErrImplicitBitflag  = errorConst("'" + implicitFlag + "' and '" + isBitflagFlag + "' can't be combined")
	ErrTagNameEmpty     = errorConst("tag name is empty")
	ErrInvalidTagOption = errorConst("invalid option")
	ErrInvalidTagFormat = errorConst("invalid tag format")
	ErrTypeIsEmpty      = errorConst("tl field type is empty")
)

type ErrRegisteredObjectNotFound struct {
	Data []byte
	Crc  uint32
}

func (e ErrRegisteredObjectNotFound) Error() string { //cover:ignore
	return fmt.Sprintf("object with provided crc not registered: 0x%08x", e.Crc)
}

type ErrPartialWrite struct {
	Has  int
	Want int
}

func (e ErrPartialWrite) Error() string { //cover:ignore
	return fmt.Sprintf("write failed: writed only %v bytes, expected %v", e.Has, e.Want)
}

type ErrUnsupportedType struct {
	Type reflect.Type
}

func (e ErrUnsupportedType) Error() string { //cover:ignore
	switch k := e.Type.Kind(); k {
	// supported, but TL doesn't support 8 and 16 bit numbers
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64:
		return fmt.Sprintf("int32, int64 or uint32 allowed, %v is unsupported by protocol", k)

	// same: supported, but not in TL, so we can't understand, how much bytes
	// we need to scan.
	case reflect.Float32, reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("only float64 is allowed, %v is unsupported by protocol", k)

	default:
		return fmt.Sprintf("invalid or unknown type %v", e.Type)
	}
}

type ErrObjectNotRegistered crc32

func (e ErrObjectNotRegistered) Error() string { //cover:ignore
	return fmt.Sprintf("object with crc 0x%08x not registered", crc32(e))
}

type ErrInvalidCRC struct {
	Got  crc32
	Want crc32
}

func (e ErrInvalidCRC) Error() string { //cover:ignore
	return fmt.Sprintf("invalid crc code: got 0x%08x; want 0x%08x", e.Got, e.Want)
}

type ErrInvalidBoolCRC crc32

func (e ErrInvalidBoolCRC) Error() string { //cover:ignore
	return fmt.Sprintf(
		"want a 0x%08x (true) or 0x%08x (false); got 0x%08x",
		crcTrue,
		crcFalse,
		crc32(e),
	)
}

type errReadCRC struct {
	Wrapped error
}

func (e errReadCRC) Error() string { return "read crc: " + e.Wrapped.Error() } //cover:ignore
func (e errReadCRC) Unwrap() error { return e.Wrapped }                        //cover:ignore

type ErrPath struct {
	Err  error
	Path string
}

func (e ErrPath) Error() string { return e.Path + ": " + e.Err.Error() } //cover:ignore
func (e ErrPath) Unwrap() error { return e.Err }                         //cover:ignore

func wrapPath(err error, path string) error { //cover:ignore
	if err == nil {
		return nil
	}

	if e := new[ErrPath](); errors.As(err, &e) {
		e.Path += "." + path

		return e
	}

	return ErrPath{
		Path: path,
		Err:  err,
	}
}
