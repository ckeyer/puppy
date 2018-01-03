package daemon

import (
	"os"
	"os/signal"
	"syscall"
)

var once bool

func WaitForExit(stopAll <-chan struct{}) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		s := <-sigChan
		switch s {
		case syscall.SIGINT, syscall.SIGTERM:
		}
	}
}
