package internal

import (
	"crypto/md5"
	"encoding/hex"
)

// Hash.
type Hash struct {}

// GetMD5 returns md5 hash as text.
func (h *Hash) GetMD5(input []byte) string {
	hash := md5.Sum(input)
	return hex.EncodeToString(hash[:])
}