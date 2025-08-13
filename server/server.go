package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	s := ZNPServer{port: 8691}
	s.Run()
}

type ZNPServer struct {
	port int
}

func (s ZNPServer) Run() {
	listener, err := net.Listen("tcp", ":8691")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on :8691...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Client connected:", conn.RemoteAddr())
	conn.Close()
}
