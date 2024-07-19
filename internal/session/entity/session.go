package entity

import "time"

type Session struct {
	Token     string
	CreatedAt time.Time
}
