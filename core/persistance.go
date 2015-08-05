package core

import (
    "reflect"
    "bytes"
    "strconv"
    "database/sql"
)

// CastField -
func CastField(p Persistable, mapping Mapping) interface{} {
    r := reflect.ValueOf(p)
    f := reflect.Indirect(r).FieldByName(mapping.Name)
    if mapping.Type == reflect.Int {
        return f.Int()
    }

    return f.String()
}

// InsertQuery -
func InsertQuery(p Persistable) (sql.Result, error) {
    insert := InsertTextQuery(p)
    parameters := InsertParametersQuery(p)

    return ExecuteQuery(insert, parameters)
}

// InsertTextQuery -
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

// InsertParametersQuery -
func InsertParametersQuery(p Persistable) []interface{} {
    parameters := make([]interface{}, len(p.Schema()))

    for index, mapping := range p.Schema() {
       parameters[index] = CastField(p, mapping)
    }

    return parameters
}
