package services

import (
    "errors"
)

// Placeholder for user authentication (for demonstration purposes)
func AuthenticateUser(username, password string) (bool, error) {
    // For simplicity, we are hardcoding a single user
    if username == "user" && password == "password" {
        return true, nil
    }
    return false, errors.New("invalid credentials")
}
