package main

import (
	"fmt"
	"io"
	"os"
)

func decrypt() {
	if fileExists(filename) {
		if err := decryptFile(filename); err != nil {
			fmt.Printf(eviError, err)
			os.Exit(1)
		}
	}
}

func encrypt() {
	if fileExists(filename) {
		if err := encryptFile(filename); err != nil {
			fmt.Printf(eviError, err)
			fmt.Print("\n")
			removeFile()
		}
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

	plainText, err := decryptFunction[encryption](b)
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

	encrypted, err := encryptFunction[encryption](b)
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
