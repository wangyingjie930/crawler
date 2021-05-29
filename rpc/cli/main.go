package main

import (
	"fmt"
	rpcdemo "learn-golang/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	// 客户端
	// 连接tcp
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoStruct.Div",
		rpcdemo.Args{A: 10, B: 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoStruct.Div",
		rpcdemo.Args{A: 10, B: 0}, &result)
	fmt.Println(111)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
