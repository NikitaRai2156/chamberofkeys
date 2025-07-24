package chamberofkeys

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkInsertString(b *testing.B) {
	ch, _ := NewChamber()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		val := "value"
		_ = ch.InsertString(key, val, time.Minute)
	}
}

func BenchmarkGetString(b *testing.B) {
	ch, _ := NewChamber()
	key := "static-key"
	ch.InsertString(key, "value", time.Minute)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ch.GetString(key)
	}
}

func BenchmarkPushFront(b *testing.B) {
	ch, _ := NewChamber()
	key := "bench-list"
	ttl := time.Minute

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("val-%d", i)
		_ = ch.PushFront(key, val, ttl)
	}
}

func BenchmarkPopBack(b *testing.B) {
	ch, _ := NewChamber()
	key := "bench-list"

	// Setup: fill the list with b.N items
	for i := 0; i < b.N; i++ {
		ch.PushBack(key, fmt.Sprintf("val-%d", i), time.Minute)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ch.PopBack(key)
	}
}
