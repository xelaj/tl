// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package diff

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/xelaj/tl/cmd/tlgen/util"
	"github.com/xelaj/tl/schema"
)

// rootCmd is cobra cli entry point of the application.
func Cmd() *cobra.Command {
	var flags Flags
	var envPath string

	cmd := &cobra.Command{
		Use:  "patch",
		Long: "patch comments from old schema to new",
		Run:  util.Run(Action, &flags, &envPath),
	}
	cmd.Flags().StringVar(&envPath, "env", "", "Path to .env file")

	cmd.Flags().StringVar(&flags.OldSchema, "old", "/dev/stdin", "Path to old schema file")
	cmd.MarkFlagDirname("old")
	cmd.Flags().StringVar(&flags.OldSchema, "new", "", "Path to new schema file")
	cmd.MarkFlagRequired("new")
	cmd.MarkFlagDirname("new")

	return cmd
}

type Flags struct {
	OldSchema string
	NewSchema string
}

func Action(ctx context.Context, app util.AppCtx[Flags]) int {
	oldSchemaFile, err := util.GetInput(app, app.Flags.OldSchema)
	if err != nil {
		app.Log.Fatal().Err(err).Msg("failed to open old schema")
		return 1
	}
	oldSchema, err := schema.Parse(app.Flags.OldSchema, oldSchemaFile)
	if err != nil {
		app.Log.Fatal().Err(err).Msg("failed to parse old schema")
		return 1
	}

	newSchemaFile, err := os.Open(app.Flags.NewSchema)
	if err != nil {
		app.Log.Fatal().Err(err).Msg("failed to open new schema")
		return 1
	}
	newSchema, err := schema.Parse(app.Flags.NewSchema, newSchemaFile)
	if err != nil {
		app.Log.Fatal().Err(err).Msg("failed to parse new schema")
		return 1
	}

	_, _ = oldSchema, newSchema

	panic("unimplemented")

	// if _, err := app.Stdout.Write([]byte(patch(oldSchema, newSchema).String())); err != nil {
	// 	app.Log.Fatal().Err(err).Msg("failed to write schema")
	// 	return 1
	// }
}
