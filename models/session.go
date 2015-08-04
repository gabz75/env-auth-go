package models

import (
  "reflect"

  "github.com/gabz75/auth-api/core"
)

// Session -
type Session struct {
  UserID int64 `json:"-"`
  Token string `json:"token"`
}

// Schema -
func (session *Session) Schema() core.Mappings {
  return core.Mappings{
    core.Mapping{
      "UserID",
      "user_id",
      reflect.Int,
    },
    core.Mapping {
      "Token",
      "token",
      reflect.String,
    },
  }
}

// Table -
func (session *Session) Table() string {
  return "sessions"
}

// Save -
func (session *Session) Save() {
    session.Token = core.GenerateToken()

    core.InsertQuery(session)
}
