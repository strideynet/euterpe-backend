package domain

import (
	"math"
	"time"
)

type Teapot struct {
	ID         string    `db:"id"`
	Name       string    `db:"name"`
	Radius     float64   `db:"radius"`
	Height     float64   `db:"height"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

// Volume returns the volume of the cylindrical teapot in cubic metres
func (t Teapot) Volume() float64 {
	// Volume = π(r^2)h
	return math.Pi * math.Pow(t.Radius, 2) * t.Height
}
