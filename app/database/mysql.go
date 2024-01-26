package database

import (
	config "Laptop/app/configs"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var db *gorm.DB

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	// tidal-advantage-410213:asia-southeast2:root
	// dsn := "<USERNAME>:<PASSWORD>@tcp(<DB_ADDRESS>)/<alta1>?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "alta1:Alta1!@tcp(34.171.123.254)/alta1?charset=utf8mb4&parseTime=True&loc=Local"

	// var err error
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return db
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return DB
}

func InitRawSql(cfg *config.AppConfig) *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	// Membuka koneksi ke database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	return db
}
