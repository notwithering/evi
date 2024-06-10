package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func getFilename() {
	for _, a := range flag.Args() {
		if !strings.HasPrefix(a, "-") {
			filename = a
			break
		}
	}

	if len(os.Args) == 1 {
		fmt.Println("evi: try 'evi -help' for more information")
		os.Exit(0)
	}

	if filename == "" {
		fmt.Printf(eviError, "no file specified")
		os.Exit(1)
	}
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

	plainText, err := decryptBytes(b)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}

	if _, err := file.Write(plainText); err != nil {
		return err
	}
	file.Close()

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

	encrypted, err := encryptBytes(b)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}

	if _, err := file.Write(encrypted); err != nil {
		return err
	}
	file.Close()

	return nil
}
