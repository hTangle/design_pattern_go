package factory_method

import (
	"testing"
)

func TestNewFruitFactory(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "apple",
			args:  struct{ name string }{name: "apple"},
			want:  "apple",
			want1: "apple",
		},
		{
			name:  "orange",
			args:  struct{ name string }{name: "orange"},
			want:  "orange",
			want1: "orange",
		}, {
			name:  "banana",
			args:  struct{ name string }{name: "banana"},
			want:  "banana",
			want1: "banana",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NewFruitFactory(tt.args.name)
			if got.CreateFruit().Type() != tt.want1 {
				t.Errorf("NewFruitFactory() got = %v, want1 %v", got.CreateFruit().Type(), tt.want1)
			}
			if got.CreateFruit().Name() != tt.want {
				t.Errorf("NewFruitFactory() got = %v, want %v", got.CreateFruit().Type(), tt.want)
			}
		})
	}
}
