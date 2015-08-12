package orm

import (
    "testing"
    "reflect"
    "time"
    "math/rand"
    "strconv"

    "github.com/stretchr/testify/assert"
)

// - * - Helpers & Mock Objects - * -

type MockObject struct {
    ID int64 `json:"id"`
    Email string `json:"email"`
    Password string `json:"password"`
}

// Schema - mapping between model and DB
func (mockObject MockObject) Schema() Mappings {
    return Mappings{
        Mapping{
            "Email",
            "email",
            reflect.String,
        },
        Mapping {
            "Password",
            "password",
            reflect.String,
        },
    }
}

func RandomMockObject() *MockObject {
    user := MockObject{ Email: "mock-" + random(999999999999999) + "@testing.com", Password: "password" }
    return &user
}

func random(n int) string {
    rand.Seed(time.Now().UTC().UnixNano())
    return strconv.Itoa(rand.Intn(n))
}

// Table - DB table name
func (mockObject MockObject) Table() string {
    return "users"
}

// - * - Test Suite - * -

func TestDeleteQuery(t *testing.T) {
    mock := RandomMockObject()
    assert.Equal(t, deleteQuery(mock), "DELETE FROM users WHERE id = $1")
}

func TestInsertQuery(t *testing.T) {
    mock := RandomMockObject()
    assert.Equal(t, insertQuery(mock), "INSERT INTO users (email, password) VALUES ($1, $2)")
}

func TestInsert(t *testing.T) {
    mock := RandomMockObject()
    result, error := Insert(mock)
    assert.Nil(t, error)
    assert.NotNil(t, result)
    r, _ := result.RowsAffected()
    assert.Equal(t, r, int64(1))
}

func TestLastInsertedID(t *testing.T) {
    mock := RandomMockObject(); Insert(mock)
    id, error := LastInsertedID(mock)
    assert.Nil(t, error)
    assert.True(t, id > 0)
}

func TestDestroy(t *testing.T) {
    mock := RandomMockObject(); Insert(mock)
    id, _ := LastInsertedID(mock)
    mock.ID = id
    result, error := Delete(mock)
    newID, _ := LastInsertedID(mock)
    assert.Nil(t, error)
    assert.NotNil(t, result)
    r, _ := result.RowsAffected()
    assert.Equal(t, r, int64(1))
    assert.NotEqual(t, id, newID)
}

func TestFailDestroy(t *testing.T) {
    mock := RandomMockObject()
    result, error := Delete(mock)
    assert.Equal(t, mock.ID, int64(0))
    assert.Nil(t, result)
    assert.NotNil(t, error)
    assert.Equal(t, error.Error(), "no rows affected")
}
