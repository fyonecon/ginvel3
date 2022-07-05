package kits

import (
	"github.com/google/uuid"
)

func UUID() string {
	id := uuid.New().String()
	return id
}
