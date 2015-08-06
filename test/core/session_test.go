package core

import (
  "testing"

  "github.com/gabz75/auth-api/core"
  "github.com/gabz75/auth-api/models"
  "github.com/stretchr/testify/assert"
)

func CreateUser() *models.User {
    var password = "qwertyuiop"
    var email = "test@gabe.com"
    user := models.User{Email: email, Password: password }
    user.Save()

    entry, _ := models.GetUserByEmailAndPassword(email, password)

    return entry
}

func TestInsertTextQuery(t *testing.T) {
    user := CreateUser()

    session := models.Session{UserID: user.ID}
    session.Save()

    user.Destroy()
}
