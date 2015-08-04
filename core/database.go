package core

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

// DatabaseConnection connect to the database
func DatabaseConnection() *sql.DB {
    db, err := sql.Open("postgres", "user=app dbname=go_auth sslmode=disable password=famousdev")
    if err != nil {
      log.Fatal(err)
    }

    return db
}

// ExecuteQuery -
func ExecuteQuery(query string, parameters []interface {}) (sql.Result, error) {
    db := DatabaseConnection()

    return db.Exec(query, parameters ...)
}

