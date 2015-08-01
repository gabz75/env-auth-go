package main

import (
    "net/http"

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
        "/users",
        PostUser,
    },
    Route{
        "GetUser",
        "GET",
        "/users",
        GetUser,
    },
    Route{
        "PostSession",
        "POST",
        "/sessions",
        PostSession,
    },
}

// LaunchRouter enable routes
func LaunchRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
}
