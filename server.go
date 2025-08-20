package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"znp-cs/status"
)

var ip string = ""

func NewZNPServer(port Port) znpServer {
	return znpServer{port: port}
}

type znpServer struct {
	port Port
}

func (s znpServer) Run() {
	address := ":8691"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on ", address)

	// Loop to be able to handle multiple connections.
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
	state := "talking"

	time.Sleep(500000)
	// wait for message

	for {
		time.Sleep(1 * time.Second)
		switch state {
		case status.TALKING:
			print(conn, state)
			sendMessage(conn, "Hi Friend")

			state = status.LISTENING
		case status.LISTENING:
			print(conn, state)

			state = status.STOPPING
		case status.STOPPING:
			print(conn, state)

			conn.Close()
			return
		default:
			fmt.Println("Erroneous state: ", state)
			state = status.STOPPING
		}
	}
}

func printC(conn net.Conn, message string) {
	fmt.Println(conn.RemoteAddr(), message)
}

func sendMessage(conn net.Conn, msg string) {
	print(conn, "Sending: "+msg)
	n, err := fmt.Fprintf(conn, msg)
	fmt.Println("N", n, "bytes written")
	if err != nil {
		fmt.Println("Sending Message: Error: ", err)
	}
}
