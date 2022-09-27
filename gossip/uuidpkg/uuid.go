package uuidpkg 

import (
	"github.com/google/uuid"
)

type UUID []byte 

func NewRandom() (UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err 
	}

	return UUID(id[:]), nil 
}