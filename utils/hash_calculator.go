package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(value string) string {
	sha := sha256.New()
	sha.Write([]byte(value))
	hash := sha.Sum(nil)
	return hex.EncodeToString(hash)
}
