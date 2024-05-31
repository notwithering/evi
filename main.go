package main

import (
	"flag"
	"fmt"
)

var (
	noDecrypt bool
	noEdit    bool
	noEncrypt bool
)

func main() {
	var keyFlag string
	flag.StringVar(&keyFlag, "key", "", "")
	flag.StringVar(&keyFlag, "k", "", "")

	flag.BoolVar(&noDecrypt, "no-decrypt", noDecrypt, "")
	flag.BoolVar(&noEdit, "no-edit", noEdit, "")
	flag.BoolVar(&noEncrypt, "no-encrypt", noEncrypt, "")

	flag.Usage = func() {
		fmt.Println("Usage: evi [options...] <file>")
		fmt.Println("  -k, -key          Preset the encryption key")
		fmt.Println("      -no-decrypt   Stop the program from decrypting the file")
		fmt.Println("      -no-edit      Simply decrypt then re-encrypt the file")
		fmt.Println("      -no-encrypt   Stop the program from re-encrypting the file")
	}

	flag.Parse()

	if !noEdit {
		getEditor()
	}
	getFilename()
	if keyFlag == "" {
		chooseKey()
	} else {
		key = []byte(keyFlag)
		burnString(&keyFlag)
	}
	if !noDecrypt {
		decrypt()
	}
	if !noEdit {
		openEditor()
	}
	if !noEncrypt {
		encrypt()
	}
}
