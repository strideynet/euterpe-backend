package domain

import (
	"time"
)

type Teapot struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Radius      float64   `db:"radius"`
	Height      float64   `db:"height"`
	CreateTime  time.Time `db:"create_time"`
	UpdateTime  time.Time `db:"update_time"`
	Description string    `db:"description"`
}
