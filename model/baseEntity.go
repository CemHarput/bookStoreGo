package model

import (
	"time"
)

type CommonFields struct {
	ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}