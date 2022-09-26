package helpers

import "github.com/google/uuid"

func GenerateId() string {
	uid := uuid.NewString()

	return uid
}
