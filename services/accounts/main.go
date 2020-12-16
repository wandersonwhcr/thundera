package main

import (
    "accounts/handlers"
    "context"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "net/http"
    "os"
    "time"
    middlewares "github.com/gorilla/handlers"
)

func main() {
    client, _ := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
    ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
    client.Connect(ctx)

    router := mux.NewRouter()

    router.HandleFunc("/", handlers.Home).Methods("GET")
    router.HandleFunc("/v1/accounts", handlers.Create).Methods("POST")

    logger := middlewares.LoggingHandler(os.Stdout, router)

    http.ListenAndServe(":" + os.Getenv("PORT"), logger)
}
