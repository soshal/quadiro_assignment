package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "quadiro_assignment/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("file:test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.Car{})

    DB = database
}
