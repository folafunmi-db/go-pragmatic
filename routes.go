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
	posts []Post
)

// to run when the program initializes
func init() {
	posts = []Post{{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// encode the array into the json format
	result, err := json.Marshal(posts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func addPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the request"}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	result, err2 := json.Marshal(posts)

	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the request"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
