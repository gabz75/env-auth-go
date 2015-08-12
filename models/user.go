package models

import(
    "reflect"
    "errors"

    "golang.org/x/crypto/bcrypt"

    "github.com/gabz75/auth-api/core/orm"
)

// User model
type User struct {
    ID int64 `json:"id"`
    Email string `json:"email"`
    Password string `json:"password"`
}

// Schema - mapping between model and DB
func (user *User) Schema() orm.Mappings {
    return orm.Mappings{
        orm.Mapping{
            "Email",
            "email",
            reflect.String,
        },
        orm.Mapping {
            "Password",
            "password",
            reflect.String,
        },
    }
}

// Table - DB table name
func (user *User) Table() string {
    return "users"
}


// Valid - Verify if the model is valid before inserting into db
func (user *User) Valid() error {
    // @TODO validate format of email and password
    if user.Email == "" {
        return errors.New("email required")
    }

    if user.Password == "" {
        return errors.New("password required")
    }

    return nil
}

// Save - insert session in DB
func (user *User) Save() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

    if err != nil {
        return err
    }

    user.Password = string(hashedPassword)

    if _, err := orm.Insert(user); err != nil {
        return err
    }

    id, err := orm.LastInsertedID(user)

    if err != nil {
        return err
    }

    user.ID = id

    return nil
}

// Destroy - delete session from DB
func (user *User) Destroy() error {
    if _, err := orm.Delete(user); err != nil {
        return err
    }

    return nil
}
