package main

import (
	"fmt"
	"math/rand"
	"time"
)

type col []int

func main() {
	c := col{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c.shuffle1()
	fmt.Println(c)
	fmt.Println("")
	c = col{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c.shuffle2()
	fmt.Println(c)
}

// method 1
func (s col) shuffle1() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range s {
		newPosition := r.Intn(len(s) - 1)
		s[i], s[newPosition] = s[newPosition], s[i]
	}
}

// method 2
func (s col) shuffle2() {
	source := int64(time.Now().Nanosecond())
	rand.Seed(source)
	for i := range s {
		newPosition := rand.Intn(len(s) - 1)
		s[i], s[newPosition] = s[newPosition], s[i]
	}
}
