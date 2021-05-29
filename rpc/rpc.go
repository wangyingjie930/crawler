package rpc

import "errors"

type DemoStruct struct {

}

type Args struct {
	A int
	B int
}

func (d DemoStruct) Div (arg Args, result *interface{}) error {
	if arg.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(arg.A) / float64(arg.B)
	return nil
}
