package main

import (
	"bufio"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/notwithering/sgr"
)

// options
var (
	algorithm, mode int
	key             []byte
)

const (
	eviError    string = sgr.FgHiRed + "error:" + sgr.Reset + " %s\n"
	eviInfo     string = sgr.FgHiBlue + "::" + sgr.Reset + " %s\n"
	eviInfoPair string = sgr.FgHiMagenta + "::" + sgr.Reset + " %s " + sgr.FgHiBlack + ":" + sgr.Reset + " %s\n"
	eviChoice   string = "   %d) %s\n"
	eviInput    string = ">> "
)

func main() {
	editor, err := envy.MustGet("EDITOR")
	if err != nil {
		fmt.Printf(eviError, "no editor specified in $EDITOR")
		return
	}

	var filename string
	for _, a := range os.Args[1:] {
		if !strings.HasPrefix(a, "-") {
			filename = a
			break
		}
	}

	if filename == "" {
		fmt.Printf(eviError, "no file specified")
		return
	}

chooseKey:
	for {
		fmt.Printf(eviInfo, "Encryption key:")
		fmt.Printf(eviInfo, "[d]etails   [a]lgorithm   [m]ode")
		fmt.Print(eviInput)

		in, err := line()
		if err != nil {
			fmt.Printf(eviError, err)
			return
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
		case "m":
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
		default:
			key = []byte(in)
			break chooseKey
		}

		fmt.Print("\n")
	}

	if fileExists(filename) {
		if err := decryptFile(filename); err != nil {
			fmt.Printf(eviError, err)
			return
		}
	}

	cmd := exec.Command(editor, os.Args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	if fileExists(filename) {
		if err := encryptFile(filename); err != nil {
			fmt.Printf(eviError, err)

			fmt.Print("\n")

			fmt.Printf(eviInfo, "Remove file? [Y/n]")
			fmt.Print(eviInput)

			in, err := line()
			if err != nil {
				fmt.Printf(eviError, err)
				return
			}

			switch strings.ToLower(in) {
			case "", "y":
				if err := os.Remove(filename); err != nil {
					fmt.Printf(eviError, err)
					return
				}
			default:
				return
			}
		}
	}
}

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

func decryptFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	file.Close()

	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	c, err := getBlock()
	if err != nil {
		return err
	}

	switch modes[mode] {
	case "GCM":
		gcm, err := cipher.NewGCM(c)
		if err != nil {
			return err
		}

		nonceSize := gcm.NonceSize()

		if len(b) < nonceSize {
			return fmt.Errorf("cipher text is smaller than the nonce size")
		}

		nonce, cipherBytes := b[:nonceSize], b[nonceSize:]
		plainText, err := gcm.Open(nil, nonce, cipherBytes, nil)
		if err != nil {
			return err
		}

		if _, err := file.Write(plainText); err != nil {
			return err
		}
	}

	return nil
}

func encryptFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	file.Close()

	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	c, err := getBlock()
	if err != nil {
		return err
	}

	switch modes[mode] {
	case "GCM":
		gcm, err := cipher.NewGCM(c)
		if err != nil {
			return err
		}

		nonce := make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return err
		}

		encrypted := gcm.Seal(nonce[:], nonce[:], b, nil)

		if _, err := file.Write(encrypted); err != nil {
			return err
		}
	}

	return nil
}

func hash256(key []byte) []byte {
	h := sha256.New()
	h.Write(key)
	return h.Sum(nil)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
