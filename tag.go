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

// how tag could look:
//
//   // field named somefield, it could be empty, value is boolean, so it can
//   // be encoded implicitly, value stored in some_flag value
//   `tl:"somefield,omitempty:some_flag:3,implicit"`
//
//   // value is abcd, it's required to exist in serialized
//   `tl:"abcd"`
type structTag struct {
	BitFlags  *bitflag
	Name      string
	IsBitflag bool
}

func (t *structTag) Ignore() bool   { return t.Name == "-" }
func (t *structTag) Implicit() bool { return t.BitFlags != nil && t.BitFlags.Implicit }

func (t *structTag) String() string {
	res := t.Name
	if t.BitFlags != nil {
		res += "," + t.BitFlags.String()
	}
	if t.IsBitflag {
		res += "," + isBitflagFlag
	}

	return res
}

type bitflag struct {
	TargetField string
	BitPosition uint8
	Implicit    bool
}

func (t *bitflag) String() string {
	res := t.TargetField + ":" + strconv.Itoa(int(t.BitPosition))
	if t.Implicit {
		res += "," + implicitFlag
	}

	return res
}

func (t *bitflag) valid() error {
	if t == nil {
		return nil
	}
	if t.BitPosition >= maxBitflagIndex {
		return ErrBitflagTooHigh
	}
	if t.Implicit && t.TargetField == "" {
		return errors.New(implicitFlag + " defined without target field to trigger")
	}

	return nil
}

func parseTag(tag, defaultName string) (*structTag, error) {
	parts := strings.Split(tag, ",")
	name := parts[0]
	if name == "" {
		name = defaultName
	}

	var optionalField *bitflag
	var isBitflag bool
	for _, option := range parts[1:] {
		switch {
		case option == implicitFlag:
			if optionalField == nil {
				optionalField = &bitflag{}
			}
			optionalField.Implicit = true

			continue

		case option == isBitflagFlag:
			isBitflag = true

		case strings.HasPrefix(option, omitemptyPrefix):
			parts := strings.Split(option, ":")

			const omitemptyParts = 3
			if len(parts) != omitemptyParts {
				return nil, errors.Wrap(ErrInvalidTagFormat, omitemptyPrefix)
			}

			if optionalField == nil {
				optionalField = &bitflag{}
			}

			optionalField.TargetField = parts[1]
			pos, err := strconv.ParseUint(parts[2], 10, 5) //nolint:gomnd // obvious
			if err != nil {
				return nil, errors.Wrap(ErrInvalidTagFormat, omitemptyPrefix)
			}
			optionalField.BitPosition = uint8(pos)

		default:
			return nil, errors.Wrap(ErrInvalidTagOption, option)
		}
	}
	if err := optionalField.valid(); err != nil {
		return nil, err
	}

	return &structTag{
		Name:      name,
		BitFlags:  optionalField,
		IsBitflag: isBitflag,
	}, nil
}
