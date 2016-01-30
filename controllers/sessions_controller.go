package controllers

import(
    "net/http"
    "encoding/json"

    "github.com/gabz75/env-auth/models"
    "github.com/gabz75/env-auth/core"
)

// PostSession - Generate a token given a valid email/password
func PostSession(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

    var user models.User;

    err := decoder.Decode(&user)
    if err != nil {
        BadRequest(err, w, r)
        return
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
        InternalServerError(err, w, r)
        return
    }
}

// GetSessions - List all available sessions
func GetSessions(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if authenticated := Authenticate(&user, w, r); !authenticated {
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    sessions := models.GetSessions(user)

    if err := json.NewEncoder(w).Encode(sessions); err != nil {
        InternalServerError(err, w, r)
        return
    }
}

// DestroySession - Logout
func DestroySession(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if authenticated := Authenticate(&user, w, r); !authenticated {
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNoContent)

    session, err := models.GetSession(user, core.ExtractToken(r.Header.Get("Authorization")))

    if err != nil {
        BadRequest(err, w, r)
        return
    }

    session.Destroy()
}
