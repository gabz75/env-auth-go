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
    var user models.User
    if authenticated := Authenticate(&user, w, r); authenticated {

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        sessions := models.GetSessions(user)

        if err := json.NewEncoder(w).Encode(sessions); err != nil {
            panic(err)
        }
    }
}

// DestroySession -
func DestroySession(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if authenticated := Authenticate(&user, w, r); authenticated {

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusNoContent)

        session, err := models.GetSession(user, core.ExtractToken(r.Header.Get("Authorization")))

        if err != nil {
            panic(err)
        }

        session.Destroy()
    }
}
