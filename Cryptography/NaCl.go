package main

import (
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
)

const (
	KeySize   = 32
	NonceSize = 24
)

var (
	ErrEncrypt = errors.New("secret: encryption failed")
	ErrDecrypt = errors.New("secret: encryption failed")
)

// Create a new random secret key
func GenerateKey() (*[KeySize]byte, error) {
	key := new([KeySize]byte)
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, err
	}

	return key, nil
}

// Create a new random secret key
func GenerateNonce() (*[NonceSize]byte, error) {
	nonce := new([NonceSize]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		return nil, err
	}

	return nonce, nil
}

// genereate a random nonce and ecnrypt the input using
// NaCl's secretbox package. The nonce is prepended to the
// ciphertext. A sealed message will be the same size as the
// // original message plus secretbox.Overhead bytes long.
func Encrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
	nonce, err := GenerateNonce()
	if err != nil {
		return nil, ErrEncrypt
	}

	out := make([]byte, len(nonce))
	copy(out, nonce[:])
	out = secretbox.Seal(out, message, nonce, key)
	return out, nil
}

// Extract the nonce from the ciphertext, and attempt to
// decrypt with NaCl's secretbox
func Decrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
	if len(message) < (NonceSize + secretbox.Overhead) {
		return nil, ErrDecrypt
	}

	var nonce [NonceSize]byte
	copy(nonce[:], message[:NonceSize])
	out, ok := secretbox.Open(nil, message[NonceSize:], &nonce, key)
	if !ok {
		return nil, ErrDecrypt
	}

	return out, nil
}
