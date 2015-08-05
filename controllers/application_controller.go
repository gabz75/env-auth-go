package controllers

import (
    "net/http"
    "encoding/json"
)

// ErrorMessage -
type ErrorMessage struct {
    Message string `json:"error"`
}

// Unauthorized -
func Unauthorized(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusUnauthorized)
    if err := json.NewEncoder(w).Encode(ErrorMessage{ Message: "Unauthorized" }); err != nil {
        panic(err)
    }
}
