package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB = nil

func InitDatabase() error {
	var err error
	// dbConfig will contains information about the database, it is the result of unmarshal the config.json file.
	dbConfig, err := GetConfiguration()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBname, dbConfig.DBport)
	db, err := gorm.Open("postgres", dsn)

	// restriction of the connection pool.
	db.DB().SetConnMaxIdleTime(1)
	db.DB().SetConnMaxLifetime(5)

	DB = db

	return nil
}
