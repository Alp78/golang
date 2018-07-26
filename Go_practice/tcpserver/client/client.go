package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// open connection
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		// no connection could be made because the target machine actively refused it
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter client name: ")
	clientName, _ := inputReader.ReadString('\n')
	trimClient := strings.Trim(clientName, "\r\n")
	//fmt.Printf("Client: %s\n", trimClient)

	// send info to server until Quit
	for {
		fmt.Printf("%s: ", trimClient)
		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")
		// fmt.Printf("Input: %s\n", trimInput)

		if trimInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimClient + ": " + trimInput))
		if err != nil {
			fmt.Println("Error writing", err.Error())
			return
		}
	}

}
