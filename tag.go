package tl

import (
	"fmt"
	"strconv"
	"strings"
)

type structTag struct {
	Name      string
	BitFlags  *bitflag
	IsBitflag bool
}

const isBitflagOption = "bitflag"

func (t *structTag) Ignore() bool   { return t.Name == "-" }
func (t *structTag) Implicit() bool { return t.BitFlags != nil && t.BitFlags.Implicit }

func (t *structTag) String() string {
	res := t.Name
	if t.BitFlags != nil {
		res += "," + t.BitFlags.String()
	}
	if t.IsBitflag {
		res += "," + isBitflagOption
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
	if t.BitPosition >= 32 {
		return fmt.Errorf("trigger bit is too high: %v", t.BitPosition)
	}
	if t.Implicit && t.TargetField == "" {
		return fmt.Errorf("%q defined without target field to trigger", implicitFlag)
	}

	return nil
}

// how tag could look:
//
//   // field named somefield, it could be empty, value is boolean, so it can be encoded implicitly
//   `tl:"somefield,omitempty:bitflag:3,implicit"`
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
			pos, err := strconv.ParseUint(parts[2], 10, 5)
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
