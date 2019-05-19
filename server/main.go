package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err.Error())
	}

	defer listener.Close()

	fmt.Println("Listening on 127.0.0.1:8081...")

	for {
		connection, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	address := connection.RemoteAddr()

	for {
		message, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Клиент("+address.String()+"): ", message)
	}
}
