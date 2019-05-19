package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	message, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("Сообщение от клиента: ", message)
}
