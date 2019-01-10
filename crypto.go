package aliyuniot

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

const (
	DefaultSignAlgorithm = HMACSHA1
	HMACSHA1             = "sha1"
	HMACSHA256           = "sha256"
)

func hmacSign(algo string, secret, data string) string {
	hfun := buildHashFunc(algo)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(hfun, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}

type hashFunc func() hash.Hash

func buildHashFunc(algo string) hashFunc {
	switch algo {
	case HMACSHA256:
		return sha256.New
	case HMACSHA1:
		return sha1.New
	}
	panic("unsupport hash algorithm:" + string(algo))
}
