package core

import (
    "reflect"
    "bytes"
    "strconv"
    "database/sql"
)

// InsertQuery - execute an insert for a Persistable object
func InsertQuery(p Persistable) (sql.Result, error) {
    insert := InsertTextQuery(p)
    parameters := InsertParametersQuery(p)

    return ExecuteQuery(insert, parameters)
}

// DeleteQuery - execute a delete for a Persistable object
func DeleteQuery(p Persistable) (sql.Result, error) {
    delete := DeleteTextQuery(p)
    parameters := []interface{} { CastField(p, "ID", reflect.Int64) }

    return ExecuteQuery(delete, parameters)
}

// DeleteTextQuery - dynammically generate delete query
func DeleteTextQuery(p Persistable) string {
    var query bytes.Buffer

    query.WriteString("DELETE FROM ")
    query.WriteString(p.Table())
    query.WriteString(" WHERE id = $1")

    return query.String()
}

// InsertTextQuery - dynammically generate insert query
func InsertTextQuery(p Persistable) string {
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

// InsertParametersQuery - dynammically generate parameters for insert
func InsertParametersQuery(p Persistable) []interface{} {
    parameters := make([]interface{}, len(p.Schema()))

    for index, mapping := range p.Schema() {
       parameters[index] = CastField(p, mapping.Name, mapping.Type)
    }

    return parameters
}

// CastField - cast a field into the right type.
func CastField(p Persistable, name string, kind reflect.Kind) interface{} {
    // @TODO improve support of other Type (float, string, etc...)
    r := reflect.ValueOf(p)
    f := reflect.Indirect(r).FieldByName(name)
    if kind == reflect.Int || kind == reflect.Int64 {
        return f.Int()
    }

    return f.String()
}
