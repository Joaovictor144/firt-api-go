package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB(){
	str := "host=localhost port=5432 user=docker dbname=api_go password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil{
		log.Fatal("Error:", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDataBase() *gorm.DB{
	return db
}