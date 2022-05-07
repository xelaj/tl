// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name:        packageName,
	Action:      Action,
	Description: "Parses tl schema and generates go code for it",
	Commands: []*cli.Command{
		{
			Name:        "gen",
			Action:      GenAction,
			Description: "generates golang structures from typelang schema",
			Flags:       []cli.Flag{},
		},
	},
	Flags:     []cli.Flag{},
	Copyright: "Xelaj software ltd.",

	Reader:    os.Stdin,
	Writer:    os.Stdout,
	ErrWriter: os.Stderr,
}

func GenAction(_ *cli.Context) error {
	return nil
}
