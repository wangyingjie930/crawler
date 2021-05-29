package main

import (
	"fmt"
	rpc2 "learn-golang/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 简易的 rpc 服务器
	rpc.Register(rpc2.DemoStruct{})
	// 监听1234端口
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		// 接受一个传入的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go func() {
			jsonrpc.ServeConn(conn)
			fmt.Printf("111111")
		}()
	}
}
