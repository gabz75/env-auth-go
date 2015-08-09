package models

import (
    "testing"
    "time"
    "math/rand"
    "strconv"

    "github.com/stretchr/testify/assert"
)

func RandomUser() *User {
    user := User{ Email: "user-" + random(100) + "@testing.com", Password: "password" }
    return &user
}

func random(n int) string {
    rand.Seed(time.Now().Unix())
    return strconv.Itoa(rand.Intn(n))
}

func TestUserSave(t *testing.T) {
    user := RandomUser()
    errorOnSave := user.Save()
    assert.Nil(t, errorOnSave)
    assert.NotNil(t, user)
    assert.NotEqual(t, user.ID, 0)
}

func TestUserDestroy(t *testing.T) {
    // @TODO this is actually not working
    user := RandomUser()
    err := user.Destroy()
    assert.Nil(t, err)
}
