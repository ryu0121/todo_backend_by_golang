package domain

import "time"

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
