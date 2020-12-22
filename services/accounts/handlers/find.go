package handlers

import (
    "accounts/entities"
    "context"
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "net/http"
)

func Find(writer http.ResponseWriter, request *http.Request) {
    database := request.Context().Value("database").(*mongo.Client)

    var accounts []*entities.Account

    cursor, _ := database.Database("accounts").Collection("accounts").Find(context.TODO(), bson.M{})

    for cursor.Next(context.TODO()) {
        var account entities.Account
        cursor.Decode(&account)
        accounts = append(accounts, &account)
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(accounts)
}
