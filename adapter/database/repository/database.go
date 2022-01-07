package repository

import (
	"e-commerce/adapter/database/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open("host=localhost user=root password=xz2k600600 dbname=ecommerce port=5433 sslmode=disable"),
		&gorm.Config{PrepareStmt: true})
	if err != nil{
		log.Fatal("Error: ", err)
	}

	config, _ := db.DB()
	config.SetMaxOpenConns(10)
	config.SetConnMaxLifetime(100)
	config.SetConnMaxLifetime(time.Hour)
	migration.RunMigration(db)
	return db

}
