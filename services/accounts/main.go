package main

import (
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func Home(writer http.ResponseWriter, request *http.Request) {
    writer.WriteHeader(http.StatusOK)
    writer.Write([]byte("Accounts"))
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", Home).Methods("GET")

    logger := handlers.LoggingHandler(os.Stdout, router)

    http.ListenAndServe(":" + os.Getenv("PORT"), logger)
}
