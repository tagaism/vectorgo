package vector

import (
	"fmt"
	"math"
)

type Vector struct {X, Y float64}

func Create(x, y float64) Vector{
	v := Vector{x, y}
	return v
}

func (a *Vector) Copy() Vector {
	return Vector{a.X, a.Y}
}

// Vectors logic
func (a *Vector) Add(b Vector) {
	(*a).X += b.X
	(*a).Y += b.Y
}

func (a *Vector) Sub(b Vector) {
	(*a).X -= b.X
	(*a).Y -= b.Y
}

func Sub(a, b Vector) Vector {
	c := Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
	return c
}

func (a *Vector) Mult(i float64) {
	(*a).X *= i
	(*a).Y *= i
}

func (a *Vector) Div(i float64) {
	(*a).X /= i
	(*a).Y /= i
}

// Calculate magnitude (length) of the vector
func (a *Vector) Mag() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// Normilize (unify) vector
func (a *Vector) Normal() {
	mag := (*a).Mag()
	(*a).X /= mag
	(*a).Y /= mag
}

func (a *Vector) Limit(max float64) {
	mag := (*a).Mag()
	delim := mag / max
	(*a).X /= delim
	(*a).Y /= delim
}

func PrintVec(a *Vector) {
	fmt.Println(*a)
}
