package db

import (
	"goapp/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "user=postgres password=00000000 dbname=db_project2 port=3306 sslmode=disable"

func Connect() error {
	var err error

	Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	Database.AutoMigrate(&models.Word{})
	Database.AutoMigrate(&models.WordAnswer{})
	return nil
}
