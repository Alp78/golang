package main

import (
	poetry "Poetry"
	"fmt"
)

func main() {

	poem := poetry.Poem{{"Come away, come away, death,",
		"And in sad cypress let me be laid.",
		"Fly away, fly away, breath;",
		"I am slain by a fair cruel maid.",
		"My shroud of white, stuck all with yew,",
		"O, prepare it!",
		"My part of death, no one so true",
		"Did share it."}, {"Not a flower, not a flower sweet,",
		"On my black coffin let there be strown.",
		"Not a friend, not a friend greet",
		"My poor corpse, where my bones shall be thrown.",
		"A thousand thousand sighs to save,",
		"Lay me, O, where",
		"Sad true lover never find my grave,",
		"To weep there!"}}

	fmt.Printf("\n")

	// %v prints in the inferred format: here a slice (of slices..)
	fmt.Printf("%v\n\n", poem)
	fmt.Printf("%v\n\n", poem[0])
	fmt.Printf("%v\n\n", poem[0][0])
	// runes are expressed in unint8: UTF-8 representation
	// string conversion needed to get the corresponding character
	fmt.Printf("%v\n\n", string(poem[0][0][0]))
	// trying to print it with %s will throw an error
	fmt.Printf("%s\n\n", poem[0][0][0])

	// %#v prints the syntax
	fmt.Printf("%#v\n\n", poem)
	fmt.Printf("%#v\n\n", poem[0])

	// %T prints the type
	fmt.Printf("%T\n\n", poem)
	fmt.Printf("%T\n\n", poem[0])
	fmt.Printf("%T\n\n", poem[0][0])
	// runes are expressed in unint8
	fmt.Printf("%T\n\n", poem[0][0][0])

	//%q escapes special chars
	q := "\nNo Escape\n"
	fmt.Printf("%q", q)
	fmt.Printf("\n\n")

	//%x prints in hexadecimal: here a line of runes
	// adding a space between % and x will make it more readable
	fmt.Printf("% x\n\n", poem[0][0])

	// prints the output of a method
	fmt.Printf("%v\n\n", poem.String())

	// stores the returned Poem from LoadPoem() method
	poemFile, err := poetry.LoadPoem("shakespeare.txt")
	if err != nil {
		fmt.Printf("Error loading poem: %s\n", err)
	}

	fmt.Printf("%v\n\n", poemFile)

	// pulling out stats to demonstrate poemFile is indeed of Poem type
	vows, cons, puns := poemFile.Stats()
	fmt.Printf("Vowels: %d\nConsonants: %d\nPunctuation Marks: %d\n\n", vows, cons, puns)
	fmt.Printf("Stanzas: %d\nLines: %d\n", poemFile.PoemStanzas(), poemFile.PoemLines())
	fmt.Printf("\n")

}
