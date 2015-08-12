package orm

import (
    "errors"
    "reflect"
    "bytes"
    "strconv"
    "database/sql"
)

// Delete - Delete a persistable object from DB
func Delete(p Persistable) (sql.Result, error) {
    result, error := ExecuteQuery(deleteQuery(p), []interface{} { CastField(p, "ID", reflect.Int64) })

    count, _ := result.RowsAffected()

    if count == int64(0) {
        return nil, errors.New("no rows affected")
    }

    return result, error
}

func deleteQuery(p Persistable) string {
    var query bytes.Buffer

    query.WriteString("DELETE FROM ")
    query.WriteString(p.Table())
    query.WriteString(" WHERE id = $1")

    return query.String()
}

// Insert - Insert a Persistable object in DB
func Insert(p Persistable) (sql.Result, error) {
    return ExecuteQuery(insertQuery(p), parameters(p))
}

// LastInsertedID - Return the last Id
func LastInsertedID(p Persistable) (int64, error) {
    db := DatabaseConnection()

    mapping := p.Schema()[0]
    firstAttribute := CastField(p, mapping.Name, mapping.Type)
    query := "SELECT id FROM " + p.Table() + " WHERE " + mapping.Name + " = $1"
    rows, error := db.Query(query, firstAttribute)

    if error != nil {
        return 0, error
    }

    defer rows.Close()

    var id int64

    for rows.Next() {
        error = rows.Scan(&id)
        if error != nil {
            return 0, error
        }
    }

    return id, nil
}

func insertQuery(p Persistable) string {
    var query bytes.Buffer
    var columns bytes.Buffer
    var values bytes.Buffer

    query.WriteString("INSERT INTO ")
    query.WriteString(p.Table())
    query.WriteString(" ")

    columns.WriteString("(")
    values.WriteString("(")

    lastIndex := len(p.Schema())

    count := 0
    for _, mapping := range p.Schema() {
        columns.WriteString(mapping.DBName)
        values.WriteString("$")
        values.WriteString(strconv.Itoa(count + 1))
        if count < lastIndex - 1{
            columns.WriteString(", ")
            values.WriteString(", ")
        }
        count++
    }

    columns.WriteString(")")
    values.WriteString(")")

    query.WriteString(columns.String())
    query.WriteString(" VALUES ")
    query.WriteString(values.String())

    return query.String()
}

func parameters(p Persistable) []interface{} {
    parameters := make([]interface{}, len(p.Schema()))

    for index, mapping := range p.Schema() {
       parameters[index] = CastField(p, mapping.Name, mapping.Type)
    }

    return parameters
}
