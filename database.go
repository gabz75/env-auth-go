package main

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

// DatabaseConnection connect to the database
func DatabaseConnection() *sql.DB {
    db, err := sql.Open("postgres", "user=app dbname=go_auth sslmode=verify-full password=famousdev")
    if err != nil {
      log.Fatal(err)
    }

    return db
}
