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

var client *mongo.Client

func main() {
    client, _ = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
    ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
    client.Connect(ctx)

    router := mux.NewRouter()

    router.Use(DatabaseMiddleware)

    router.HandleFunc("/", handlers.Home).Methods("GET")
    router.HandleFunc("/v1/accounts", handlers.Create).Methods("POST")

    logger := middlewares.LoggingHandler(os.Stdout, router)

    http.ListenAndServe(":" + os.Getenv("PORT"), logger)
}

func DatabaseMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        request = request.WithContext(context.WithValue(request.Context(), "database", client))
        next.ServeHTTP(writer, request)
    })
}
