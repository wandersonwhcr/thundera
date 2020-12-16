package uuid

import (
    "encoding/json"
    "github.com/google/uuid"
)

type UUID struct {
    Data []byte
}

func New() (UUID) {
    data, _ := uuid.New().MarshalBinary()

    return UUID{data}
}

func (u UUID) String() (string) {
    value, _ := uuid.FromBytes(u.Data)

    return value.String()
}

func (u UUID) MarshalJSON() ([]byte, error) {
    value, _ := uuid.FromBytes(u.Data)

    return json.Marshal(value)
}
