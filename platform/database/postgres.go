package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// 데이터베이스 기본 선언
var postgresClient *gorm.DB

// PostgresClient Postgres 데이터베이스 연결
func PostgresClient() (*gorm.DB, error) {
	var err error
	if postgresClient == nil {
		// 데이터베이스 생성
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB_NAME"),
			os.Getenv("POSTGRES_SSL_MODE"),
		)
		postgresClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	return postgresClient, err
}
