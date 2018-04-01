/*
Program that opens and reads a utf-8 CSV file comprising any number of columns and records,
stores its content in a slice, sort the content of the columns in alphanumerical order
taking the first column as reference and aligning the other columns to that order.
the program then prints out the sorted slice.

In this example a csv with 2 columns is used:
	a) English words
	b) the French translation of these words

The prgram will then sort a slice of this file taking the English words as reference,
and alignt he correct order with their corresponding French translation.

There are only 2 columns here, but any number of columns could be used, with the same resulting
order alignment.

Alternatively, this program can convert the csv in JSON format, after having sorted the slice.
*/

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// defining a Word type as struct of 2 string elements, 1 for each language
type Word struct {
	English string
	French  string
}

// defining a type as slice of Word
type byEnglish []Word

// Use the sort.Sort function to sort the slice of Word (byEnglish).
// Implement the sort interface with a Less function that compares English and then French.

// Len is the number of elements in the collection
func (a byEnglish) Len() int { return len(a) }

// Swap swaps the elements with indexes i and j
func (a byEnglish) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less reports whether the element with
// index i should sort before the element with index j
func (a byEnglish) Less(i, j int) bool {
	if a[i].English < a[j].English {
		return true
	}
	if a[i].English > a[j].English {
		return false
	}
	return a[i].French < a[j].French
}

func main() {
	// opens the csv file and allocate memory for its content to be read
	csvFile, _ := os.Open("voc_en_full.csv")

	// reads the content of the file and stored it in a buffer
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// declaring a slice to contain the csv file content
	var words []Word

	// infinite for loop until "break"
	for {
		// reads each line of the csv file and stores its content in a variable
		line, err := reader.Read()

		// if the end of the file is reached, then break the for loop
		if err == io.EOF {
			break

			// if however there is another type of error, the logs the error and go to next instruction
		} else if err != nil {
			log.Fatal(err)
		}

		// appends each element of the csv line to a respective column
		// this is where the number of columns can vary indefinitely
		words = append(words, Word{
			line[0],
			line[1],
		})
	}

	// sorting the slice leveraging the sort.Sort functions
	sort.Sort(byEnglish(words))

	// converting the slice to JSON and prints it out
	// this would however require the definition of the JSON labels for each column in the struct defintion
	// wordJSON, _ := json.Marshal(words)
	// fmt.Println(string(wordJSON))

	// printing out the content of the sorted slice
	for i := 0; i < len(words); i++ {
		fmt.Printf("%s : %s\n", words[i].English, words[i].French)
	}

}
