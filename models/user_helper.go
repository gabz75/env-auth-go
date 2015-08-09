package models

import (
    "errors"
    "fmt"

    "golang.org/x/crypto/bcrypt"

    "github.com/gabz75/auth-api/core/orm"
)

// GetUserByEmailAndPassword fetch user by email
func GetUserByEmailAndPassword(email string, password string) (*User, error) {
    db := orm.DatabaseConnection()

    rows, err := db.Query("SELECT * FROM users WHERE email = $1", email)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var user User;
    var id int64
    var userEmail string
    var hashedPassword string

    for rows.Next() {
        err = rows.Scan(&id, &userEmail, &hashedPassword)
        if err != nil {
            panic(err)
        }
        user = User{ ID: id, Email: userEmail, Password: hashedPassword }
    }

    nomatch := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

    if nomatch != nil || user.ID == 0 {
        return nil, errors.New("invalid email/password")
    }

    return &user, nil
}

// GetUserByToken -
func GetUserByToken(token string) (*User, error) {
    db := orm.DatabaseConnection()

    query := "SELECT users.* FROM users INNER JOIN sessions ON sessions.user_id = users.id WHERE sessions.token = $1"
    rows, err := db.Query(query, token)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var user User
    var id int64
    var userEmail string
    var hashedPassword string

    for rows.Next() {
        err = rows.Scan(&id, &userEmail, &hashedPassword)
        if err != nil {
            panic(err)
        }
        user = User{ ID: id, Email: userEmail, Password: hashedPassword }
    }

    if user.ID == 0 {
        return nil, errors.New("invalid session")
    }

    return &user, nil
}

// AvailableEmail -
func AvailableEmail(email string) bool {
    db := orm.DatabaseConnection()

    rows, err := db.Query("SELECT count(*) FROM users WHERE email = $1", email)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var count int64

    for rows.Next() {
        err = rows.Scan(&count)
        if err != nil {
            panic(err)
        }
    }

    fmt.Println(count)

    return count == 0
}
