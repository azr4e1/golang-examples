package main

import (
	"crypto/aes"
	"crypto/cipher"
)

// Encrypt secures a message using AES-GCM
func EncryptAES(key, message []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrEncrypt
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, ErrEncrypt
	}

	nonce, err := GenerateNonce()
	if err != nil {
		return nil, ErrEncrypt
	}

	out := gcm.Seal(nonce[:], nonce[:], message, nil)
	return out, nil
}
