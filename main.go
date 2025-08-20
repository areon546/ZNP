package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	role    string = ""
	address IP4    = ""
	port    Port   = "8691"
)

func main() {
	flag.StringVar(&role, "r", role, "Specify whether the program should run as a Client or a Server")
	flag.StringVar(&ip, "ip", "", "Specify the IP address you are making calls to as a client")

	flag.Parse()

	determineRole()
}

func determineRole() {
	switch role {
	case "m":
		runMessager()
	case "server":
		runServer()
	case "client":
		runClient()
	default:
		print("Unexpected role:", role)
	}
}

func runMessager() {
	m := NewMessageConnection(port, ProcessState)

	root := IP4("127.0.0.1")
	m.Connect(root)
	m.Close()
}

func runServer() {
	s := NewZNPServer(port)
	s.conn.Host(port)
}

func runClient() {
	c := NewZNPClient(port, os.Stdin)
	ip := IP4(address)
	c.Connect(ip)
}

// Misc
func print(a ...any) {
	fmt.Println(a...)
}
