package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func MD5(text string) string {
	hashed := md5.New()
	hashed.Write([]byte(text))
	return hex.EncodeToString(hashed.Sum(nil))
}

func SHA256(text string) string {
	hashed := sha256.New()
	hashed.Write([]byte(text))
	return hex.EncodeToString(hashed.Sum(nil))
}
