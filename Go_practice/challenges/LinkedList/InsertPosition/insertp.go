package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func main() {
	file, err := os.Open("./input/input00.txt")
	checkError(err)
	reader := bufio.NewReader(file)

	writer := bufio.NewWriter(os.Stdout)

	// read the first line of the file
	llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	// initialize an empty list
	llist := SinglyLinkedList{}

	// fill the list with the next N lines
	for i := 0; i < int(llistCount); i++ {
		// converts the string of one line to int64
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		// converts to int32
		llistItem := int32(llistItemTemp)
		// append the int32 to the list
		llist.insertNodeIntoSinglyLinkedList(llistItem)
	}

	dataTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	data := int32(dataTemp)

	positionTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	position := int32(positionTemp)

	llist_head := insertNodeAtPosition(llist.head, data, position)

	printSinglyLinkedList(llist_head, " ", writer)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {

	insertNode := &SinglyLinkedListNode{
		data: data,
		next: nil,
	}

	if position == 0 {
		insertNode.next = llist
		llist = insertNode
	} else {
		count := 1
		beforeNode := llist
		for {
			currentNode := beforeNode.next
			if count == int(position) {
				insertNode.next = currentNode
				beforeNode.next = insertNode
				break
			}
			beforeNode = beforeNode.next
			count++
		}
	}
	return llist
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		} else if node == nil {
			fmt.Fprintf(writer, "\n")
		}
	}
}
