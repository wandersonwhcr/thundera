package main

import (
    "encoding/json"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func Home(writer http.ResponseWriter, request *http.Request) {
    home := map[string]string{
        "name": "accounts",
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(home)
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", Home).Methods("GET")

    logger := handlers.LoggingHandler(os.Stdout, router)

    http.ListenAndServe(":" + os.Getenv("PORT"), logger)
}
