package models

import (
  "reflect"

  "github.com/gabz75/go-auth-api/core"
  "github.com/gabz75/go-auth-api/core/orm"
)

// Session - hold tokens to authenticate the User
type Session struct {
    ID int64 `json:"id"`
    UserID int64 `json:"-"`
    Token string `json:"token"`
}

// Schema - mapping between model and DB
func (session *Session) Schema() orm.Mappings {
    return orm.Mappings{
        orm.Mapping{
            "UserID",
            "user_id",
            reflect.Int,
        },
        orm.Mapping {
            "Token",
            "token",
            reflect.String,
        },
    }
}

// Table - DB table name
func (session *Session) Table() string {
    return "sessions"
}

// Save - insert session in DB
func (session *Session) Save() error {
    session.Token = core.GenerateToken()

    if _, err := orm.Insert(session); err != nil {
      return err
    }

    return nil
}

// Destroy - delete session from DB
func (session *Session) Destroy() error {
    if _, err := orm.Delete(session); err != nil {
        return err
    }

    return nil
}
