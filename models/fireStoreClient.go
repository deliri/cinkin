package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// FsClient is a firestore Client to the database
func FsClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "sneakymommies"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	return client
}
