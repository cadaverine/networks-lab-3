package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err.Error())
	}

	defer listener.Close()

	fmt.Println("Listening on 127.0.0.1:8081...")

	connectionsNum := 0

	for {
		connection, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}

		connectionsNum++
		fmt.Println("New connection. Id: ", connectionsNum)

		fileName := "test-file-" + strconv.Itoa(connectionsNum) + ".txt"

		go handleConnection(connection, fileName)
	}
}

func handleConnection(connection net.Conn, fileName string) {
	defer connection.Close()

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	io.Copy(file, connection)
}
