package crypto

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rel := ""
		rel += strconv.Itoa(2020)
		rel += "年, 祝大家"
		rel += "万事如意"
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString(strconv.Itoa(2020))
		builder.WriteString("年, 祝大家")
		builder.WriteString("万事如意")

		_ = builder.String()
	}
}

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d年, 祝大家：%s", 2020, "万事如意")
	}
}
