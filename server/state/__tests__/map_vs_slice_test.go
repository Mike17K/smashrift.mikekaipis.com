package main

import (
	"sync"
	"testing"
)

func BenchmarkSyncMapRange(b *testing.B) {
	var sm sync.Map
	for i := 0; i < 10000; i++ {
		sm.Store(i, i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sm.Range(func(key, value any) bool {
			_ = key.(int)
			_ = value.(int)
			return true
		})
	}
}

func BenchmarkSliceLoop(b *testing.B) {
	slice := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		slice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range slice {
			_ = v
		}
	}
}
