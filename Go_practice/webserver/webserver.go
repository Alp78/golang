package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// HelloServer writes the content of the URL path after /
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Connected to server.")
	fmt.Fprintln(w, "Hello", req.URL.Path[1:])
}
