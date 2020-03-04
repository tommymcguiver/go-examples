package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//Sha256 ...
func Sha256(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}

//HmacSha256 ...
func HmacSha256(data string, key string) (string, error) {
	mac := hmac.New(sha256.New, []byte(key))
	_, err := mac.Write([]byte(data))

	if err != nil {
		return "", err
	}

	// to lowercase hexits
	return hex.EncodeToString(mac.Sum(nil)), err

}
