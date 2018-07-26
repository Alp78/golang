package main

import "fmt"

// any type that has a method called getGreeting()
// and returning a string satisfies the bot interface
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	// interfaces allow to accommodate different types for a same function logic
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
