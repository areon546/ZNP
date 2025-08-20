package messages

import (
	"bufio"
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

func (m Message) String() string {
	return string(m)
}

func Send(conn net.Conn, msg Message) error {
	// TODO: Split message up into chunks if larger than buffer size 1024

	_, err := conn.Write(msg)
	if err != nil {

		fmt.Println("Err sending message:", err)
		return err
	}

	return nil
}

// Blocks thread whilst waiting for message to be returned.
func Receive(conn net.Conn) (n int, msg Message, err error) {
	msg = msgBuffer(bufferSize)
	n, err = conn.Read(msg)

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

// IN and OUT

// Read 1 line from the io.Reader
func Read(in io.Reader) Message {
	scanner := bufio.NewScanner(in)
	scanner.Scan()

	if scanner.Err() != nil {
		// Handle error.
	}

	return NewMessage(scanner.Text())
}
