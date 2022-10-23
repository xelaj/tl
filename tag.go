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

type structTag struct {
	BitFlags  *bitflag
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
func ParseTag(tag, defaultName string) (t *structTag, err error) {
	parts := strings.Split(tag, ",")
	name := parts[0]
	if name == "" {
		name = defaultName
	}

	t = &structTag{}
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

func (t *structTag) Ignore() bool { return t.Name == "-" }

func (t *structTag) String() string {
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

func (t *structTag) valid() error {
	if t == nil {
		return nil
	}
	if t.Implicit && (t.BitFlags == nil || t.BitFlags.TargetField == "") {
		return ErrImplicitNoTarget
	}

	if f := t.BitFlags; f != nil {
		if f.BitPosition > maxBitflagIndex {
			return ErrBitflagTooHigh
		}
		// other checks for omitempty tag goes here
	}

	return nil
}

type bitflag struct {
	TargetField string
	BitPosition uint8
}

func (t *bitflag) String() string {
	return strings.Join([]string{
		omitemptyPrefix, t.TargetField, strconv.Itoa(int(t.BitPosition)),
	}, ":")
}

const omitemptyParts = 3

func parseOmitemptyTag(opt string) (*bitflag, error) {
	parts := strings.Split(opt, ":")

	if len(parts) != omitemptyParts {
		return nil, errors.Wrap(ErrInvalidTagFormat, omitemptyPrefix)
	}

	pos, err := parseUintMax32(parts[2])
	if err != nil {
		return nil, errors.Wrap(err, omitemptyPrefix)
	}

	return &bitflag{
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
