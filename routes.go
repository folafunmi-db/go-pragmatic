package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/folafunmi-db/go-pragmatic/entity"
	"github.com/folafunmi-db/go-pragmatic/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	posts, err := repo.FindAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func addPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	post.Id = rand.Int63()
	repo.Save(post)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
