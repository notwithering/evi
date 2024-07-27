package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/notwithering/memory"
)

func chooseKey() {
chooseKey:
	for {
		fmt.Printf(eviInfo, "Encryption key:")
		fmt.Printf(eviInfo, "[d]etails")
		fmt.Print(eviInput)

		in, err := line(true)
		if err != nil {
			fmt.Printf(eviError, err)
			os.Exit(1)
		}

		switch strings.ToLower(in) {
		case "d":
			fmt.Print("\n")
			fmt.Printf(eviInfoPair, "Editor", editor)
			fmt.Printf(eviInfoPair, "Encryption", "AES-256-GCM")
			fmt.Printf(eviInfoPair, "File", filename)
			fmt.Printf(eviInfoPair, "Hashing", "SHA-256")
			fmt.Print("\n")
		default:
			key = hash([]byte(in))
			memory.Zero(&in)

			if confirmKey {
				fmt.Print(eviInput)

				in, err := line(true)
				if err != nil {
					fmt.Printf(eviError, err)
					os.Exit(1)
				}

				confirmedKey := hash([]byte(in))
				memory.Zero(&in)

				if !keysMatch(key, confirmedKey) {
					fmt.Printf(eviError, "keys do not match")
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
	fmt.Printf(eviInfo, "Remove file? [y/N]")
	fmt.Print(eviInput)

	in, err := line(false)
	if err != nil {
		fmt.Printf(eviError, err)
		os.Exit(1)
	}

	switch strings.ToLower(in) {
	case "y":
		if err := os.Remove(filename); err != nil {
			fmt.Printf(eviError, err)
			os.Exit(1)
		}
	default:
		return
	}
}
