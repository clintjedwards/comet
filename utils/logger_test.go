package utils

import "testing"

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Log().Info("")
	}
}
