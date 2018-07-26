package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server...")

	// create listener
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	// listen and accept connections from clients
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}

}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 256)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("%v off\n", conn.RemoteAddr())
			return
		}
		trimBytes := bytes.Trim(buf, "\x00")
		fmt.Printf("%v\n", string(trimBytes))
	}
}
