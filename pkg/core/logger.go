package core

import (
	"context"
	"fmt"
	"github.com/lee-lou2/api/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// NewWriter 신규 작성자
type NewWriter struct{}

func (w NewWriter) Write(p []byte) (int, error) {
	message := string(p)

	go func(message string) {
		client, collection, err := database.GetCollection("logs", "api")
		if err != nil {
			panic(err)
		}
		defer client.Disconnect(context.Background())

		if _, err := collection.InsertOneDocument(
			bson.M{"message": message},
		); err != nil {
			fmt.Println(err)
		}
	}(message)

	return len(p), nil
}

// SetLogger 로그 설정
func SetLogger() {
	log.SetOutput(&NewWriter{})
}
