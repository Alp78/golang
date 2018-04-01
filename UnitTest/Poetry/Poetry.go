package poetry

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
	return Poem{}
}

// function that takes a string as argument (by opening a file containing some coded in UTF-8)
// scans its content with bufio methods and return it in Poem type
func LoadPoem(name string) (Poem, error) {

	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	poem := Poem{}
	stanza := Stanza{}

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			poem = append(poem, stanza)
			stanza = Stanza{}
			continue
		}

		stanza = append(stanza, Line(line))
	}
	poem = append(poem, stanza)

	if scan.Err() != nil {
		return nil, scan.Err()
	}
	return poem, nil
}

// implementing the sort.Sort methods that satisfy the interface
func (s Stanza) Len() int {
	return len(s)
}
func (s Stanza) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns True if a line is shorter
func (s Stanza) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (poem Poem) PoemStanzas() int {
	return len(poem)
}

func (s Stanza) StanzaLines() int {
	return len(s)
}

func (poem Poem) PoemLines() (count int) {
	for _, s := range poem {
		count += s.StanzaLines()
	}
	return
}

// NumWords returns number of words
func (poem Poem) NumWords() int {
	count := 0
	for _, stanza := range poem {
		for _, line := range stanza {
			// each line stored as string
			sl := string(line)
			// stores all words in a slice
			parts := strings.Split(sl, " ")
			count += len(parts)
		}
	}
	return count
}

// NumTokenWord returns the number occurrences of a word defined in a variable
func (poem Poem) NumTokenWord(token []string) int {
	count := 0
	for _, stanza := range poem {
		for _, line := range stanza {
			// each line stored as string
			sl := string(line)
			// if the string contains the token
			parts := strings.Split(sl, " ")

			for _, part := range parts {
				if part == token[0] || part == token[1] {
					count++
				}
			}
		}
	}
	return count
}

func (poem Poem) Stats() (numVows, numCons, numPuns int) {
	for _, stanza := range poem {
		for _, line := range stanza {
			for _, symbol := range line {
				switch symbol {
				case 'a', 'e', 'i', 'o', 'u':
					numVows++
				case ',', ' ', '!', '.', ';':
					numPuns++
				default:
					numCons++
				}
			}
		}
	}
	return
}

// function that has Poem as receiver -> method attached to the instance "poem"
// prints the poem with each stanza and line separated by a linebreak
func (poem Poem) String() string {
	result := ""
	for _, stanza := range poem {
		// adds a linebreak between each stanza
		result += "\n"
		for _, line := range stanza {
			// converts the line to string
			result += fmt.Sprintf("%s\n", line)
			// adds a linebreak between each line
			result += "\n"
		}
	}
	return result
}
