package util

import (
	"context"
	"errors"
	"io"
	"os"
	"os/signal"

	"dario.cat/mergo"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func defaultLogger(w io.Writer) zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{Out: w}).With().Timestamp().Logger()
}

// RunCmd is a helper function for cobra.Command.PreRun that loads flags
// from environment and merges them with flags from command line.
func Run[T any](action func(ctx context.Context, flags AppCtx[T]) int, flags *T, envPath *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		log := defaultLogger(os.Stderr)

		if envPath != nil && *envPath != "" {
			if err := godotenv.Load(*envPath); err != nil {
				log.Warn().Err(err).Msg("failed to load .env file")
			}
		}

		var envFlags T
		if err := env.Parse(&envFlags); err != nil {
			log.Fatal().Err(err).Msg("failed to parse environment variables")
			return
		}

		if err := mergo.Merge(flags, &envFlags); err != nil {
			log.Fatal().Err(err).Msg("failed to merge flags")
			return
		}

		code := action(cmd.Context(), AppCtx[T]{
			IsPipeline: isPipeline(os.Stdin),
			Stdin:      os.Stdin,
			Stdout:     os.Stdout,
			Log:        log,
			Flags:      *flags,
		})
		os.Exit(code)
	}
}

func GetInput[T any](app AppCtx[T], pathFlag string) (io.Reader, error) {
	if app.IsPipeline {
		return app.Stdin, nil
	} else if pathFlag == "" {
		return nil, errors.New("you must specify file flag when using pipeline")
	}

	return os.Open(pathFlag)
}

// newContext creates a new context with all wrappers, including:
// - context attached to signals
// - logger injected in context
func NewContext(stdin io.Reader, stdout, stderr io.Writer) (context.Context, context.CancelFunc) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	return ctx, cancel
}

func isPipeline(in *os.File) bool {
	stat, _ := in.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

type AppCtx[T any] struct {
	IsPipeline bool
	Stdin      io.Reader
	Stdout     io.Writer
	Log        zerolog.Logger
	Flags      T
}
