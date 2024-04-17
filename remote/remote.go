package remote

import (
	"fmt"
	"net"
)

func Start() {
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error occured during connection: ", r)
			conn.Close()
			fmt.Println("Connection closed")
		}
	}()

	buffer := make([]byte, 1024)
	conn.Read(buffer)

	fmt.Println(string(buffer))
}
