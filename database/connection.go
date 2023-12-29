package database

import (
	"github.com/project/project-skripsi/go-be/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Return Conn to postgresql
func Connection() {
	dsn := "host=localhost user=postgres password=bisma052002 dbname=project_skripsi port=5432 sslmode=disable"
	Connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("couldnt connect to the database")
	}

	DB = Connection

	Connection.AutoMigrate(&models.Users{})
}
