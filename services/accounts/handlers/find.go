package handlers

import (
    "encoding/json"
    "net/http"
)

func Find(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode([]string{})
}
