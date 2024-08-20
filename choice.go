package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/go-ansi"
	"github.com/notwithering/memory"
)

func chooseKey() {
chooseKey:
	for {
		ansi.Printf(eviInfo, "Encryption key:")
		ansi.Printf(eviInfo, "[d]etails")
		fmt.Print(eviInput)

		in, err := line(true)
		if err != nil {
			ansi.Printf(eviError, err)
			os.Exit(1)
		}

		switch strings.ToLower(in) {
		case "d":
			fmt.Print("\n")
			ansi.Printf(eviInfoPair, "Editor", editor)
			ansi.Printf(eviInfoPair, "Encryption", "AES-256-GCM")
			ansi.Printf(eviInfoPair, "File", filename)
			ansi.Printf(eviInfoPair, "Hashing", "SHA-256")
			fmt.Print("\n")
		default:
			key = hash([]byte(in))
			memory.Zero(&in)

			if confirmKey {
				fmt.Print(eviInput)

				in, err := line(true)
				if err != nil {
					ansi.Printf(eviError, err)
					os.Exit(1)
				}

				confirmedKey := hash([]byte(in))
				memory.Zero(&in)

				if !keysMatch(key, confirmedKey) {
					ansi.Printf(eviError, "keys do not match")
					os.Exit(1)
				}
			}

			break chooseKey
		}
	}
}

func keysMatch(key, confirmedKey []byte) bool {
	if len(key) != len(confirmedKey) {
		return false
	}
	for i, b := range key {
		if b != confirmedKey[i] {
			return false
		}
	}
	return true
}

func removeFile() {
	ansi.Printf(eviInfo, "Remove file? [y/N]")
	fmt.Print(eviInput)

	in, err := line(false)
	if err != nil {
		ansi.Printf(eviError, err)
		os.Exit(1)
	}

	switch strings.ToLower(in) {
	case "y":
		if err := os.Remove(filename); err != nil {
			ansi.Printf(eviError, err)
			os.Exit(1)
		}
	default:
		return
	}
}
