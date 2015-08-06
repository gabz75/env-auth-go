package controllers

import (
    "net/http"
    "encoding/json"

    "github.com/gabz75/auth-api/models"
    "github.com/gabz75/auth-api/core"
)

// ErrorMessage -
type ErrorMessage struct {
    Message string `json:"error"`
}

// Unauthorized - send unauthorized http status code
func Unauthorized(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusUnauthorized)
    if err := json.NewEncoder(w).Encode(ErrorMessage{ Message: "Unauthorized" }); err != nil {
        panic(err)
    }
}

// Authenticate - autenticate a user and return true or return false and send an unauthorized status
func Authenticate(currentUser *models.User, w http.ResponseWriter, r *http.Request) bool {
    token := core.ExtractToken(r.Header.Get("Authorization"))

    user, err := models.GetUserByToken(token)

    if err != nil {
        Unauthorized(w, r)
        return false
    }

    currentUser.ID = user.ID
    currentUser.Email = user.Email
    currentUser.Password = user.Password

    if currentUser.ID == 0 {
        Unauthorized(w, r)
        return false
    }

    return true
}
