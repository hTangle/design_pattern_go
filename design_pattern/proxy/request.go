package proxy

import (
	"fmt"
	"time"
)

type Request struct {
}

func (r *Request) Post(body interface{}) error {
	fmt.Printf("%v", body)
	return nil
}

type RequestProxy struct {
	Request Request
}

func (r *RequestProxy) Post(body interface{}) error {
	start := time.Now()
	err := r.Request.Post(body)
	end := time.Now().Sub(start)
	fmt.Printf("cost: %v", end)
	return err
}
