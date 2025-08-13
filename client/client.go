package main

import (
	"fmt"
	"net"
)

func main() {
	c := ZNPClient{}
	c.Connect("localhost")
}

type ZNPClient struct{}

func (c ZNPClient) Connect(address string) {
	conn, err := net.Dial("tcp", address+":8691")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server on localhost:8080")
}
