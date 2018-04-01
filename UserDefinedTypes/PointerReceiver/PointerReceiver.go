/*
Program that retreives the size of a url body, storing the data in a user-defined struct
using pointers/receivers to carry out the operations
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// user-defined type that encapsulates a webpage
type webPage struct {
	url  string
	body []byte
	err  error
}

// getURL() is downloading the webpage
// (w *webPage): "*" pointer receiver: when using structs -> use pointers if the variables are modified
// using direct values wont work (all would be nil) -> another copy of the variable would affected in the function
// this would not be necessary with value types such as slices or maps -> as the value is passed around
// or in case the struct is not modified in the function
// -> getURL() on a instance "w" of webPage type
func (w *webPage) getURL() {
	// use the url passed through the webPage struct to store the data and error in resp, err
	resp, err := http.Get(w.url)
	// if error, store the error into the struct "err" variable
	if err != nil {
		w.err = err
		return
	}
	// defer body closure
	defer resp.Body.Close()

	// get the content of the url body (byte[]) from resp.Body
	// and store it in the struct "body" variable
	w.body, err = ioutil.ReadAll(resp.Body)
	// if error, store the error into the struct "err" variable
	if err != nil {
		w.err = err
	}
}

// function that store "nil" in w.err if there was no error
func (w *webPage) isOK() bool {
	return w.err == nil
}

func main() {
	// using a pointer with "&" (or not) to declare w as webPage type
	// and fill it in with a url
	w := &webPage{url: "http://www.concur.com"}

	// calling the function with that webPage instance
	w.getURL()

	// embedding the printing of output within the isOK function
	if w.isOK() {
		// printing out the instance's data
		fmt.Printf("URL: %s Error: %s Length: %d", w.url, w.err, len(w.body))
	} else {
		fmt.Printf("Something went wrong: %s", w.err)
	}

	fmt.Printf("\n")
}
