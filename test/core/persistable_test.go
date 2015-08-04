package core

import (
  "testing"

  "github.com/gabz75/auth-api/core"
  "github.com/gabz75/auth-api/models"
  "github.com/stretchr/testify/assert"
)


func TestCreateUser(t *testing.T) {
    var password = "qwertyuiop"
    var email = "test@gabe.com"
    user := models.User{Email: email, Password: password }
    user.Save()

    entry, err := models.GetUserByEmailAndPassword(email, password)

    assert.NotEqual(t, entry.ID, 0)
    assert.NotNil(t, entry)
    assert.Nil(t, err)

    entry.Destroy()

    entry, err = models.GetUserByEmailAndPassword(email, password)
    assert.Nil(t, entry)
    assert.NotNil(t, err)
}

func TestInsertTextQuery(t *testing.T) {
    var password = "qwertyuiop"
    var email = "test@gabe.com"
    user := models.User{Email: email, Password: password }
    user.Save()

    var token = core.GenerateToken()
    session := models.Session{UserID: user.ID, Token: token}

    insert := core.InsertTextQuery(&session)
    assert.Equal(t,  insert, "INSERT INTO sessions (user_id, token) VALUES ($1, $2)")

    parameters := core.InsertParametersQuery(&session)
    assert.Equal(t, parameters, []interface{}{ user.ID, token })

    entry, _ := models.GetUserByEmailAndPassword(email, password)
    entry.Destroy()
}
