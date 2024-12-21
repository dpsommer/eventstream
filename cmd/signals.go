package main

import (
	"os"
	"os/signal"
	"syscall"
)

func setupSignals() <-chan os.Signal {
	c := make(chan os.Signal, 1)

	sig := []os.Signal{
		syscall.SIGINT,
		syscall.SIGTERM,
	}

	signal.Notify(c, sig...)

	return c
}
