package core

import (
    "reflect"
)

// Mapping - describe the internal db structure of a model
type Mapping struct {
    Name string
    DBName string
    Type reflect.Kind
}

// Mappings - slice of Mapping
type Mappings []Mapping
