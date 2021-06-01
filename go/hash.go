package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

const _hashlength = 32

// Keccak256 is the Keccak-256 hashing method
type HmacSha256 struct {
	secret string
}

// New creates a new Keccak-256 hashing method
func New(secret string) *HmacSha256 {
	return &HmacSha256{
		secret: secret,
	}
}

// HashLength returns the length of hashes generated by Hash() in bytes
func (h *HmacSha256) HashLength() int {
	return _hashlength
}

// Hash generates a Keccak-256 hash from a byte array
func (h *HmacSha256) Hash(d ...[]byte) []byte {
	hash := hmac.New(sha256.New, []byte(h.secret))
	for _, value := range d {
		hash.Write(value)
	}

	return hash.Sum(nil)
}