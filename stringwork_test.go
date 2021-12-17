package simpleTesting

import "testing"

var num = 100000

func BenchmarkStringFromAssignment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringFromAssignment(num)
	}
}

func BenchmarkStringFromAppendJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringFromAppendJoin(num)
	}
}

func BenchmarkStringsFromBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsFromBuffer(num)
	}
}
