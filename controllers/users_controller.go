package controllers

import(
    "net/http"
    "encoding/json"

    "github.com/gabz75/go-auth-api/models"
)

// PostUser -
func PostUser(w http.ResponseWriter, r *http.Request) {
    var user models.User;

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&user)

    if err != nil {
        BadRequest(err, w, r)
        return
    }

    if err := user.Valid(); err != nil {
        UnprocessableEntity(err, w, r)
        return
    }

    if err := user.Save(); err != nil {
        InternalServerError(err, w, r)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)

    if err := json.NewEncoder(w).Encode(user); err != nil {
        InternalServerError(err, w, r)
        return
    }
}
