package controllers

import(
    "net/http"
    "encoding/json"

    "github.com/gabz75/auth-api/models"
    "github.com/gabz75/auth-api/services"
)

type AuthenticationToken struct {
	Token string `json:"token"`
}

type AuthenticationError struct {
	Message string `json:"error"`
}

// PostSession - Generate a token given a valid email/password
func PostSession(w http.ResponseWriter, r *http.Request) {
	db := services.DatabaseConnection()
	decoder := json.NewDecoder(r.Body)

    var user models.User;

    err := decoder.Decode(&user)

    if err != nil {
        panic(err)
    }

    result := models.GetUserByEmailAndPassword(db, user.Email, user.Password)

    if result != nil {
    	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(401)
	    if err := json.NewEncoder(w).Encode(AuthenticationError{Message: "invalid credentials"}); err != nil {
	        panic(err)
	    }
    } else {
    	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(201)
	    if err := json.NewEncoder(w).Encode(AuthenticationToken{Token: services.GenerateToken()}); err != nil {
	        panic(err)
	    }
    }
}

// GetSessions -
func GetSessions(w http.ResponseWriter, r *http.Request) {

}