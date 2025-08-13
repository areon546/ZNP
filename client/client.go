package client

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"znp-cs/shared"
)

const bufferSize int = 1024

func NewZNPClient(port int) zNPClient {
	return zNPClient{port: port}
}

type zNPClient struct {
	port int
}

func (c zNPClient) Connect(address string) {
	address += ":" + strconv.Itoa(c.port)
	fmt.Println("Client connecting to port", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	fmt.Println("Connected to server on ", address)
	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	time.Sleep(1000000000)
	fmt.Println("Ready to receive message from server ")

	n, msg := receiveMessage(conn)
	fmt.Println(msg[:n])
	conn.Close()
}

func receiveMessage(c net.Conn) (int, []byte) {
	return shared.ReceiveMessage(c)
}

func processMessage(msg []byte) {
	fmt.Println(len(msg))
}
