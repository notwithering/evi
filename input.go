package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/notwithering/zero"
)

func line(password bool) (string, error) {
	if err := keyboard.Open(); err != nil {
		return "", err
	}
	defer keyboard.Close()

	var input []byte
	var position int

	var reprintFromPosition = func() {
		var distance int = len(input[position:]) + 1
		if password {
			for range distance - 1 {
				fmt.Print("*")
			}
		} else {
			fmt.Print(string(input[position:]))
		}
		fmt.Print(" ")
		for range distance {
			fmt.Print("\x1b[D")
		}
	}

	for {
		r, key, err := keyboard.GetKey()
		if err != nil {
			return "", err
		}

		switch key {
		case keyboard.KeyEnter:
			fmt.Print("\r\n")
			if password {
				defer zero.Zero(&input)
			}
			return string(input), nil
		case keyboard.KeyArrowLeft:
			if position > 0 {
				position--
				fmt.Print("\x1b[D")
			}
		case keyboard.KeyArrowRight:
			if position < len(input) {
				position++
				fmt.Print("\x1b[C")
			}
		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			if len(input) > 0 && position > 0 {
				var inputBuf []byte
				for i, r := range input {
					if i != position-1 {
						inputBuf = append(inputBuf, r)
					}
				}
				input = inputBuf
				position--
				fmt.Print("\x1b[D")

				reprintFromPosition()
			}
		case keyboard.KeyDelete:
			if len(input) > 0 && position < len(input) {
				var inputBuf []byte
				for i, r := range input {
					if i != position {
						inputBuf = append(inputBuf, r)
					}
				}
				input = inputBuf

				reprintFromPosition()
			}
		case keyboard.KeyCtrlA:
			for range position {
				fmt.Print("\x1b[D")
			}
			position = 0
		case keyboard.KeyCtrlE:
			for range len(input) - position {
				fmt.Print("\x1b[C")
			}
			position = len(input)
		// TODO: Add CTRL+U, CTRL+K, and CTRL+W
		default:
			if key == keyboard.KeySpace {
				r = ' '
			}

			if r < ' ' || r > '~' {
				continue
			}

			inputBuf := append(input, 0)
			copy(inputBuf[position+1:], inputBuf[position:])
			inputBuf[position] = byte(r)
			input = inputBuf

			if password {
				fmt.Print("*")
			} else {
				fmt.Print(string(r))
			}

			position++

			reprintFromPosition()
		}
	}
}
