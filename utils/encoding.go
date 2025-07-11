package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func EncodeToShortPath(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	// Ambil 6 karakter base64 agar pendek
	return base64.URLEncoding.EncodeToString(hash[:])[:6]
}
