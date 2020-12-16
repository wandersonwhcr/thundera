package handlers

import (
    "encoding/json"
    "net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
    home := map[string]string{
        "name": "accounts",
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(home)
}
