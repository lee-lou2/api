package utils

import (
	"context"
	"github.com/lee-lou2/hub/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// MongoDBWriter MongoDB Writer
type MongoDBWriter struct {
	collection *mongo.Collection
}

// Write 로그 등록
func (w *MongoDBWriter) Write(p []byte) (int, error) {
	message := string(p)
	_, err := w.collection.InsertOne(context.Background(), bson.M{"message": message})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// SetLogger 로그 설정
func SetLogger() {
	client, err := database.MongoClient()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// MongoDBWriter 생성
	collection := client.Database("logs").Collection("api")
	_, err = collection.InsertOne(
		context.Background(),
		bson.M{"message": "Application started."},
	)

	if err != nil {
		log.Fatal(err)
	}
	writer := &MongoDBWriter{collection: collection}

	// 로그 출력 방식 변경
	log.SetOutput(writer)
}
