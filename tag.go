// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// tag name of struct tags.
const (
	tagName = "tl"

	implicitFlag    = "implicit"
	isBitflagFlag   = "bitflag"
	omitemptyPrefix = "omitempty"

	maxBitflagIndex = 32
)

type StructTag struct {
	BitFlags  *Bitflag
	Implicit  bool
	Name      string
	IsBitflag bool
}

// ParseTag is a function which parses struct field tag for structes, defined by
// you.
//
// how tag could look:
//
//	// field named somefield, it could be empty, value is boolean, so it can
//	// be encoded implicitly, value stored in some_flag value
//	`tl:"somefield,omitempty:some_flag:3,implicit"`
//
//	// value is abcd, it's required to exist in serialized
//	`tl:"abcd"`
func ParseTag(tag, defaultName string) (t *StructTag, err error) {
	parts := strings.Split(tag, ",")
	name := parts[0]
	if name == "" {
		name = defaultName
	}

	t = &StructTag{
		Name: name,
	}
	for _, option := range parts[1:] {
		switch {
		case option == implicitFlag:
			t.Implicit = true

		case option == isBitflagFlag:
			t.IsBitflag = true

		case strings.HasPrefix(option, omitemptyPrefix):
			if t.BitFlags, err = parseOmitemptyTag(option); err != nil {
				return nil, err
			}

		default:
			return nil, errors.Wrap(ErrInvalidTagOption, option)
		}
	}
	if err := t.valid(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *StructTag) Ignore() bool { return t.Name == "-" } //cover:ignore

func (t *StructTag) String() string { //cover:ignore
	res := t.Name
	if t.BitFlags != nil {
		res += "," + t.BitFlags.String()
	}
	if t.Implicit {
		res += "," + implicitFlag
	}
	if t.IsBitflag {
		res += "," + isBitflagFlag
	}

	return res
}

func (t *StructTag) valid() error {
	switch {
	case t == nil:
		return nil
	case t.Name == "":
		return ErrTagNameEmpty
	case t.Implicit && (t.BitFlags == nil || t.BitFlags.TargetField == ""):
		return ErrImplicitNoTarget
	case t.Implicit && t.IsBitflag:
		return ErrImplicitBitflag
	case t.BitFlags != nil && t.BitFlags.BitPosition > maxBitflagIndex:
		return ErrBitflagTooHigh
	default:
		return nil
	}
}

type Bitflag struct {
	TargetField string
	BitPosition uint8
}

func (t *Bitflag) String() string {  //cover:ignore
	return strings.Join([]string{
		omitemptyPrefix, t.TargetField, strconv.Itoa(int(t.BitPosition)),
	}, ":")
}

const omitemptyParts = 3

func parseOmitemptyTag(opt string) (*Bitflag, error) {
	parts := strings.Split(opt, ":")

	if len(parts) != omitemptyParts {
		return nil, errors.Wrap(ErrInvalidTagFormat, omitemptyPrefix)
	}

	pos, err := parseUintMax32(parts[2])
	if err != nil {
		return nil, errors.Wrap(err, omitemptyPrefix)
	}

	return &Bitflag{
		TargetField: parts[1],
		BitPosition: pos,
	}, nil
}

const (
	bit32       = 5  // 5 bits to make 32 different variants
	defaultBase = 10 // base 10 of numbers
)

func parseUintMax32(s string) (uint8, error) {
	if pos, err := strconv.ParseUint(s, defaultBase, bit32); err == nil {
		return uint8(pos), nil
	}

	return 0, ErrInvalidTagFormat
}
