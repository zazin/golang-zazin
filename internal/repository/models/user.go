package models

import "time"

type Model struct {
	ID        uint
	Name      string
	CreatedAt time.Time
}
