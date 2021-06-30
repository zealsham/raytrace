package tuple

import "math"

//addresses floating point number comparaison computers have

const EPSILON = 10e-6

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func Point(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0.0}
}

func IsVector(tup Tuple) bool {
	return Equal(tup.W, 0.0)
}

func IsPoint(tup Tuple) bool {
	return Equal(tup.W, 1.0)
}

func EqualTuple(tup1 Tuple, tup2 Tuple) bool {
	return (Equal(tup1.X, tup2.X) && Equal(tup1.Y, tup2.Y) && Equal(tup1.Z, tup2.Z) && Equal(tup1.W, tup2.W))
}
func AddTuple(tup1 Tuple, tup2 Tuple) Tuple {
	//shows the new location if you followd tup2 from tup1
	return Tuple{tup1.X + tup2.X, tup1.Y + tup2.Y, tup1.Z + tup2.Z, tup1.W + tup2.W}
}

func SubTuple(tup1 Tuple, tup2 Tuple) Tuple {
	//basically moving backward by the given  vector
	return Tuple{tup1.X - tup2.X, tup1.Y - tup2.Y, tup1.Z - tup2.Z, tup1.W - tup2.W}
}

func Equal(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func NegTuple(tup Tuple) Tuple {
	// the opposite vector
	return Tuple{-tup.X, -tup.Y, -tup.Z, -tup.W}
}

func MultTuple(scalar float64, tup Tuple) Tuple {
	//determines the point scalar times ahead of the vector
	return Tuple{scalar * tup.X, scalar * tup.Y, scalar * tup.Z, scalar * tup.W}
}

func DivTuple(divider float64, tup Tuple) Tuple {
	return Tuple{tup.X / divider, tup.Y / divider, tup.Z / divider, tup.W / divider}
}

func MagnitudeTuple(tup Tuple) float64 {
	// the total distance covered by a vector
	x := math.Pow(tup.X, 2)
	y := math.Pow(tup.Y, 2)
	z := math.Pow(tup.Z, 2)
	w := math.Pow(tup.W, 2)
	magnitude := math.Sqrt(x + y + z + w)
	return magnitude
}

func NormalizeTuple(tup Tuple) Tuple {
	// transforms a vector into a unit vector
	return Tuple{tup.X / MagnitudeTuple(tup), tup.Y / MagnitudeTuple(tup), tup.Z / MagnitudeTuple(tup), tup.W / MagnitudeTuple(tup)}
}

func DotProduct(tup1 Tuple, tup2 Tuple) float64 {
	result := (tup1.X * tup2.X) + (tup1.Y * tup2.Y) + (tup1.Z * tup2.Z) + (tup1.W * tup2.W)
	return result

}
func CrossProduct(tup1, tup2 Tuple) Tuple {
	x := tup1.Y*tup2.Z - tup1.Z*tup2.Y
	y := tup1.Z*tup2.X - tup1.X*tup2.Z
	z := tup1.X*tup2.Y - tup1.Y*tup2.X

	return Vector(x, y, z)
}
