package models

import(
    "database/sql"
    "golang.org/x/crypto/bcrypt"
    "errors"

    "github.com/gabz75/auth-api/services"
)

// User Model
type User struct {
    ID int `json:"id"`
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
    db := services.DatabaseConnection()
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
func GetUserByEmailAndPassword(db *sql.DB, email string, password string) error {
    rows, err := db.Query("SELECT * FROM users WHERE email = $1", email)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var user User;

    for rows.Next() {
        var id int
        var email string
        var hashedPassword string
        err = rows.Scan(&id, &email, &hashedPassword)

        user = User{ID: id, Email: email, Password: hashedPassword}
    }
    
    if user.ID == 0 {
        return errors.New("invalid email/password")
    }

    passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

    if passwordMatch != nil {
        return errors.New("invalid email/password")   
    }

    return nil
}

func GetLastUser(db *sql.DB) User {
    rows, err := db.Query("SELECT * FROM users ORDER BY id desc LIMIT 1")

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var user User;

    for rows.Next() {
        var id int
        var email string
        var password string
        err = rows.Scan(&id, &email, &password)

        user = User{ID: id, Email: email, Password: password}
    }

    return user
}
