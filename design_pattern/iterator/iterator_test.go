package iterator

import (
	"testing"
)

func TestArray_Iterator(t *testing.T) {
	tests := []struct {
		name string
		a    Array[int]
	}{
		// TODO: Add test cases.
		{
			name: "test",
			a:    []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Iterator()
			index := 0
			for got.HasNext() {
				if got.Current() != tt.a[index] {
					t.Errorf("error: %v!=%v", got.Current(), tt.a[index])
				}
				index++
				got.Next()
			}
		})
	}
}
