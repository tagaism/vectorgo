package vector_test
import (
	"testing"
	"github.com/tagaism/vectorgo"
)

func TestCreate(test *testing.T) {
	var x = float64(1.1)
	var y = float64(2.2)
	vec := vector.Create(x, y)

	if vec.X != x || vec.Y != y {
		test.Fail()
	}
}

func TestCopy(test *testing.T) {
	vec := vector.Create(1.1, 1.1)

	copied_vec := vec.Copy()

	if vec.X != copied_vec.X || vec.Y != copied_vec.Y {
		test.Fail()
	}

	if &vec == &copied_vec {
		test.Errorf("Vectors' pointers must not be equal.")
	}
}

func TestAdd(test *testing.T) {
	vec1 := vector.Vector{1.0, 1.0}
	vec2 := vector.Vector{2.0, 2.0}

	vec1.Add(vec2)
	if vec1.X != 3.0 || vec1.Y != 3.0 {
		test.Fail()
	}
}

func TestSub(test *testing.T) {
	vec1 := vector.Vector{3.0, 3.0}
	vec2 := vector.Vector{2.0, 2.0}

	vec1.Sub(vec2)
	if vec1.X != 1.0 || vec1.Y != 1.0 {
		test.Fail()
	}
	
}

func TestMult(test *testing.T) {
	vec := vector.Vector{3, 4}
	var m = float64(2)

	vec.Mult(m)
	if vec.X != 6 || vec.Y != 8 {
		test.Fail()
	}
}

func TestDiv(test *testing.T) {
	vec := vector.Vector{3, 4}
	var m = float64(2)

	vec.Div(m)
	if vec.X != 1.5 || vec.Y != 2 {
		test.Fail()
	}
}

func TestMag(test *testing.T) {
	vec := vector.Vector{3, 4}

	magnitude := vec.Mag()
	if magnitude != 5.0 {
		test.Fail()
	}
}

func TestNormal(test *testing.T) {
	vec := vector.Vector{3, 4}

	magnitude := vec.Mag()
	normalized_x := vec.X / magnitude
	normalized_y := vec.Y / magnitude

	vec.Normal()
	if vec.X != normalized_x || vec.Y != normalized_y {
		test.Fail()
	}
}

func TestLimit(test *testing.T) {
	var x, y = float64(3), float64(4)
	vec := vector.Create(x, y)
	max_val := 2.0

	magnitude := vec.Mag()
	delim := magnitude / max_val
	limited_x := x / delim
	limited_y := y / delim

	vec.Limit(2)
	if vec.X != limited_x || vec.Y != limited_y {
		test.Fail()
	}
}
