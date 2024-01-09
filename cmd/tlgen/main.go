// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package main

import (
	"os"

	"github.com/xelaj/tl/cmd/tlgen/root"
	"github.com/xelaj/tl/cmd/tlgen/util"
)

func main() {
	ctx, cancel := util.NewContext(os.Stdin, os.Stdout, os.Stderr)
	defer cancel()

	root.Cmd().ExecuteContext(ctx)
}
