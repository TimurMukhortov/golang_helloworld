package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("hello, world\n")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer print string started from Hello, *and path*
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
