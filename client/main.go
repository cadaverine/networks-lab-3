package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"path"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleError(value interface{}, err error) interface{} {
	checkError(err)
	return value
}

func main() {
	connection := handleError(net.Dial("tcp", "127.0.0.1:8081")).(net.Conn)
	defer connection.Close()

	for {
		var filePath string

		fmt.Print("Set the path of file to be send: ")
		fmt.Scanln(&filePath)

		fileName := path.Base(filePath)
		fileNameBytes := []byte(fileName)
		fileNameBytesLen := int64(len(fileNameBytes))

		fileBytes := handleError(ioutil.ReadFile(filePath)).([]byte)
		fileBytesLen := int64(len(fileBytes))

		fmt.Println("Send file to server: ", fileName)

		nameBuf := make([]byte, 8)
		binary.PutVarint(nameBuf, fileNameBytesLen)

		sizeBuf := make([]byte, 8)
		binary.PutVarint(sizeBuf, fileBytesLen)

		// длина названия файла
		io.CopyN(connection, bytes.NewReader(nameBuf), 8)
		// байты названия файла
		io.CopyN(connection, bytes.NewReader(fileNameBytes), fileNameBytesLen)
		// размер файла
		io.CopyN(connection, bytes.NewReader(sizeBuf), 8)
		// байты файла
		io.CopyN(connection, bytes.NewReader(fileBytes), fileBytesLen)
	}
}
