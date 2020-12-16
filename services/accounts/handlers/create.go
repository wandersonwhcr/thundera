package handlers

import (
    "accounts/helpers/uuid"
    "encoding/json"
    "net/http"
)

func Create(writer http.ResponseWriter, request *http.Request) {
    identifier := map[string]string{
        "_id": uuid.New().String(),
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusCreated)
    json.NewEncoder(writer).Encode(identifier)
}
