package birdge

const (
	ShapeCircle = "circle"
	ShapeSquare = "square"
)

type Shape interface {
	GetShape() string
	GetInfo() (string, string)
}

type ShapeBase struct {
	Color Color
}

type Circle struct {
	ShapeBase
}

func (c *Circle) GetShape() string {
	return ShapeCircle
}
func (c *Circle) GetInfo() (string, string) {
	return c.GetShape(), c.Color.GetColor()
}

type Square struct {
	ShapeBase
}

func (s *Square) GetShape() string {
	return ShapeSquare
}

func (s *Square) GetInfo() (string, string) {
	return s.GetShape(), s.Color.GetColor()
}

func NewCircleShapeWithColor(color Color) Shape {
	return &Circle{ShapeBase{Color: color}}
}

func NewSquareShapeWithColor(color Color) Shape {
	return &Square{ShapeBase{Color: color}}
}
