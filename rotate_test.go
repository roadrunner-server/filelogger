//go:build linux

package filelogger

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Example of how to rotate in response to SIGHUP.
func ExampleLogger_Rotate() {
	l := &Logger{}
	log.SetOutput(l)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			err := l.Rotate()
			if err != nil {
				panic(err)
			}
		}
	}()
}
