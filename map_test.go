package golang_map_bench

import "testing"

func BenchmarkNativeMap(b *testing.B) {
	m := NewNativeMap()
	StartProducer(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get()
	}
}

func BenchmarkAtomicMap(b *testing.B) {
	m := NewAtomicMap()
	StartProducer(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get()
	}
}

func BenchmarkRWLockMap(b *testing.B) {
	m := NewRWLockMap()
	StartProducer(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	m := NewSyncMap()
	StartProducer(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get()
	}
}
