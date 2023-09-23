package database

import (
	"fmt"
	"github.com/project_management/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota", host, user, password, dbname, port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo conectar a la base de datos", err)
	}
	log.Println("conectado a la base de datos", conn)
	DB = conn
	conn.AutoMigrate(&model.User{})
	conn.AutoMigrate(&model.Project{})
	conn.AutoMigrate(&model.Task{})
}
