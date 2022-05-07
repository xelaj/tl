// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE.md for details.

package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name:   packageName,
	Action: Action,
	Description: "Parses tl schema and generates go code for it",
	Flags: []cli.Flag{

	},
	Copyright: "Xelaj software ltd.",

	Reader:    os.Stdin,
	Writer:    os.Stdout,
	ErrWriter: os.Stderr,
}
