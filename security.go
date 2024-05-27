package main

import (
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
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
