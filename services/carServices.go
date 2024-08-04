package services

import (
    "quadiro_assignment/models"
    "quadiro_assignment/config"
)

func CreateCar(car models.Car) error {
    return config.DB.Create(&car).Error
}

//
func GetCars() ([]models.Car, error) {
    var cars []models.Car
    err := config.DB.Find(&cars).Error
    return cars, err
}

func UpdateCar(car models.Car) error {
    return config.DB.Save(&car).Error
}

func DeleteCar(id string) error {
    return config.DB.Delete(&models.Car{}, id).Error
}

