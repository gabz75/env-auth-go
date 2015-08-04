package core

import (
    "reflect"
)

// Mapping -
type Mapping struct {
    Name string
    DBName string
    Type reflect.Kind
}

// Mappings -
type Mappings []Mapping
