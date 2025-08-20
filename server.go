package main

import (
	"fmt"
	"znp-cs/messages"
	"znp-cs/status"
)

var ip string = ""

func NewZNPServer(port Port) znpServer {
	z := znpServer{}
	m := NewMessageConnection(port, z.processState())
	z.conn = m
	return z
}

type znpServer struct {
	conn MessageConnection

	port Port
}

func (z znpServer) processState() StateProcessor {
	return func(c Connection, s status.State, msg messages.Message) (status.State, messages.Message) {
		switch s {
		case status.INIT:
			s = status.TALK
		case status.TALK:
			err := c.TALK(messages.NewMessage("Hi Friend"))
			fmt.Println("ERRRR: ", err)

			s = status.HEAR
		case status.HEAR:
			heard, err := c.HEAR() // Using an intermediary variable since msg would get re-declared if here
			fmt.Println("ERRRR: ", err)

			if errIsEOF(err) {
				s = status.STOP
				break
			}
			s = status.TALK
			msg = heard
		case status.STOP:
			msg = messages.EmptyMessage()
		default:
			fmt.Print("Server encountered error: ", status.Lookup(s), msg)
		}

		return s, msg
	}
}

// func (s znpServer) Run() {
// 	address := ":8691"
// 	listener, err := net.Listen("tcp", address)
// 	if err != nil {
// 		fmt.Println("Error creating listener:", err)
// 		os.Exit(1)
// 	}
// 	defer listener.Close()
//
// 	fmt.Println("Listening on ", address)
//
// 	// Loop to be able to handle multiple connections.
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection:", err)
// 			continue
// 		}
// 		go handleConnection(conn)
// 	}
// }
//
// func handleConnection(conn net.Conn) {
// 	fmt.Println("Client connected:", conn.RemoteAddr())
// 	var state status.State = status.TALK
//
// 	for {
// 		time.Sleep(1 * time.Second)
// 		msg := messages.NewMessage("Hi Friend")
//
// 		switch state {
// 		case status.TALK:
// 			printC(conn, state, msg)
// 			sendMessage(conn, msg)
//
// 			state = status.HEAR
// 		case status.HEAR:
// 			printC(conn, state, messages.EmptyMessage())
//
// 			state = status.STOP
// 		case status.STOP:
// 			printC(conn, state, messages.EmptyMessage())
//
// 			conn.Close()
// 			return
// 		default:
// 			fmt.Println("Erroneous state: ", state)
// 			state = status.STOP
// 		}
// 	}
// }

// func printC(conn net.Conn, s status.State, message messages.Message) {
// 	fmt.Println(conn.RemoteAddr(), status.Lookup(s), message)
// }
//
// func sendMessage(conn net.Conn, msg messages.Message) {
// 	_, err := fmt.Fprintf(conn, msg.String())
// 	if err != nil {
// 		fmt.Println("Error Sending Message: ", err)
// 	}
// }
