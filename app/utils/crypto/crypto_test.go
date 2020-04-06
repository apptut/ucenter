package crypto_test

import (
	"fmt"
	"testing"
	"ucenter/app/utils/crypto"
)

func TestMD5(t *testing.T) {
	fmt.Println(crypto.MD5("hello"))
}

func TestSHA256(t *testing.T) {
	fmt.Println(crypto.SHA256("hello"))
}
