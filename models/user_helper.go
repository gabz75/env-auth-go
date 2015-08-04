package models

import (
    "errors"
    "golang.org/x/crypto/bcrypt"

    "github.com/gabz75/auth-api/core"
)

// GetUserByEmailAndPassword fetch user by email
func GetUserByEmailAndPassword(email string, password string) (*User, error) {
    db := core.DatabaseConnection()

    rows, err := db.Query("SELECT * FROM users WHERE email = $1", email)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var user User;
    var id int64
    var _email string
    var hashedPassword string

    for rows.Next() {
        err = rows.Scan(&id, &_email, &hashedPassword)
        user = User{ ID: id, Email: _email, Password: hashedPassword }
    }

    nomatch := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

    if nomatch != nil || user.ID == 0 {
        return nil, errors.New("invalid email/password")
    }

    return &user, nil
}
