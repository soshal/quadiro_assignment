package routes

import (
    "quadiro_assignment/controllers"
    "quadiro_assignment/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    admin := r.Group("/admin")
    {
        admin.POST("/login", controllers.AdminLogin)
        admin.Use(middleware.AuthMiddleware) // Protect the following routes with authentication middleware
        admin.POST("/cars", controllers.CreateCar)
        admin.GET("/cars", controllers.GetCars)
        admin.PUT("/cars/:id", controllers.UpdateCar)
        admin.DELETE("/cars/:id", controllers.DeleteCar)
        admin.GET("/dashboard", controllers.AdminDashboard) // New dashboard route
    }

    user := r.Group("/user")
    {
        user.POST("/login", controllers.UserLogin)
        user.GET("/cars", controllers.GetCars)
    }

    return r
}
