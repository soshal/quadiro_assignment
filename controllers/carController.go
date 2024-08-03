package controllers

import (
    "net/http"
    "quadiro_assignment/models"
    "quadiro_assignment/services"
    "github.com/gin-gonic/gin"
)

func CreateCar(c *gin.Context) {
    var car models.Car
    if err := c.ShouldBindJSON(&car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.CreateCar(car); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": car})
}

func GetCars(c *gin.Context) {
    cars, err := services.GetCars()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": cars})
}

func UpdateCar(c *gin.Context) {
    var car models.Car
    if err := c.ShouldBindJSON(&car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.UpdateCar(car); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": car})
}

func DeleteCar(c *gin.Context) {
    id := c.Param("id")
    if err := services.DeleteCar(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": "Car deleted successfully"})
}

// New function to render the dashboard
func AdminDashboard(c *gin.Context) {
    cars, err := services.GetCars()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    totalCars := len(cars)
    c.HTML(http.StatusOK, "dashboard.html", gin.H{
        "title": "Admin Dashboard",
        "cars": cars,
        "totalCars": totalCars,
    })
}
