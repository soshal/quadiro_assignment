package main

import (
    "quadiro_assignment/routes"
    "quadiro_assignment/config"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := routes.SetupRouter()

    // Initialize the database connection
    config.ConnectDatabase()

    r.LoadHTMLGlob("templates/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Assignment for Quadiro Technologies",
        })
    })

    // Run the server
    r.Run(":8080")
}
