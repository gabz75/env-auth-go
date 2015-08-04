package core

// Persistable -
type Persistable interface {
  Schema() Mappings
  Table() string
}
