package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening on %s", host)
	for  {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) *rpc.Client {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
	return jsonrpc.NewClient(conn)
}
