package core

import (
	"context"
	"github.com/lee-lou2/api/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"os"
)

// MongoDBWriter MongoDB Writer
type MongoDBWriter struct {
	collection *database.NewCollection
}

// Write 로그 등록
func (w *MongoDBWriter) Write(p []byte) (int, error) {
	message := string(p)
	_, err := w.collection.InsertOneDocument(bson.M{"message": message})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// SetLogger 로그 설정
func SetLogger() {
	client, collection, err := database.GetCollection("logs", "api")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	if _, err := collection.InsertOneDocument(
		bson.M{"message": "Application started."},
	); err != nil {
		log.Fatal(err)
	}

	// 로그 출력 방식 변경
	writer := &MongoDBWriter{collection: collection}
	log.SetOutput(io.MultiWriter(os.Stdout, writer))
}
