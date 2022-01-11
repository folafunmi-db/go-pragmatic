package repository

import (
	"context"
	"log"
	"../entity"
	"cloud.google.com/go/firestore"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

// NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId string = "go-reviews"
	collectionName string = "posts"
)

func Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id": post.Id,
		"Title": post.Title,
		"Text": post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to add a new post: %v", err)
		return nil, err
	}

	// if there are no error return the posts and nil
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
		return nil, err
		}

		post := entity.Post{
			Id: doc.Data()["Id"](int64),
			Title: doc.Data()["Title"](string),
			Text: doc.Data()["Text"](string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
