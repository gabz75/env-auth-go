package orm

import (
    "database/sql"
    "log"

    "github.com/gabz75/go-auth-api/core"

    _ "github.com/lib/pq"
)

// DatabaseConnection connect to the database
func DatabaseConnection() *sql.DB {
    config := core.GetConfig()
    db, err := sql.Open("postgres", "user=" + config.Db.Username + " dbname=" + config.Db.Database + " sslmode=disable password=" + config.Db.Password)
    if err != nil {
      log.Fatal(err)
    }

    return db
}

// ExecuteQuery - initialize db connection and execute a query
func ExecuteQuery(query string, parameters []interface {}) (sql.Result, error) {
    db := DatabaseConnection()

    return db.Exec(query, parameters ...)
}
