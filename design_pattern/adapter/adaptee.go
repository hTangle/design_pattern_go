package adapter

import "fmt"

type Adaptee struct {
}

func (a *Adaptee) Request() {
	fmt.Printf("requested by adaptee")
}
