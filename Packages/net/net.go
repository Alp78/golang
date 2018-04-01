/*
program that cretates and starts a web server

ListenAndServe() starts an HTTP server with a given address and handler.
The handler is usually nil, which means to use DefaultServeMux.
Handle and HandleFunc add handlers to DefaultServeMux
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"poetry"
	"sort"
)

type config struct {
	// defining the keys in json object `json:""`
	Route      string   `json:"route"`
	BindAdress string   `json:"address"`
	ValidPoems []string `json:"valid"`
	Tokens     []string `json:"tokens"`
}

// creates a config struct instance to store the decoded json
var c config

type poemWithTitle struct {
	Title  string
	Body   poetry.Poem
	Words  int
	Tokens int
}

func poemHandler(w http.ResponseWriter, r *http.Request) {

	// htttp Request
	// need to parse "r"
	r.ParseForm()
	// assigns the file name to a variable
	// ex: http://localhost:8088/poem?name=shakespeare
	poemName := r.Form["name"][0]

	// if the poemName is in valid poems (json config)
	found := false
	for _, v := range c.ValidPoems {
		if v == poemName {
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Not found (invalid)", http.StatusNotFound)
		return
	}

	// load the poem from text file
	poem, err := poetry.LoadPoem(poemName)

	numWords := poem.NumWords()

	// passing value of Tokens field of JSON file
	wordToken := poem.NumTokenWord(c.Tokens)

	// if the poem isn't found
	if err != nil {
		http.Error(w, "File not found", http.StatusInternalServerError)
	} else {
		// creating an instance of struct to contain a title, body and number of words
		// and affecting the values of the loaded poem
		pwt := poemWithTitle{poemName, poem, numWords, wordToken}
		// format into JSON
		enc := json.NewEncoder(w)
		enc.Encode(pwt)

		// Fprintf(w io.writer, format string): formats according to format specifier
		// returns the number of bytes written
		fmt.Fprintf(w, "Poem %s :\n", poemName)

		// sorting each stanza (lines sorted by length)
		for i := 0; i < len(poem); i++ {
			sort.Sort(poem[i])
		}

		// prints the loaded poem from text file
		// http://localhost:8088/poem?name=shakespeare		// sorting the stanza lines
		fmt.Fprintf(w, "%s\n\nNumber of words: %d\nNumber of %s: %d\n", poem, poem.NumWords(), c.Tokens, poem.NumTokenWord(c.Tokens))
	}
}

func main() {
	// opens the config file containing a json object
	// f is an io.Reader
	f, err := os.Open("config")
	if err != nil {
		os.Exit(1)
	}

	dec := json.NewDecoder(f)

	// decode in "c" and stores the error
	err = dec.Decode(&c)

	// close file
	f.Close()
	// if error exit
	if err != nil {
		os.Exit(1)
	}

	// creates a url on the server: config struct field Route
	http.HandleFunc(c.Route, poemHandler)

	// starts server @ config struct field BindAddress
	http.ListenAndServe(c.BindAdress, nil)

}
