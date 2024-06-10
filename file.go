package main

import (
	"flag"
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
	for _, a := range flag.Args() {
		if !strings.HasPrefix(a, "-") {
			filename = a
			break
		}
	}

	if len(os.Args) == 1 {
		fmt.Println("evi: try 'evi -help' for more information")
		os.Exit(0)
	}

	if filename == "" {
		fmt.Printf(eviError, "no file specified")
		os.Exit(1)
	}
}
