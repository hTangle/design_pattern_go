package factory_method

import (
	"testing"
)

func TestNewSimpleFruitFactory(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "apple",
			args: args{
				name: "apple",
			},
			want: "apple",
		},
		{
			name: "banana",
			args: args{
				name: "banana",
			},
			want: "banana",
		},
		{
			name: "orange",
			args: args{
				name: "orange",
			},
			want: "orange",
		},
	}
	for _, tes := range tests {
		t.Run(tes.name, func(ts *testing.T) {
			if got := NewSimpleFruitFactory(tes.args.name); !(got.Name() == tes.want) {
				ts.Errorf("NewSimpleFruitFactory Name=%s, want %s", got.Name(), tes.want)
			}
		})
	}
}
