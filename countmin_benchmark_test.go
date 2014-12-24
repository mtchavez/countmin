package countmin

import (
	"fmt"
	"testing"
)

func Benchmark_New_1000(b *testing.B)  { benchNew(b, 1000) }
func Benchmark_New_10000(b *testing.B) { benchNew(b, 10000) }

func benchNew(b *testing.B, total int) {
	for i := 0; i < b.N; i++ {
		New(total, total)
	}
}

func Benchmark_Add_1000(b *testing.B)   { benchAdd(b, 1000) }
func Benchmark_Add_10000(b *testing.B)  { benchAdd(b, 10000) }
func Benchmark_Add_100000(b *testing.B) { benchAdd(b, 1000000) }
func benchAdd(b *testing.B, total int) {
	cm := New(40, 200)
	b.ResetTimer()
	var i int64
	for i = 0; i < int64(b.N); i++ {
		cm.Add([]byte(fmt.Sprintf("http://domain%d.com/page%d", i, i)), i)
	}
}
