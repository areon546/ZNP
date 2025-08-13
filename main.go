package main

import (
	"flag"
	"fmt"
	"znp-cs/client"
	"znp-cs/server"
)

var role string = ""

func main() {
	flag.StringVar(&role, "r", role, "Specify whether the program should run as a Client or a Server")

	flag.Parse()

	switch role {
	case "client":
		fallthrough
	case "c":
		runClient()
	case "server":
		fallthrough
	case "s":
		runServer()
	default:
		fmt.Println("No role given")
		fmt.Println(role)
	}
}

func runClient() {
	c := client.NewZNPClient(8691)
	c.Connect("localhost")
}

func runServer() {
	s := server.NewZNPServer(8691)
	s.Run()
}
