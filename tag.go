// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"fmt"
	"strconv"
	"strings"
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
	res := t.TargetField + ":" + strconv.FormatUint(uint64(t.BitPosition), 10)
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
		return fmt.Errorf("trigger bit is too high: %v", t.BitPosition)
	}
	if t.Implicit && t.TargetField == "" {
		return fmt.Errorf("%q defined without target field to trigger", implicitFlag)
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
			if len(parts) != 3 {
				return nil, fmt.Errorf("%v: invalid flag format, got %q", omitemptyPrefix, option)
			}

			if optionalField == nil {
				optionalField = &bitflag{}
			}

			optionalField.TargetField = parts[1]
			pos, err := strconv.ParseUint(parts[2], 10, 5) //nolint:gomnd // obvious
			if err != nil {
				return nil, fmt.Errorf("%v: invalid flag format, got %q", omitemptyPrefix, option)
			}
			optionalField.BitPosition = uint8(pos)

		default:
			return nil, fmt.Errorf("invalid option %q", option)
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
