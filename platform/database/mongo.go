package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type NewCollection struct {
	*mongo.Collection
}

func MongoClient() (*mongo.Client, error) {
	// MongoDB 클라이언트 설정
	uri := fmt.Sprintf(
		"mongodb://%s:%s",
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
	)
	userName := os.Getenv("MONGO_USER_NAME")
	password := os.Getenv("MONGO_PASSWORD")
	clientOptions := options.Client().ApplyURI(uri)

	// 인증 정보 추가
	credential := options.Credential{
		Username: userName,
		Password: password,
	}
	clientOptions.Auth = &credential

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetCollection(dbName string, colName string) (*mongo.Client, *NewCollection, error) {
	// 컬렉션 선택
	client, _ := MongoClient()
	collection := &NewCollection{client.Database(dbName).Collection(colName)}
	return client, collection, nil
}

func (c *NewCollection) InsertOneDocument(document interface{}) (primitive.ObjectID, error) {
	// 문서 등록
	insertResult, err := c.InsertOne(context.Background(), document)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func (c *NewCollection) GetDocumentByID(id primitive.ObjectID, document interface{}) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 문서 조회
	err := c.FindOne(context.Background(), filter).Decode(document)
	if err != nil {
		return err
	}

	return nil
}

func (c *NewCollection) DeleteDocumentByID(id primitive.ObjectID) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 문서 삭제
	deleteResult, err := c.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("document not found")
	}

	return nil
}

func (c *NewCollection) UpdateDocumentByID(id primitive.ObjectID, update interface{}) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 문서 업데이트
	updateResult, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("document not found")
	}

	return nil
}

func (c *NewCollection) PushToDocumentArray(id primitive.ObjectID, field string, values ...interface{}) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 필드에 값을 추가
	update := bson.M{"$push": bson.M{field: bson.M{"$each": values}}}
	updateResult, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("document not found")
	}

	return nil
}

func (c *NewCollection) IncrementDocumentField(id primitive.ObjectID, field string, value int) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 필드의 값을 증가
	update := bson.M{"$inc": bson.M{field: value}}
	updateResult, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("document not found")
	}

	return nil
}
