package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	ID          uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	StudentName string `json:"student"`
	ClassID     uint64 `json:"class_id" gorm:"not null"`
	Class       Class  `gorm:"belongs_to:ClassID;"`
}

type Class struct {
	ID        uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ClassName string `json:"class"`
}

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=louis password=123 dbname=studentdb port=5432 sslmode=disable TimeZone=Europe/Paris"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Student{}, &Class{})
}
