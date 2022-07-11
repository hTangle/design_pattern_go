package facade

import "fmt"

const (
	ColorRed    = "red"
	ColorBlue   = "blue"
	ShapeCircle = "circle"
	ShapeSquare = "square"
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

type Shape interface {
	GetShape() string
}

type Circle struct {
}

func (c *Circle) GetShape() string {
	return ShapeCircle
}

type Square struct {
}

func (s *Square) GetShape() string {
	return ShapeSquare
}

type Drawer struct {
}

func (d *Drawer) Draw() {
	circle := &Circle{}
	red := &Red{}
	blue := &Blue{}
	square := &Square{}
	fmt.Printf("get one %s with %s and one %s with %s", circle.GetShape(), red.GetColor(), square.GetShape(), blue.GetColor())
}
