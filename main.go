package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("hello, world\n")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/getUsers", getUsers)
	http.ListenAndServe(":8080", nil)
}

// HelloServer print string started from Hello, *and path*
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func getUsers(w http.ResponseWriter, req *http.Request) {
	userFirst := User{"Tom", 15}
	userSecnod := User{"Alex", 25}
	dataArray := []User{userFirst, userSecnod}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&dataArray)
}

// User struct
type User struct {
	Name string `json:"test"`
	Age  uint
}
