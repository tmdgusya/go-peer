package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic("Server couldn't estabilish as well")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic("Something went wrong during estabilishing connection with client")
		}
		fmt.Println(conn)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
}
