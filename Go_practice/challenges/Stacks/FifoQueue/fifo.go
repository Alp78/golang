package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	fifo []int
}

func main() {

	file, err := os.Open("./input/input16.txt")
	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	queries := make([][]int, 0)
	for {
		strLine := make([]string, 0)

		line, err := reader.ReadString('\n')
		strLine = strings.Fields(line)

		intLine := make([]int, len(strLine))

		for i := range strLine {
			tempInt, _ := strconv.Atoi(strLine[i])
			intLine[i] = tempInt
		}
		queries = append(queries, intLine)
		if err == io.EOF {
			break
		}
	}

	queries = append(queries[:0], queries[1:]...)

	q := &queue{
		fifo: []int{},
	}

	for _, query := range queries {
		switch query[0] {
		case 1:
			q.put(query[1])
		case 2:
			q.pop()
		case 3:
			q.peek()
		}
	}
}

func (q *queue) put(a int) {
	q.fifo = append(q.fifo, a)
}

func (q *queue) pop() {
	q.fifo = append(q.fifo[:0], q.fifo[1:]...)
}

func (q *queue) peek() {
	fmt.Println(q.fifo[0])
}
