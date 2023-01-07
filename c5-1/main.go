package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}

	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	fmt.Printf("killing server with pid=%d\n", pid)

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can't remove pid file - %s", err)
	}

	return nil
}
