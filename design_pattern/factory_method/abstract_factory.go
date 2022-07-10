package factory_method

type Container interface {
	Put() string
}

type GlassBottle struct {
}

func (g *GlassBottle) Put() string {
	return "put into glass bottle"
}

type Cans struct {
}

func (c *Cans) Put() string {
	return "put into cans"
}

type CanningFactory interface {
	CreateContainer() Container
	CreateFruit() Fruit
}

type AppGlassFactory struct {
}

func (a *AppGlassFactory) CreateContainer() Container {
	return &GlassBottle{}
}
func (a *AppGlassFactory) CreateFruit() Fruit {
	return &Apple{
		name:  "apple",
		type_: "apple",
	}
}

type BananaCansFactory struct {
}

func (b *BananaCansFactory) CreateContainer() Container {
	return &Cans{}
}

func (b *BananaCansFactory) CreateFruit() Fruit {
	return &Banana{
		name:  "banana",
		type_: "banana",
	}
}
