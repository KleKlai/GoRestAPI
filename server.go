package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	const port string = ":8080"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/post", addPost).Methods("POST")

	log.Printf("Server is running on port %s", port)

	http.ListenAndServe(port, router)
}
