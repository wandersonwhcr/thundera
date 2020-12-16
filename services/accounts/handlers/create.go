package handlers

import (
    "accounts/helpers/uuid"
    "encoding/json"
    "net/http"
)

func Create(writer http.ResponseWriter, request *http.Request) {
    _id := uuid.New()

    identifier := map[string]uuid.UUID{
        "_id": _id,
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusCreated)
    json.NewEncoder(writer).Encode(identifier)
}
