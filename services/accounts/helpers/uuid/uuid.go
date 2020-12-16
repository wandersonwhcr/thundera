package uuid

import (
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
