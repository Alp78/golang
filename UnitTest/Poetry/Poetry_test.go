package poetry

import (
	"testing"
)

func TestPoemStanzas(t *testing.T) {
	poem := Poem{}
	if poem.PoemStanzas() != 0 {
		t.Fatalf("Empty poem is not empty!")
	}

	poem = Poem{{"Come away, come away, death,",
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

	if poem.PoemStanzas() != 2 {
		t.Fatalf("Expected stanza count: %d", poem.PoemStanzas())
	}
}
func TestPoemLines(t *testing.T) {
	poem := Poem{}
	if poem.PoemLines() != 0 {
		t.Fatalf("Empty poem is not empty!")
	}

	poem = Poem{{"Come away, come away, death,",
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

	if poem.PoemLines() != 16 {
		t.Fatalf("Expected lines count: %d", poem.PoemLines())
	}
}
func TestStats(t *testing.T) {
	poem := Poem{}
	vows, cons, puns := poem.Stats()
	if vows != 0 || cons != 0 || puns != 0 {
		t.Fatalf("Poem not empty!")
	}

	poem = Poem{{"Hello, World!"}}
	vows, cons, puns = poem.Stats()
	if vows != 3 || cons != 7 || puns != 3 {
		t.Fatalf("Expected number of vowels: %d, consonants: %d, punctuation marks: %d", vows, cons, puns)
	}
}

func TestNumWords(t *testing.T) {
	p := Poem{{"hello"}}
	if p.NumWords() != 1 {
		t.Fatalf("Wrong number of words.")
	}

	p = Poem{{"hello world"}}
	if p.NumWords() != 2 {
		t.Fatalf("Wrong number of words.")
	}
}

func TestNumTokenWords(t *testing.T) {
	p := Poem{{"the The they They"}}
	w := []string{"the", "The"}
	if p.NumTokenWord(w) != 2 {
		t.Fatalf("Wrong token count.")
	}
}
