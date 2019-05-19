package main

import (
	"fmt"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	for {
		scanned := ""

		fmt.Scanln(&scanned)

		_, err = connection.Write([]byte(scanned + "\n"))
		if err != nil {
			panic(err.Error())
		}
	}
}
