package main

import (
    "encoding/json"
    "net/http"
)

// GetUser -
func GetUser(w http.ResponseWriter, r *http.Request) {
    db := DatabaseConnection()

    user := GetLastUser(db)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(user); err != nil {
        panic(err)
    }
}

// PostUser -
func PostUser(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var user User;
    var status = 422;

    err := decoder.Decode(&user)

    if err != nil {
        panic(err)
    }

    if user.Valid() {
        user.Save()
        status = 200
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(status)

    if err := json.NewEncoder(w).Encode(user); err != nil {
        panic(err)
    }
}

// PostSession -
func PostSession(w http.ResponseWriter, r *http.Request) {

}
