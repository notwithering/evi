package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

var encryptions = []string{
	"SHA-256/AES-256-GCM",
	"SHA-512/AES-128-CBC",
}

var encryptFunction = []func([]byte) ([]byte, error){
	encryptsha256aes256gcm,
	encryptsha512aes128cbc,
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

func encryptsha512aes128cbc(b []byte) ([]byte, error) {
	hkey := sha512.Sum512(key)

	c, err := aes.NewCipher(hkey[:16])
	if err != nil {
		return nil, err
	}

	padding := c.BlockSize() - len(b)%c.BlockSize()
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	b = append(b, padtext...)

	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(c, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], b)

	return ciphertext, nil
}

var decryptFunction = []func([]byte) ([]byte, error){
	decryptsha256aes256gcm,
	decryptsha512aes128cbc,
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

func decryptsha512aes128cbc(ciphertext []byte) ([]byte, error) {
	hkey := sha512.Sum512(key)

	c, err := aes.NewCipher(hkey[:16]) // Use the first 16 bytes of the SHA-512 hash
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(c, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	length := len(ciphertext)
	if length == 0 {
		return nil, fmt.Errorf("input is empty")
	}

	padlen := int(ciphertext[length-1])
	if padlen > length {
		return nil, fmt.Errorf("padding size is larger than the input")
	}

	for _, v := range ciphertext[length-padlen:] {
		if int(v) != padlen {
			return nil, fmt.Errorf("invalid padding")
		}
	}

	return ciphertext[:length-padlen], nil
}
