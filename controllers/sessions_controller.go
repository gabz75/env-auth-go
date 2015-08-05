package controllers

import(
    "net/http"
    "encoding/json"

    "github.com/gabz75/auth-api/models"
    "github.com/gabz75/auth-api/core"
)

// AuthenticationToken -
type AuthenticationToken struct {
	Token string `json:"token"`
}

// PostSession - Generate a token given a valid email/password
func PostSession(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

    var user models.User;

    err := decoder.Decode(&user)
    if err != nil {
        panic(err)
    }

    entry, err := models.GetUserByEmailAndPassword(user.Email, user.Password)
    if err != nil {
        Unauthorized(w, r)
        return
    }

    session := models.Session{UserID: entry.ID}
    session.Save()

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(session); err != nil {
        panic(err)
    }
}

// GetSessions -
func GetSessions(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    token := core.ExtractToken(r.Header.Get("Authorization"))
    user, err := models.GetUserFromToken(token)

    if err != nil {
        Unauthorized(w, r)
        return
    }

    if err := json.NewEncoder(w).Encode(user); err != nil {
        panic(err)
    }
}
