package main

import (
  "time"
  "math/rand"
)

// RandomPassword generate random password
func RandomPassword() string {
    return RandomString(40)
}

// RandomEmail generate random email
func RandomEmail() string {
    return RandomString(4) + "@famo.us"
}

// RandomString generate random string
func RandomString(n int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    rand.Seed(time.Now().UTC().UnixNano())

    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }

    return string(b)
}
