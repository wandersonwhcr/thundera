package handlers

import (
    "accounts/helpers/uuid"
    "context"
    "encoding/json"
    "go.mongodb.org/mongo-driver/mongo"
    "net/http"
)

type Account struct {
    Id uuid.UUID `json:"_id" bson:"_id"`
}

func Create(writer http.ResponseWriter, request *http.Request) {
    database := request.Context().Value("database").(*mongo.Client)

    account := Account{uuid.New()}

    database.Database("accounts").Collection("accounts").InsertOne(context.TODO(), account)

    identifier := map[string]uuid.UUID{
        "_id": account.Id,
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusCreated)
    json.NewEncoder(writer).Encode(identifier)
}
