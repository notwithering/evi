package main

import (
	"crypto/aes"
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

func hash(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}
