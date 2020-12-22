package uuid

import (
    "encoding/json"
    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/bsontype"
    "go.mongodb.org/mongo-driver/bson/primitive"
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

func (u UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
    return bson.MarshalValue(primitive.Binary{bsontype.BinaryUUID, u.Data})
}

func (u *UUID) UnmarshalBSONValue(t bsontype.Type, data []byte) (error) {
    _, Data := bson.RawValue{Type: t, Value: data}.Binary()
    u.Data = Data
    return nil
}
