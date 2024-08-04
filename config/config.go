package config

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "quadiro_assignment/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    // Migrate the schema
    err = DB.AutoMigrate(&models.Car{})
    if err != nil {
        panic("Failed to migrate database schema!")
    }
}
