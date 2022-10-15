// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

type Schema struct { //nolint:govet // it's ok
	Objects      []*Object
	TypeComments map[string]string
}

func (s *Schema) String() string { return "" }
