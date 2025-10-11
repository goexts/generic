package slices_test

import (
	"math/rand"
	"testing"

	"github.com/goexts/generic/slices"
)

func generateTestData(n, m int) ([]int, []int) {
	source := make([]int, n)
	for i := range source {
		source[i] = i
	}

	target := make([]int, m)
	for i := range target {
		target[i] = rand.Intn(n * 2) // 50% 的概率在 source 中
	}

	return source, target
}

func BenchmarkFilterIncluded(b *testing.B) {
	sizes := []struct {
		n int // source 大小
		m int // target 大小
	}{
		{10, 5},
		{100, 10},
		{1000, 50},
		{10000, 100},
	}

	for _, size := range sizes {
		source, target := generateTestData(size.n, size.m)
		b.Run(
			b2s(size.n, size.m),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = slices.FilterIncluded(source, target)
				}
			},
		)
	}
}

func b2s(n, m int) string {
	return "n=" + itoa(n) + ",m=" + itoa(m)
}

// 简单的 int 转 string 函数，避免引入 strconv
func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var b [32]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}
