package main

import (
	"flag"
	"fmt"

	"github.com/notwithering/memory"
)

var (
	confirmKey bool
	noDecrypt  bool
	noEdit     bool
	noEncrypt  bool
)

func main() {
	var keyFlag string
	flag.BoolVar(&confirmKey, "confirm-key", confirmKey, "")
	flag.BoolVar(&confirmKey, "c", confirmKey, "")

	flag.StringVar(&keyFlag, "key", keyFlag, "")
	flag.StringVar(&keyFlag, "k", keyFlag, "")

	flag.BoolVar(&noDecrypt, "no-decrypt", noDecrypt, "")
	flag.BoolVar(&noEdit, "no-edit", noEdit, "")
	flag.BoolVar(&noEncrypt, "no-encrypt", noEncrypt, "")

	flag.Usage = func() {
		fmt.Println("Usage: evi [options...] <file>")
		fmt.Println(" -c, -confirm-key  Program asks for key twice")
		fmt.Println(" -h, -help         Show this help menu")
		fmt.Println(" -k, -key          Preset the encryption key")
		fmt.Println("     -no-decrypt   Stop the program from decrypting the file")
		fmt.Println("     -no-edit      Stop the program from opening the editor")
		fmt.Println("     -no-encrypt   Stop the program from re-encrypting the file")
	}

	flag.Parse()

	if !noEdit {
		getEditor()
	}
	getFilename()
	if keyFlag == "" {
		chooseKey()
	} else {
		key = hash([]byte(keyFlag))
		memory.Zero(&keyFlag)
	}
	if !noDecrypt {
		decrypt()
	}
	if !noEdit {
		openEditor()
	}

	// TODO: instead of only re-encryptying after the editor closes: re-encrypt after anything is written to the file.

	if !noEncrypt {
		encrypt()
	}
}
