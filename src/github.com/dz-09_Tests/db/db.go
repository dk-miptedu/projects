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
	Host     = "localhost"
	Port     = 5432
	User     = "mplspsql"
	Password = "mplspsql"
	Dbname   = "transactionsdb"
)

var DB *gorm.DB

func Connect() {
	var err error
	//	dsn := "host=localhost user=user password=password dbname=transactionsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", Host, Port, User, Password, Dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database\n", err)
		os.Exit(2)
	}
	log.Println("Connection Opened to Database")
}

func Migrate() {
	// DB.AutoMigrate(&models.Users{}, &models.Transactions{})
	// log.Println("Database Migration Completed")
	DB.AutoMigrate(&models.Users{})
	log.Println("Model USERS Migration Completed")
	DB.AutoMigrate(&models.Transactions{})
	log.Println("Model Transactions Migration Completed")
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to close database connection\n", err)
	}
	sqlDB.Close()
	log.Println("atabase Migration Completed")
}
