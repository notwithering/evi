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
		fmt.Printf(eviInfo, "[d]etails   [a]lgorithm   [m]ode")
		fmt.Print(eviInput)

		in, err := line()
		if err != nil {
			fmt.Printf(eviError, err)
			os.Exit(1)
		}

		fmt.Print("\n")

		switch strings.ToLower(in) {
		case "d":
			fmt.Printf(eviInfoPair, "Algorithm", algorithms[algorithm])
			fmt.Printf(eviInfoPair, "Editor", editor)
			fmt.Printf(eviInfoPair, "File", filename)
			fmt.Printf(eviInfoPair, "Hashing", "SHA256")
			fmt.Printf(eviInfoPair, "Mode", modes[mode])
		case "a":
			chooseAlgorithm()
		case "m":
			chooseMode()
		default:
			key = []byte(in)
			break chooseKey
		}

		fmt.Print("\n")
	}
}

func chooseAlgorithm() {
chooseAlgorithm:
	for {
		fmt.Printf(eviInfo, "Algorithm:")
		for i, a := range algorithms {
			fmt.Printf(eviChoice, i+1, a)
		}

		fmt.Print(eviInput)

		index, err := chooseIndex()
		if err != nil {
			fmt.Printf(eviError, err)
			continue chooseAlgorithm
		}

		if index < 0 || index >= len(algorithms) {
			fmt.Printf(eviError, "index out of range")
			continue chooseAlgorithm
		}

		algorithm = index

		break chooseAlgorithm
	}
}

func chooseMode() {
chooseMode:
	for {
		fmt.Printf(eviInfo, "Mode:")
		for i, a := range modes {
			fmt.Printf(eviChoice, i+1, a)
		}

		fmt.Print(eviInput)

		index, err := chooseIndex()
		if err != nil {
			fmt.Printf(eviError, err)
			continue chooseMode
		}

		if index < 0 || index >= len(modes) {
			fmt.Printf(eviError, "index out of range")
			continue chooseMode
		}

		mode = index

		break chooseMode
	}
}
