package vectorgo

import (
	"math"
)

type Vector struct {
	x float64
	y float64
}

func Create(x, y float64) Vector{
	v := Vector{x, y}
	return v
}