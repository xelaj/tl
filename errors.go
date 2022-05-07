// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"fmt"
	"reflect"
)

type ErrOutIsNil null

func (ErrOutIsNil) Error() string {
	return "can't unmarshal to nil value"
}

type ErrRegisteredObjectNotFound struct {
	Data []byte
	Crc  uint32
}

func (e ErrRegisteredObjectNotFound) Error() string {
	return fmt.Sprintf("object with provided crc not registered: 0x%08x", e.Crc)
}

type ErrMustParseSlicesExplicitly null

func (ErrMustParseSlicesExplicitly) Error() string {
	return "got vector CRC code when parsing unknown object: vectors can't be parsed as predicted objects"
}

type ErrorPartialWrite struct {
	Has  int
	Want int
}

func (e ErrorPartialWrite) Error() string {
	return fmt.Sprintf("write failed: writed only %v bytes, expected %v", e.Has, e.Want)
}

type ErrImplicitInteger null

func (ErrImplicitInteger) Error() string {
	return "got int or uint, must be converted to int32, int64 or uint32 explicitly"
}

type ErrUnsupportedInt struct {
	Kind reflect.Kind
}

func (e ErrUnsupportedInt) Error() string {
	return fmt.Sprintf("only int32, int64 or uint32 are allowed, %v is unsupported by protocol", e.Kind)
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

type ErrUnexpectedNil null

func (ErrUnexpectedNil) Error() string {
	return "unexpected nil value"
}

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
