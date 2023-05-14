package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func GenerateID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {

	r, err := uuid.Parse(s)

	return ID(r), err
}
