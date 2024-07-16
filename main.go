package main

import (
	"fmt"
	"food-reserve/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func main() {
	db := ConnectDB()
	MigrateDb(db)
}
func ConnectDB() *gorm.DB {
	conn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		"127.0.0.1",
		"5432",
		"food",
		"",
		"",
		"disable")

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		//Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	// db.LogMode(true)
	// db.SetLogger(log)
	// db.SingularTable(true)
	return db
}

func MigrateDb(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Order{})
}
