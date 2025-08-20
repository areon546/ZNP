package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"znp-cs/messages"
	"znp-cs/status"
)

const bufferSize int = 1024

func NewZNPClient(port int, in io.Reader) znpClient {
	z := znpClient{}
	c := NewMessageConnection("127.0.0.1", "8961", z.stateProcessor)

	z.conn = c

	return z
}

type znpClient struct {
	conn MessageConnection

	input io.Reader
	scan  bufio.Scanner

	output io.Writer
}

func (z znpClient) stateProcessor() stateProcessor {
	sp := func(m Connection, state status.State, msg messages.Message) (status.State, messages.Message) {
		fmt.Println(state, msg)

		state = status.STOPPING
		return state, msg
	}
	return sp
}

// Read 1 line from the io.Reader
func (z znpClient) Read() messages.Message {
	scanner := bufio.NewScanner(z.input)
	scanner.Scan()

	if scanner.Err() != nil {
		// Handle error.
	}

	return messages.NewMessage(scanner.Text())
}

// MISC

func receiveMessage(c net.Conn) (int, messages.Message) {
	return messages.Receive(c)
}

func process(msg messages.Message) status.State {
	fmt.Println(len(msg), msg)
	return status.STOPPING
}
