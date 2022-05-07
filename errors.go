// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"fmt"
	"reflect"
)

type ErrRegisteredObjectNotFound struct {
	Data []byte
	Crc  uint32
}

func (e ErrRegisteredObjectNotFound) Error() string {
	return fmt.Sprintf("object with provided crc not registered: 0x%08x", e.Crc)
}

type ErrorPartialWrite struct {
	Has  int
	Want int
}

func (e ErrorPartialWrite) Error() string {
	return fmt.Sprintf("write failed: writed only %v bytes, expected %v", e.Has, e.Want)
}

type ErrUnsupportedInt struct {
	Kind reflect.Kind
}

func (e ErrUnsupportedInt) Error() string {
	return fmt.Sprintf("int32, int64 or uint32 allowed, %v is unsupported by protocol", e.Kind)
}

type ErrUnsupportedFloat struct {
	Kind reflect.Kind
}

func (e ErrUnsupportedFloat) Error() string {
	return fmt.Sprintf("only float64 is allowed, %v is unsupported by protocol", e.Kind)
}

type ErrUnsupportedType struct {
	Type reflect.Type
}

func (e ErrUnsupportedType) Error() string {
	return fmt.Sprintf("invalid or unknown type %v", e.Type)
}

type errorConst string

func (e errorConst) Error() string { return string(e) }

const (
	ErrUnexpectedNil = errorConst("unexpected nil value")
	ErrImplicitInt   = errorConst("value must be converted to int32, int64 or uint32 explicitly")

	ErrBitflagTooHigh   = errorConst("trigger bit is more than 32")
	ErrInvalidTagOption = errorConst("invalid option")
	ErrInvalidTagFormat = errorConst("invalid tag format")
)

type ErrPath struct {
	Err  error
	Path string
}

func (e ErrPath) Error() string { return e.Path + ": " + e.Err.Error() }
func (e ErrPath) Unwrap() error { return e.Err }

func wrapPath(err error, path string) error {
	if err == nil {
		return nil
	}

	if err, ok := as118[ErrPath](err); ok {
		err.Path += "." + path

		return err
	}

	return ErrPath{
		Path: path,
		Err:  err,
	}
}
