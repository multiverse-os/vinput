package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func NewHandler() {
	shutdownSignals := make(chan bool, 1)
	go signalHandler(shutdownSignals)
}

func signalHandler(shutdownSignals <-chan bool) {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, os.Kill)
	for {
		s := <-signals
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, os.Kill:
			fmt.Println("[portal-gun:virtui][SHUTDOWN] User initiated shutdown by pressing CTRL+C...")
			os.Exit(0)
		}
	}
}
