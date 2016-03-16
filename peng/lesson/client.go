package lesson

import (
	"fmt"
	"log"
	//"net/rpc" 方法2
	"net/rpc/jsonrpc" //方法3
	"os"
)

type Args struct {
	A, B int
}
type Quotiant struct {
	Quo, Rem int
}

func InitClient() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddr := os.Args[1]

	//方法1 client, err := rpc.DialHTTP("tcp", serverAddr+":1234")
	//方法2 client, err := rpc.Dial("tcp", serverAddr+":1234")
	client, err := jsonrpc.Dial("tcp", serverAddr+":1234")

	if err != nil {
		log.Fatal("dialhttp", err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Math.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Match error", err)
	}

	fmt.Printf("math: %d*%d = %d\n", args.A, args.B, reply)

	var quo Quotiant
	err = client.Call("Math.Divide", args, &quo)
	if err != nil {
		log.Fatal("Match error", err)
	}
	fmt.Printf("math: %d / %d = %d reminder %d\n", args.A, args.B, quo.Quo, quo.Rem)
}
