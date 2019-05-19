package main

import (
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	connection.Write([]byte("Я клиент!\n"))
}
