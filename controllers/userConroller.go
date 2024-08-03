package controllers

import (
    "net/http"
    "quadiro_assignment/services"
    "github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    if username == "admin" && password == "admin" {
        c.JSON(http.StatusOK, gin.H{"token": "admin_token"})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    }
}

func UserLogin(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    authenticated, err := services.AuthenticateUser(username, password)
    if authenticated {
        c.JSON(http.StatusOK, gin.H{"token": "user_token"})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
    }
}
