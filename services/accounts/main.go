package main

import (
    "accounts/handlers"
    "github.com/gorilla/mux"
    "net/http"
    "os"
    middlewares "github.com/gorilla/handlers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.Home).Methods("GET")

    logger := middlewares.LoggingHandler(os.Stdout, router)

    http.ListenAndServe(":" + os.Getenv("PORT"), logger)
}
