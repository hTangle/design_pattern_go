package birdge

const (
	ColorRed  = "red"
	ColorBlue = "blue"
)

type Color interface {
	GetColor() string
}

type Red struct {
}

func (r *Red) GetColor() string {
	return ColorRed
}

type Blue struct {
}

func (b *Blue) GetColor() string {
	return ColorBlue
}
