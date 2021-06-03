package rpcDemo

import (
	"fmt"
	"log"
	"net/rpc"
)

func DemoClient() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("PersonService.Hello", "hanhe", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
