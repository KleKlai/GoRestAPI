package repostiroy

import (
	"context"
	"log"

	"github.com/KleKlai/GoRestAPI/entity"
	"google.golang.org/api/firestore/v1"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct {
}

func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	credentialsFile string = "gorestapi-3b6ab-firebase-adminsdk-5q9xu-8b5f5b1b5f.json"
	projectId       string = "gorestapi-3b6ab"
	collectionName  string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {

	ctx := context.Background()

	// credentialsJSON, err := ioutil.ReadFile(credentialsFile)
	// if err != nil {
	// 	log.Fatalf("Failed to read credentials file: %v", err)
	// }

	// creds, err := google.CredentialsFromJSON(ctx, credentialsJSON, firestore.CloudPlatformScope)
	// if err != nil {
	// 	log.Fatalf("Failed to create Firestore client: %v", err)
	// }

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})
}
