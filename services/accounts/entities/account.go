package entities

import (
    "accounts/helpers/uuid"
)

type Account struct {
    Id uuid.UUID `json:"_id" bson:"_id"`
}
