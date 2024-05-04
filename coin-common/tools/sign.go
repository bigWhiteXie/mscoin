package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateHmacSha256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
