package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)

	fmt.Println("Mencoba koneksi dengan config:")
	fmt.Printf("host=%s user=%s password=%s dbname=%s port=%d\n", host, user, password, dbname, port)
	fmt.Printf("DB HOST: %s, PORT: %d, USER: %s, NAME: %s\n", host, port, user, dbname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
		os.Exit(1)
	}
	fmt.Println("Berhasil koneksi ke database")
}
