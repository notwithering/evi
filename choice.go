package main

import (
	"fmt"
	"os"
	"strings"
)

func chooseKey() {
chooseKey:
	for {
		fmt.Printf(eviInfo, "Encryption key:")
		fmt.Printf(eviInfo, "[d]etails   [e]ncryption")
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
			fmt.Printf(eviInfoPair, "Encryption", encryptions[encryption])
			fmt.Printf(eviInfoPair, "File", filename)
			fmt.Print("\n")
		case "e":
			fmt.Print("\n")
			chooseEncryption()
			fmt.Print("\n")
		default:
			key = []byte(in)
			break chooseKey
		}
	}
}

func chooseEncryption() {
chooseEncryption:
	for {
		fmt.Printf(eviInfo, "Encryption:")
		for i, a := range encryptions {
			fmt.Printf(eviChoice, i+1, a)
		}

		fmt.Print(eviInput)

		index, err := chooseIndex()
		if err != nil {
			fmt.Printf(eviError, err)
			continue chooseEncryption
		}

		if index < 0 || index >= len(encryptions) {
			fmt.Printf(eviError, "index out of range")
			continue chooseEncryption
		}

		encryption = index

		break chooseEncryption
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
