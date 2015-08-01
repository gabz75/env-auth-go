package main

import(
    "database/sql"
    "golang.org/x/crypto/bcrypt"
)

// User Model
type User struct {
    ID int64 `json:"id"`
    Email string `json:"email"`
    Password string `json:"password"`
}

func (user *User) Valid() bool {
    if user.Email == "" {
        return false
    }
    if user.Password == "" {
        return false
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    user.Password = string(hashedPassword)

    return true
}

func (user *User) Save() {
    db := DatabaseConnection()
    CreateUser(db, user.Email, user.Password)
}

// CreateUser create user with email and password
func CreateUser(db *sql.DB, email string, password string) {
    _, err := db.Exec("INSERT INTO users (email, password) VALUES  ('" + email + "', '" + password + "')")

    if err != nil {
        panic(err)
    }
}

// GetUserByEmail fetch user by email
func GetUserByEmail(db *sql.DB, email string) User {
    rows, err := db.Query("SELECT * FROM users WHERE email = $1", email)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var user User;

    for rows.Next() {
        var id int64
        var email string
        var password string
        err = rows.Scan(&id, &email, &password)

        user = User{ID: id, Email: email, Password: password}
    }

    return user
}

func GetLastUser(db *sql.DB) User {
    rows, err := db.Query("SELECT * FROM users ORDER BY id desc LIMIT 1")

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var user User;

    for rows.Next() {
        var id int64
        var email string
        var password string
        err = rows.Scan(&id, &email, &password)

        user = User{ID: id, Email: email, Password: password}
    }

    return user
}
