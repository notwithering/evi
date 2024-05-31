package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
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

func encryptBytes(b []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce[:], nonce[:], b, nil), nil
}

func decryptBytes(b []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	if len(b) < nonceSize {
		return nil, fmt.Errorf("cipher text is smaller than the nonce size")
	}

	nonce, cipherBytes := b[:nonceSize], b[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherBytes, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func burnString(s *string) {
	byteSlice := []byte(*s)

	for i := range byteSlice {
		byteSlice[i] = 0
	}

	*s = string(byteSlice)
}

func burnBytes(b []byte) {
	for i := range b {
		b[i] = 0
	}
}
