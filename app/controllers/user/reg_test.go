package user

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
	"ucenter/app/utils/crypto"
)

func BenchmarkMd5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = crypto.MD5("12345678")
	}
}

func BenchmarkBcrypt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = bcrypt.GenerateFromPassword([]byte("12345678"), 18)
	}
}
