package entity

import "time"

type User struct {
	Id        string
	Name      string
	CreatedAt time.Time
}
