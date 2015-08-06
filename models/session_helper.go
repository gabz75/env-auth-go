package models

import (
    "errors"

    "github.com/gabz75/auth-api/core"
)

// GetSessions -
func GetSessions(user User) []Session {
    db := core.DatabaseConnection()

    rows, err := db.Query("SELECT * FROM sessions WHERE user_id = $1", user.ID)

    defer rows.Close()

    var sessions []Session
    var id int64
    var userID int64
    var token string

    for rows.Next() {
        err = rows.Scan(&id, &userID, &token)
        if err != nil {
            panic(err)
        }
        sessions = append(sessions, Session{ID: id, UserID: userID, Token: token})
    }

    return sessions
}

// GetSession -
func GetSession(user User, token string) (*Session, error) {
    db := core.DatabaseConnection()

    rows, err := db.Query("SELECT * FROM sessions WHERE user_id = $1 AND token = $2", user.ID, token)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    var session Session
    var id int64
    var userID int64
    var sessionToken string

    for rows.Next() {
        err = rows.Scan(&id, &userID, &sessionToken)
        if err != nil {
            panic(err)
        }
        session = Session{ID: id, UserID: userID, Token: sessionToken}
    }

    if session.ID == 0 {
        return nil, errors.New("no session found")
    }

    return &session, nil
}
