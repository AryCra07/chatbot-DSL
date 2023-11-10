package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSha256(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(hashBytes)
	return hashCode
}
