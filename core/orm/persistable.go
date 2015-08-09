package orm

import (
    "reflect"
)

// Mapping - describe the internal db structure of a model
type Mapping struct {
    Name string
    DBName string
    Type reflect.Kind
}

// Mappings - Collection of Mapping
type Mappings []Mapping

// Persistable -
type Persistable interface {
  Schema() Mappings
  Table() string
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
