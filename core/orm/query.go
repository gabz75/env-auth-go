package orm

import (
    "reflect"
    "bytes"
    "strconv"
    "database/sql"
)

// Delete - Delete a persistable object from DB
func Delete(p Persistable) (sql.Result, error) {
    return ExecuteQuery(deleteQuery(p), []interface{} { CastField(p, "ID", reflect.Int64) })
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
