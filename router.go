package main

import (
    "net/http"

    "github.com/gabz75/go-auth-api/controllers"
    "github.com/gabz75/go-auth-api/core"

    "github.com/gorilla/mux"
)

// Route struct
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes collection of route
type Routes []Route

var routes = Routes{
    Route{
        "PostUser",
        "POST",
        "/v1/users",
        controllers.PostUser,
    },
    Route{
        "PostSession",
        "POST",
        "/v1/sessions",
        controllers.PostSession,
    },
    Route{
        "GetSessions",
        "GET",
        "/v1/sessions",
        controllers.GetSessions,
    },
    Route{
        "DestroySession",
        "DELETE",
        "/v1/sessions",
        controllers.DestroySession,
    },
}

// LaunchRouter enable routes
func LaunchRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = core.Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
}
