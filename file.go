package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/k0kubun/go-ansi"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func getFilename() {
	filename = flag.Arg(0)

	if len(os.Args) == 1 {
		fmt.Println("evi: try 'evi -help' for more information")
		os.Exit(0)
	}

	if filename == "" {
		ansi.Printf(eviError, "no file specified")
		os.Exit(1)
	}
}

func decryptFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	plainText, err := decryptBytes(b)
	if err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.Write(plainText); err != nil {
		return err
	}

	return nil
}

func encryptFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	encrypted, err := encryptBytes(b)
	if err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.Write(encrypted); err != nil {
		return err
	}

	return nil
}
