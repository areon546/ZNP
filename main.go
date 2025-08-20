package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	role    string = ""
	address IP4    = ""
	port    Port   = "8961"
)

func main() {
	fmt.Println("I am A Client Main")

	flag.StringVar(&role, "r", role, "Specify whether the program should run as a Client or a Server")
	flag.StringVar(&ip, "ip", "", "Specify the IP address you are making calls to as a client")

	flag.Parse()

	determineRole()
}

func determineRole() {
	switch role {
	case "server":
		runServer()
	case "client":
		runClient()
	default:
		print("Unexpected role:", role)
	}
}

func runServer() {
	s := NewZNPServer(port)
	s.Run()
}

func runClient() {
	c := NewZNPClient(8691, os.Stdin)
	c.Connect(ip)
}

// Misc
func print(a ...any) {
	fmt.Println(a...)
}
