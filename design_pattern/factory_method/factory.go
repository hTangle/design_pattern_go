package factory_method

import "errors"

type FruitFactory interface {
	CreateFruit() Fruit
}

type AppleFruitFactory struct {
	name string
}

func (f *AppleFruitFactory) CreateFruit() Fruit {
	return &Apple{name: f.name, type_: "apple"}
}

type BananaFruitFactory struct {
	name string
}

func (f *BananaFruitFactory) CreateFruit() Fruit {
	return &Banana{name: f.name, type_: "banana"}
}

type OrangeFruitFactory struct {
	name string
}

func (f *OrangeFruitFactory) CreateFruit() Fruit {
	return &Orange{name: f.name, type_: "orange"}
}

func NewFruitFactory(name string) (FruitFactory, error) {
	switch name {
	case "apple":
		return &AppleFruitFactory{name: name}, nil
	case "banana":
		return &BananaFruitFactory{name: name}, nil
	case "orange":
		return &OrangeFruitFactory{name: name}, nil
	}
	return nil, errors.New("not supported")
}
