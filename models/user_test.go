package models

import (
    "testing"
    "time"
    "math/rand"
    "strconv"

    "github.com/stretchr/testify/assert"
)

func RandomUser() *User {
    user := User{ Email: "user-" + random(9999999999999) + "@testing.com", Password: "password" }
    return &user
}

func random(n int) string {
    rand.Seed(time.Now().UTC().UnixNano())
    return strconv.Itoa(rand.Intn(n))
}

func TestUserSave(t *testing.T) {
    user := RandomUser()
    errorOnSave := user.Save()
    assert.Nil(t, errorOnSave)
    assert.True(t, user.ID > 0)
    assert.NotEqual(t, user.Password, "password")
    assert.Equal(t, len(user.Password), 60)
}

func TestErrorUserSave(t *testing.T) {
    user := RandomUser()
    errorOnSave := user.Save()
    assert.Nil(t, errorOnSave)
    assert.True(t, user.ID > 0)
    // duplicate email
    errorOnSave = user.Save()
    assert.NotNil(t, errorOnSave)
}

func TestUserDestroy(t *testing.T) {
    user := RandomUser()
    errorOnSave := user.Save()
    assert.Nil(t, errorOnSave)
    assert.True(t, user.ID > 0)
    err := user.Destroy()
    assert.Nil(t, err)
}

func TestErrorUserDestroy(t *testing.T) {
    user := User {ID: 0, Email: "unknown@testing.com", Password: "password" }
    err := user.Destroy()
    assert.NotNil(t, err)
}

func TestUserValid(t *testing.T) {
    user := User{}
    assert.NotNil(t, user.Valid())
    user.Email = "missing@password.com"
    assert.NotNil(t, user.Valid())
    user.Password = "herewego"
    assert.Nil(t, user.Valid())
}
