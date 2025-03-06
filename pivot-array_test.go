package main

import (
	"testing"
)

func BenchmarkPivotArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pivotArray([]int{-3, 4, 3, 2}, 2)
	}
}
