package main

import (
	"fmt"
	"goRpcClientDemo/rpcDemo"
)

func main() {
	fmt.Println("goRpcClientDemo server starting ... ")
	rpcDemo.DemoClient()
}
