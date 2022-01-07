package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	// for deployment to heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Up and running...")
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPosts).Methods("POST")

	log.Println("Server is listening on port: ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
