package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	post []Post
)

// to run when the program initializes
func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}

}

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	// encode the array into the json format
	result, err := json.Marshal(posts)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
	return
}
