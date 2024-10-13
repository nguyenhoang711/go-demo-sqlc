package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()
    hash.Write([]byte(key))
    return hex.EncodeToString(hash.Sum(nil))
}