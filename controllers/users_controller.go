package controllers

import(
    "net/http"
    "encoding/json"

    "github.com/gabz75/auth-api/models"
)

// PostUser -
func PostUser(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var user models.User;
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
