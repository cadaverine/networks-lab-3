package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	for {
		filePath := ""

		fmt.Print("Set the path of file to be send: ")
		fmt.Scanln(&filePath)

		binary, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Send file to server: ", filePath)

		io.Copy(connection, bytes.NewReader(binary))
	}
}
