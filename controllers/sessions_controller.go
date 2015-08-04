package controllers

import(
    "net/http"
    "encoding/json"

    "github.com/gabz75/auth-api/models"
)

// AuthenticationToken -
type AuthenticationToken struct {
	Token string `json:"token"`
}

// ErrorMessage -
type ErrorMessage struct {
	Message string `json:"error"`
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
    w.WriteHeader(201)
    if err := json.NewEncoder(w).Encode(session); err != nil {
        panic(err)
    }
}

// GetSessions -
func GetSessions(w http.ResponseWriter, r *http.Request) {

}

// Unauthorized -
func Unauthorized(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusUnauthorized)
    if err := json.NewEncoder(w).Encode(ErrorMessage{ Message: "Unauthorized" }); err != nil {
        panic(err)
    }
}
