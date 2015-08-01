package main

import (
    "log"
    "net/http"
)

func main() {
    log.Fatal(http.ListenAndServe(":8080", LaunchRouter()))

    // db := DatabaseConnection()

    // email := RandomEmail()

    // CreateUser(db, email, RandomPassword())

    // fmt.Println(GetUserByEmail(db, email))

    // defer db.Close()


    // Comparing the password with the hash
    // err = bcrypt.CompareHashAndPassword(hashedPassword, password)
    // fmt.Println(err) // nil means it is a match

}
