package models

import(
    "reflect"

    "golang.org/x/crypto/bcrypt"

    "github.com/gabz75/auth-api/core"
)

// User Model
type User struct {
    ID int64 `json:"id"`
    Email string `json:"email"`
    Password string `json:"password"`
}

// Schema -
func (user *User) Schema() core.Mappings {
  return core.Mappings{
    core.Mapping{
      "Email",
      "email",
      reflect.String,
    },
    core.Mapping {
      "Password",
      "password",
      reflect.String,
    },
  }
}

// Table -
func (user *User) Table() string {
  return "users"
}


// Valid -
func (user *User) Valid() bool {
    if user.Email == "" {
        return false
    }
    if user.Password == "" {
        return false
    }

    return true
}

// Save -
func (user *User) Save() {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

    if err != nil {
        panic(err)
    }

    user.Password = string(hashedPassword)

    if _, err := core.InsertQuery(user); err != nil {
        panic(err)
    }
}

// Destroy -
func (user *User) Destroy() {
    db := core.DatabaseConnection()

    if _, err := db.Exec("DELETE FROM users WHERE id = $1", user.ID); err != nil {
        panic(err)
    }
}
