package vector

import (
	"fmt"
)

type Vector struct {X, Y float64}

func Create(x, y float64) Vector{
	v := Vector{x, y}
	return v
}

// Vectors logic
func (a *Vector) Add(b Vector) {
	(*a).X += b.X
	(*a).Y += b.Y
}

func (a *Vector) Mult(i float64) {
	(*a).X *= i
	(*a).Y *= i
}

func (a *Vector) Div(i float64) {
	(*a).X /= i
	(*a).Y /= i
}

func Printlen(a *Vector) {
	fmt.Println((*a).X, (*a).Y)
}