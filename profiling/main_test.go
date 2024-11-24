package main

import (
	"testing"
)

// // Benchmark for the unoptimized code
// func BenchmarkUnoptimizedWork(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		unoptimizedWork()
// 	}
// }

// Benchmark for the optimized code
func BenchmarkOptimizedWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optimizedWork()
	}
}
