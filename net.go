package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
	"znp-cs/messages"
	"znp-cs/status"
)

type (
	Port string
	IP4  string

	// The State Processor is a function that determines how a node reacts to a message.
	// Important states:
	// 	INIT		- the node starts on this state, implementations should use this as a signal to control which state to swap to at the start of a conversation
	//  HEAR		- waiting for message from other before it does anything
	//  TALK 		- wants to send a message to the other node
	//  STOP 		- closing the connection
	// How a state processor works is incredibly implementation dependent.
	//
	// The base theory behind the StateProcessor function is that a MessageConnection will loop through and call it repeatedly until it receives the STOP command.
	// When it receives the STOP command, it will close the connection.
	//
	// NOTE: The implementation of the state processor must handle the relevant errors, eg EOF
	StateProcessor func(Connection, status.State, messages.Message) (status.State, messages.Message)

	Connection interface {
		Host(Port)   // starts a new.Conn listener (Listen PORT)
		Connect(IP4) // starts a net.Conn with the given address (Dial IP)

		TALK(messages.Message) error     // Sends a Message through the connection
		HEAR() (messages.Message, error) // Reads a Message from the connection
		Close()                          // Closes the connection
	}

	Client interface {
		read()
		write()
	}
)

func (p Port) String() string {
	return string(p)
}

func errIsEOF(err error) bool {
	return errors.Is(err, io.EOF)
}

// A basic State Processor that will alternate between TALK and HEAR until it HEARs STOP
func ProcessState(m Connection, state status.State, msg messages.Message) (status.State, messages.Message) {
	switch state {
	// Start of conversation
	case status.INIT:
		msg = messages.EmptyMessage()
		state = status.HEAR

	// Sending message
	case status.TALK:
		fmt.Print("TALK: ")
		msg = messages.Read(os.Stdin)
		err := m.TALK(msg)

		// TODO: Process Error
		fmt.Println(err)

		state = status.HEAR

	// Listening to Message
	case status.HEAR:

		msg, err := m.HEAR()

		if errIsEOF(err) {
			state = status.STOP
			break
		}

		state = status.TALK

		fmt.Println("Heard message: ", msg)
		fmt.Println("HEAR, Switching to TALK")

		startsWithSTOP := strings.HasPrefix(msg.String(), "STOP")
		if startsWithSTOP {
			state = status.STOP
		}

	// Stop
	case status.STOP:
		msg = messages.EmptyMessage()
	default:
		fmt.Println("HEAR == status.HEAR", state, status.HEAR)
		fmt.Println("Erroneous state: ", state)
		state = status.STOP
	}
	return state, msg
}

func NewMessageConnection(port Port, stateProcessor StateProcessor) MessageConnection {
	m := MessageConnection{}

	m.process = stateProcessor
	m.port = port

	return m
}

type MessageConnection struct {
	port Port
	conn net.Conn // NEEDS TO BE INITIALISED

	// A simple processor could be a function that
	// outputs the message to m.output, and reads from m.input.
	// EG: output to os.Stdout, input from os.Stdin
	process StateProcessor
}

// Code for acting as a Server
func (m MessageConnection) Host(port Port) {
	address := ":" + port.String()
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
		go m.converse(conn)
	}
}

// Boilerplate for making a connection as a Client.
func (m MessageConnection) Connect(address IP4) {
	port := string(address) + ":" + string(m.port)

	fmt.Println("Node connecting to port", port)
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	fmt.Println("Connected to ", conn.RemoteAddr())
	// conn.Close()
	m.converse(conn)
}

// the node talking across the
func (m MessageConnection) converse(conn net.Conn) {
	m.conn = conn
	fmt.Println("Ready to receive message from ", conn.RemoteAddr())
	state := status.INIT
	reply := messages.EmptyMessage()

	for {
		time.Sleep(1 * time.Second)

		state, reply = m.process(m, state, reply)

		if state == status.STOP {
			m.Close()
			break
		}
	}
}

// Send a message through the connection.
func (m MessageConnection) TALK(msg messages.Message) error {
	err := messages.Send(m.conn, msg)

	fmt.Println("TALK:", msg)
	return err
}

func (m MessageConnection) HEAR() (messages.Message, error) {
	n, msg, err := messages.Receive(m.conn)
	heard := msg[:n]

	fmt.Println("HEAR:", heard)
	return heard, err
}

func (m MessageConnection) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

// Read Write
