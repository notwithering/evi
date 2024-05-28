package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/term"
)

// FIXME: make this more like a text box (add backspace, cursor control, keybinds, etc.)
func line(password bool) (string, error) {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	var input []byte

	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return "", err
		}

		if r == rune(tcell.KeyEnter) {
			return string(input), nil
		}

		if r < ' ' || r > '~' {
			continue
		}

		input = append(input, byte(r))
		if password {
			fmt.Print("*")
		} else {
			fmt.Print(string(r))
		}
	}
}

func chooseIndex() (int, error) {
	in, err := line(false)
	if err != nil {
		return 0, err
	}

	inputtedIndex, err := strconv.Atoi(in)
	if err != nil {
		return 0, err
	}

	return inputtedIndex - 1, nil
}
