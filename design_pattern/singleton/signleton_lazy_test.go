package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	if GetInstance() != GetInstance() {
		t.Errorf("instance get failed")
	}
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("parallel get instance failed")
			}
		}
	})
}

func TestGetLazyInstance(t *testing.T) {
	if GetLazyInstance() != GetLazyInstance() {
		t.Errorf("lazy instance get failed")
	}
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("parallel get lazy failed")
			}
		}
	})
}
