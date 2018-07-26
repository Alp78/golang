package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// func Get(url string) (resp *Response, err error)
	// https://golang.org/pkg/net/http/#Get
	// https://golang.org/pkg/net/http/#Response
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	/*
		// make: initialize the slice with n empty bytes
		// 20000: big enough to host the body content (best guess)
		bs := make([]byte, 20000)
		//Read will insert the body content in the empty byte slice
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/
	// cleaner and more efficient way:
	// func Copy(dst Writer, src Reader) (written int64, err error)
	// io.Copy(os.Stdout, resp.Body)

	// lw: implements the Writer interface
	// as it has a Write method
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// custom implementation of the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("")
	fmt.Printf("Byte slice size: %v", len(bs))
	return len(bs), nil
}
