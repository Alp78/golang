package main

import (
	"fmt"
	"poetry"
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

	vows, cons, puns := poem.Stats()
	fmt.Printf("\n")
	fmt.Printf("Vowels: %d\nConsonants: %d\nPunctuation Marks: %d\n\n", vows, cons, puns)
	fmt.Printf("Stanzas: %d\nLines: %d\n", poem.PoemStanzas(), poem.PoemLines())
	fmt.Printf("\n")
}
