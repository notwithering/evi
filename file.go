package main

import (
	"fmt"
	"os"
	"strings"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func getFilename() {
	for _, a := range os.Args[1:] {
		if !strings.HasPrefix(a, "-") {
			filename = a
			break
		}
	}

	if filename == "" {
		fmt.Printf(eviError, "no file specified")
		os.Exit(1)
	}
}
