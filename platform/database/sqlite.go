package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

// 데이터베이스 기본 선언
var sqliteClient *gorm.DB

// SQLiteClient SQLite 데이터베이스 연결
func SQLiteClient() (*gorm.DB, error) {
	var err error
	if sqliteClient == nil {
		// 데이터베이스 생성
		dsn := os.Getenv("SQLITE_HOST")
		sqliteClient, err = gorm.Open(
			sqlite.Open(dsn), &gorm.Config{},
		)
	}
	return sqliteClient, err
}
