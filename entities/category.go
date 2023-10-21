package entities

import "time"

type Category struct {
	Id         uint
	Name       string
	Created_at time.Time
	Updated_at time.Time
}