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

// BadRequest - 400 HTTP response
func BadRequest(err error, w http.ResponseWriter, r *http.Request) {
    ErrorHandler(err.Error(), http.StatusBadRequest, w, r)
}

// UnprocessableEntity - 422 HTTP response
func UnprocessableEntity(err error, w http.ResponseWriter, r *http.Request) {
    ErrorHandler(err.Error(), 422, w, r)
}

// InternalServerError - 500 HTTP response
func InternalServerError(err error, w http.ResponseWriter, r *http.Request) {
    core.LogFatal(err.Error())
    ErrorHandler("Internal Server Error", http.StatusInternalServerError, w, r)
}

// ErrorHandler - error handler
func ErrorHandler(errorMessage string, httpStatus int, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(httpStatus)
    if err := json.NewEncoder(w).Encode(ErrorMessage{ Message: errorMessage }); err != nil {
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

// Unauthorized - send unauthorized http status code
func Unauthorized(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusUnauthorized)
    if err := json.NewEncoder(w).Encode(ErrorMessage{ Message: "Unauthorized" }); err != nil {
        panic(err)
    }
}
