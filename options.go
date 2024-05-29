package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

var encryptions = []string{
	"SHA-256/AES-256-GCM",
}

var encryptFunction = []func([]byte) ([]byte, error){
	encryptsha256aes256gcm,
}

func encryptsha256aes256gcm(b []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(key)
	hkey := h.Sum(nil)

	c, err := aes.NewCipher(hkey)
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

var decryptFunction = []func([]byte) ([]byte, error){
	decryptsha256aes256gcm,
}

func decryptsha256aes256gcm(b []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(key)
	hkey := h.Sum(nil)

	c, err := aes.NewCipher(hkey)
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
