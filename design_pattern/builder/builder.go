package builder

import "fmt"

const (
	defaultA = 1
	defaultB = 2
	defaultC = 3
)

type KV struct {
	name string
	a    int
	b    int
	c    int
}

type Option struct {
	a int
	b int
	c int
}

type ConfigOptionFunc func(option *Option)

func NewKVConfig(name string, opts ...ConfigOptionFunc) (kv *KV, err error) {
	if name == "" {
		return nil, fmt.Errorf("name should not be empty")
	}
	option := &Option{
		a: defaultA,
		b: defaultB,
		c: defaultC,
	}
	for _, opt := range opts {
		opt(option)
	}
	return &KV{
		name: name,
		a:    option.a,
		b:    option.b,
		c:    option.c,
	}, nil
}
