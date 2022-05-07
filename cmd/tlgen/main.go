// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE.md for details.

package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := app.RunContext(initCtx(), os.Args)
	if err != nil {
		panic(err)
	}
}

func initCtx() context.Context {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sig
		cancel()
		// don't close channel, caus it may cause panic, when app receive
		// multiple signals, and writing them to closed channel
	}()
	return ctx
}
