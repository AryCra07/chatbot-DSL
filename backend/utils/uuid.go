package utils

import (
	"github.com/google/uuid"
)

func GetUUID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
