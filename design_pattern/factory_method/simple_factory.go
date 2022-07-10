package factory_method

type Fruit interface {
	Name() string
	Type() string
}
type Apple struct {
	name  string
	type_ string
}

func (a *Apple) Name() string {
	return a.name
}

func (a *Apple) Type() string {
	return a.type_
}

type Banana struct {
	name  string
	type_ string
}

func (a *Banana) Name() string {
	return a.name
}

func (a *Banana) Type() string {
	return a.type_
}

type Orange struct {
	name  string
	type_ string
}

func (a *Orange) Name() string {
	return a.name
}
func (a *Orange) Type() string {
	return a.type_
}

func NewSimpleFruitFactory(name string) Fruit {
	switch name {
	case "apple":
		return &Apple{name: name}
	case "banana":
		return &Banana{name: name}
	case "orange":
		return &Orange{name: name}
	}
	return nil
}
