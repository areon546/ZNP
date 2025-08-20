package main

import (
	"bufio"
	"fmt"
	"io"
	"znp-cs/messages"
	"znp-cs/status"
)

const bufferSize int = 1024

func NewZNPClient(port Port, in io.Reader) znpClient {
	z := znpClient{}
	c := NewMessageConnection(port, z.stateProcessor())

	z.conn = c

	return z
}

type znpClient struct {
	conn MessageConnection

	input io.Reader
	scan  bufio.Scanner

	output io.Writer
}

func (z znpClient) stateProcessor() StateProcessor {
	sp := func(m Connection, state status.State, msg messages.Message) (status.State, messages.Message) {
		fmt.Println(state, msg)

		state = status.STOP
		return state, msg
	}
	return sp
}

func (z znpClient) Connect(ip IP4) {
	z.conn.Connect(ip)
}
