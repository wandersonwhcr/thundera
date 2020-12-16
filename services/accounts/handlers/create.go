package handlers

import (
    "encoding/json"
    "net/http"
)

func Create(writer http.ResponseWriter, request *http.Request) {
    identifier := map[string]string{
        "_id": "399a1586-d2a9-4f0d-abcd-9e5db0834e2f",
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusCreated)
    json.NewEncoder(writer).Encode(identifier)
}
