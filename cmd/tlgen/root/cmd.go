// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package root

import (
	"github.com/spf13/cobra"

	"github.com/xelaj/tl/cmd/tlgen/root/diff"
	"github.com/xelaj/tl/cmd/tlgen/root/fmt"
	"github.com/xelaj/tl/cmd/tlgen/root/gen"
)

// rootCmd is cobra cli entry point of the application.
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "tlgen",
		Long: "",
	}

	cmd.AddCommand(diff.Cmd())
	cmd.AddCommand(fmt.Cmd())
	cmd.AddCommand(gen.Cmd())

	return cmd
}
