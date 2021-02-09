package domain

import (
	"math"
	"time"
)

type Teapot struct {
	ID         string
	Name       string
	Radius     float64
	Height     float64
	CreateTime time.Time
	UpdateTime time.Time
}

// Volume returns the volume of the cylindrical teapot in cubic metres
func (t Teapot) Volume() float64 {
	// Volume = π(r^2)h
	return math.Pi * math.Pow(t.Radius, 2) * t.Height
}
