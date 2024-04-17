package main

import "net"

func Makeconneciton() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		panic("Couldn't connect to server")
	}

	conn.Write([]byte("Hello from client"))
}

func main() {
	Makeconneciton()
}
