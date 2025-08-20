package messages

import (
	"errors"
	"fmt"
	"io"
	"net"
)

const bufferSize int = 1024

type Message []byte

func NewMessage(msg string) Message {
	return []byte(msg)
}

func EmptyMessage() Message {
	return NewMessage("")
}

func Send(conn net.Conn, msg Message) error {
	// Split message up into chunks if larger than buffer size 1024

	if len(msg) > bufferSize {
		panic("TOO LONG")
		// Split up message into chunks
	}

	buffer := msg // TODO: Enforce buffer size limit is 1024
	return sendMessage(conn, buffer)
}

func sendMessage(conn net.Conn, msg Message) error {
	n, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Err sending message:", err)
		return err
	}

	fmt.Println("Bytes Written: ", n)
	return nil
}

// Blocks thread whilst waiting for message to be returned.
func Receive(conn net.Conn) (n int, msg Message) {
	msg = msgBuffer(bufferSize)
	n, err := conn.Read(msg)

	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println("Err receiving message:", err)
	}
	fmt.Println("Number bytes read", n, "Error? ", err)

	return
}

// Creates a byte array for use by Writers and Readers
func msgBuffer(size int) []byte {
	return make([]byte, size)
}
