package shared

import (
	"errors"
	"fmt"
	"io"
	"net"
)

const bufferSize int = 1024

// Creates a byte array for use by Writers and Readers
func msgBuffer(size int) []byte {
	return make([]byte, size)
}

func SendMessage(conn net.Conn, msg []byte) {
	// Split message up into chunks if larger than buffer size 1024

	if len(msg) > bufferSize {
		panic("TOO LONG")
		// Split up message into chunks
	}

	buffer := msg
	sendMessage(conn, buffer)
}

func sendMessage(conn net.Conn, msg []byte) {
	send := msgBuffer(bufferSize)
	n, err := conn.Write(send)
	if err != nil {
		fmt.Println("Err sending message:", err)
	}

	fmt.Println("Bytes Written: ", n)
}

// Blocks thread whilst waiting for message to be returned.
func ReceiveMessage(conn net.Conn) (n int, msg []byte) {
	msg = msgBuffer(bufferSize)
	n, err := conn.Read(msg)

	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println("Err receiving message:", err)
	}
	fmt.Println("Number bytes read", n, "Error? ", err)

	return
}
