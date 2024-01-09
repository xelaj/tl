package fmt

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/xelaj/tl/cmd/tlgen/util"
	"github.com/xelaj/tl/schema"
)

// rootCmd is cobra cli entry point of the application.
func Cmd() *cobra.Command {
	var flags Flags
	var envPath string

	cmd := &cobra.Command{
		Use:  "fmt",
		Long: "",
		Run:  util.Run(Action, &flags, &envPath),
	}
	cmd.Flags().StringVar(&envPath, "env", "", "Path to .env file")

	cmd.Flags().StringVar(&flags.File, "file", "", "Path to schema file")
	cmd.MarkFlagDirname("file")

	return cmd
}

type Flags struct {
	File string
}

func Action(ctx context.Context, app util.AppCtx[Flags]) int {
	in, err := util.GetInput(app, app.Flags.File)
	if err != nil {
		app.Log.Fatal().Err(err).Msg("failed to open schema")
		return 1
	}

	schema, err := schema.Parse(app.Flags.File, in)
	if err != nil {
		app.Log.Fatal().Err(err).Msg("failed to parse schema")
		return 1
	}

	if _, err := app.Stdout.Write([]byte(schema.String())); err != nil {
		app.Log.Fatal().Err(err).Msg("failed to write schema")
		return 1
	}

	return 0
}
