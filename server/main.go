package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

var connectionsNum int

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleError(value interface{}, err error) interface{} {
	checkError(err)
	return value
}

func handleConnection(connection net.Conn, connectionID int) {
	defer connection.Close()

	fmt.Println("New connection ID:", connectionID)

	var filesNum int
	for {
		nameBuf := new(bytes.Buffer)
		fileBuf := new(bytes.Buffer)
		nameLenBuf := new(bytes.Buffer)
		fileSizeBuf := new(bytes.Buffer)

		// получаем длину названия файла
		io.CopyN(nameLenBuf, connection, 8)
		nameSize, _ := binary.Varint(nameLenBuf.Bytes())
		// получаем название файла
		io.CopyN(nameBuf, connection, nameSize)
		fileName := nameBuf.String()
		// получаем размер файла
		io.CopyN(fileSizeBuf, connection, 8)
		fileSize, _ := binary.Varint(fileSizeBuf.Bytes())
		// получаем файл
		io.CopyN(fileBuf, connection, fileSize)

		file := handleError(os.Create(fileName)).(*os.File)
		defer file.Close()

		io.CopyN(file, fileBuf, fileSize)

		filesNum++
		fmt.Println("---------------------------")
		fmt.Println("Connection ID:     ", connectionID)
		fmt.Println("File name:         ", fileName)
		fmt.Println("File size:         ", strconv.FormatInt(fileSize, 10), "bytes")
		fmt.Println("Sent files number: ", filesNum)
		fmt.Println("---------------------------")
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	checkError(err)

	defer listener.Close()

	fmt.Println("Listening on 127.0.0.1:8081...")

	for {
		connection, err := listener.Accept()
		checkError(err)

		connectionsNum++
		go handleConnection(connection, connectionsNum)
	}
}
