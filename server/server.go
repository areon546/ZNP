package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	s := zNPServer{port: 8691}
	s.Run()
}

func NewZNPServer(port int) zNPServer {
	return zNPServer{port: port}
}

type zNPServer struct {
	port int
}

func (s zNPServer) Run() {
	address := ":8691"
	fmt.Println(strconv.Itoa(s.port))
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

	time.Sleep(500000)
	// wait for message
	sendMessage(conn, "Hi Friend")

	conn.Close()
}

func sendMessage(conn net.Conn, msg string) {
	n, err := fmt.Fprintf(conn, msg)
	fmt.Println("N", n, "bytes written")
	if err != nil {
		fmt.Println("Sending Message: Error: ", err)
	}
}

func receiveMessage(conn net.Conn) (msg string) {
	return
}
