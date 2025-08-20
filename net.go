package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
	"znp-cs/messages"
	"znp-cs/status"
)

type (
	Port string
	IP4  string

	// The State Processor is a function that determines how a node reacts to a message.
	// Important states:
	//  LISTENING		- waiting for message from other before it does anything
	//  TALKING 		- wants to send a message to the other node
	//  STOPPING 		- closing the connection
	// How a state processor works is incredible implementation dependant.
	stateProcessor func(Connection, status.State, messages.Message) (status.State, messages.Message)

	Connection interface {
		Connect()                 // starts a net.Conn with the given address (Dial IP)
		Send(messages.Message)    // Sends a Message through the connection
		Listen() messages.Message // Reads a Message from the connection
		Close()                   // Closes the connection
	}

	Client interface {
		read()
		write()
	}
)

func processState(m Connection, state status.State, msg messages.Message) (status.State, messages.Message) {
	switch state {

	// Send
	case status.TALKING:
		in := m.Read()
		m.Send(in)

		state = status.LISTENING

	// Listen
	case status.LISTENING:
		message := m.Listen()

		state = process(message)

	// Stop
	case status.STOPPING:
		break
	default:
		fmt.Println("Erroneous state: ", state)
		state = status.STOPPING
	}
	return status.STOPPING, messages.EmptyMessage()
}

func NewMessageConnection(address, port string, stateProcessor stateProcessor) MessageConnection {
	m := MessageConnection{}

	m.process = processState
	m.address = IP4("127.0.0.1")
	m.port = Port("25565")

	return m
}

type MessageConnection struct {
	port    Port
	address IP4
	conn    net.Conn

	// A simple processor could be a function that
	// outputs the message to m.output, and reads from m.input.
	// EG: output to os.Stdout, input from os.Stdin
	process stateProcessor
}

func (m MessageConnection) Connect() {
	address := string(m.address) + ":" + string(m.port)

	fmt.Println("Client connecting to port", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	fmt.Println("Connected to server on ", conn.RemoteAddr())
	m.connect(conn)
}

func (m MessageConnection) connect(conn net.Conn) {
	m.conn = conn
	fmt.Println("Ready to receive message from server ")
	var state status.State = status.LISTENING
	reply := messages.EmptyMessage()

	for {
		time.Sleep(1 * time.Second)

		state, reply = m.process(m, state, reply)

		if state == status.STOPPING {
			m.Close()
			break
		}

		switch state {
		default:
			fmt.Println("Erroneous state: ", state)
			state = "stopping"
		}
	}
}

// Send a message through the connection.
func (m MessageConnection) Send(msg messages.Message) {
	err := messages.Send(m.conn, msg)

	isEOF := errors.As(err, io.EOF)
	if isEOF {
		m.Close()
	}
}

func (m MessageConnection) Listen() messages.Message {
	n, msg := receiveMessage(m.conn)

	return msg[:n]
}

func (m MessageConnection) Close() {
	m.conn.Close()
}

// Read Write

// Read 1 line from the io.Reader
func (m MessageConnection) Read() messages.Message {
	scanner := bufio.NewScanner(m.input)
	scanner.Scan()

	if scanner.Err() != nil {
		// Handle error.
	}

	return messages.NewMessage(scanner.Text())
}
