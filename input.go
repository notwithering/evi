package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func line() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	in, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimRight(in, "\n"), nil
}

func chooseIndex() (int, error) {
	in, err := line()
	if err != nil {
		return 0, err
	}

	inputtedIndex, err := strconv.Atoi(in)
	if err != nil {
		return 0, err
	}

	return inputtedIndex - 1, nil
}
