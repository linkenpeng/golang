package lesson

import (
	"errors"
	"fmt"
	// "net/http" 方法1需要的包
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Math int

func (m *Math) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *Math) Divide(args *Args, quo *Quotiant) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func InitServer() {
	math := new(Math)
	rpc.Register(math)

	// 方法2
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(2)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Listen error:", err)
		os.Exit(2)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn error:", err)
			continue
		}

		//rpc.ServeConn(conn) 方法2

		//方法3
		jsonrpc.ServeConn(conn)
	}
	/*
		方法1 http
		rpc.HandleHTTP()

		err := http.ListenAndServe(":1234", nil)
		if err != nil {
			fmt.Println(err.Error())
		}
	*/
}
