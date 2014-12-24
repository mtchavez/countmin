package countmin

import "testing"

func Benchmark_New_1000(b *testing.B)  { benchNew(b, 1000) }
func Benchmark_New_10000(b *testing.B) { benchNew(b, 10000) }

func benchNew(b *testing.B, total int) {
	for i := 0; i < b.N; i++ {
		New(total, total)
	}
}
