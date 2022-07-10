package adapter

type Adapter struct {
	Adaptee Adaptee
}

func (a *Adapter) Request() {
	a.Adaptee.Request()
}
