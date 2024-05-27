package main

import (
	"crypto/aes"
	"crypto/cipher"
)

var algorithms = []string{
	"AES-256",
}

var modes = []string{
	"GCM",
}

func getBlock() (cipher.Block, error) {
	switch algorithms[algorithm] {
	default: // AES-256
		return aes.NewCipher(hash256(key))
	}
}
