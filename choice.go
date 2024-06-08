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

			break chooseKey
		}
	}
}

func removeFile() {
	fmt.Printf(eviInfo, "Remove file? [Y/n]")
	fmt.Print(eviInput)

	in, err := line(false)
	if err != nil {
		fmt.Printf(eviError, err)
		os.Exit(1)
	}

	switch strings.ToLower(in) {
	case "", "y":
		if err := os.Remove(filename); err != nil {
			fmt.Printf(eviError, err)
			os.Exit(1)
		}
	default:
		return
	}
}
