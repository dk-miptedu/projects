package db

import (
	"fmt"
	"golang_hws/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mplspsql"
	password = "mplspsql"
	dbname   = "transactionsdb"
)

var DB *gorm.DB

func Connect() {
	var err error
	//	dsn := "host=localhost user=user password=password dbname=transactionsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database\n", err)
		os.Exit(2)
	}
	log.Println("Connection Opened to Database")
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Transaction{})
	log.Println("Database Migration Completed")
}
