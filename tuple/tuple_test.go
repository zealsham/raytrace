package tuple

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorCreation(t *testing.T) {
	data := Vector(1, 2, 4)
	expected := Tuple{1, 2, 4, 0}

	assert.Equal(t, data, expected, "should be equal")
}

func TestPointCreation(t *testing.T) {
	data := Point(1, 3, 6)
	expected := Tuple{1, 3, 6, 1.0}

	assert.Equal(t, data, expected, "should be equal")
}
func TestPoint(t *testing.T) {
	point := Tuple{4.3, -4.2, 3.1, 1.0}

	assert := assert.New(t)

	assert.Equal(point.X, 4.3, "they should be equal")
	assert.Equal(point.Y, -4.2, "they should be equal")
	assert.Equal(point.Z, 3.1, "they should be equal")
	assert.Equal(point.W, 1.0, "they should be equal")
	assert.Equal(IsPoint(point), true, "should return true if it's a point")
	assert.Equal(IsVector(point), false, "a point isnt a vector")

}

func TestVector(t *testing.T) {
	vector := Tuple{4.3, -4.2, 3.1, 0.0}
	assert := assert.New(t)
	assert.Equal(vector.X, 4.3, "should be equal")
	assert.Equal(vector.Y, -4.2, "should be eqaul")
	assert.Equal(vector.Z, 3.1, "should be equal")
	assert.Equal(vector.W, 0.0, "should be equal")
	assert.Equal(IsPoint(vector), false, "a vector isnt a point")
	assert.Equal(IsVector(vector), true, "should be true if it's a vector")
}

func TestEqualTuple(t *testing.T) {
	tup1 := Tuple{1.2, 3.2, 2.1, 0.0}
	tup2 := Tuple{1.2, 3.2, 2.1, 0.0}
	tup3 := Tuple{1.2, 3.2, 2.1, 1.0}
	assert := assert.New(t)
	assert.Equal(EqualTuple(tup1, tup2), true, "the two  should be equal")
	assert.Equal(EqualTuple(tup2, tup3), false, "they shoudn't be equal in this case")

}

func TestAddTuple(t *testing.T) {
	tup1 := Tuple{3, -2, 5, 1}
	tup2 := Tuple{-2, 3, 1, 0}
	expected := Tuple{1, 1, 6, 1}

	assert := assert.New(t)
	assert.Equal(AddTuple(tup1, tup2), expected, "should be equal")
}

func TestSubtractTuple(t *testing.T) {
	tup1 := Tuple{3, 2, 1, 1}
	tup2 := Tuple{5, 6, 7, 1}
	expected := Tuple{-2, -4, -6, 0}
	testCaseVector := Tuple{2, 4, 5, 0}

	assert := assert.New(t)
	assert.Equal(SubTuple(tup1, tup2), expected, "should be equal")
	assert.Equal(IsVector(SubTuple(tup1, tup2)), true, "subtracting point from point should return a vector")
	assert.Equal(IsPoint(SubTuple(tup1, expected)), true, "subtracting a vector from point should return a point")
	assert.Equal(IsVector(SubTuple(expected, testCaseVector)), true, "subtracting two vector should return a new vector")
}

func TestNegateTuple(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	expected := Tuple{-1, 2, -3, 4}

	assert.Equal(t, NegTuple(tup), expected, "should be equal to the expected")
}

func TestMultTuple(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	scalar := 3.5

	//a * tuple should  be
	expected := Tuple{3.5, -7, 10.5, -14}

	assert.Equal(t, MultTuple(scalar, tup), expected, "should be equal")
}

func TestDivTuple(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	scalar := 2.0
	expected := Tuple{0.5, -1, 1.5, -2}

	assert.Equal(t, DivTuple(scalar, tup), expected, "should be equal")
}
func TestMagniVector(t *testing.T) {
	testdata := []Tuple{{0, 1, 0, 0}, {0, 0, 1, 0}, {1, 2, 3, 0}, {-1, -2, -3, 0}}
	expected := []float64{1, 1, math.Sqrt(14), math.Sqrt(14)}

	assert := assert.New(t)

	for i, v := range testdata {
		assert.Equal(Equal(MagnitudeTuple(v), expected[i]), true, "should be equal to")
	}
}

func TestNormVector(t *testing.T) {

	//my insticts tells me this can be done in a better way, currently just have a one day experience using
	//tesitfy library, maybe as i gain mor experience i'll figure a better way out
	testdata := []Tuple{{4, 0, 0, 0}, {1, 2, 3, 0}}
	expected := []Tuple{{1, 0, 0, 0}, {0.26726, 0.53452, 0.80178, 0}}
	magcalls := []Tuple{NormalizeTuple(testdata[0]), NormalizeTuple(testdata[1])}

	assert := assert.New(t)

	for i, v := range expected {
		assert.Equal(Equal(v.X, magcalls[i].X), true, "should be equal")
		assert.Equal(Equal(v.Y, magcalls[i].Y), true, "should be equal")
		assert.Equal(Equal(v.Z, magcalls[i].Z), true, "should be equal")
		assert.Equal(Equal(v.W, magcalls[i].W), true, "should be equal")
		assert.Equal(Equal(MagnitudeTuple(magcalls[i]), 1.0), true, "should be equal")
	}

}
func TestDot(t *testing.T) {
	testdata1 := Tuple{1, 2, 3, 0}
	testdata2 := Tuple{2, 3, 4, 0}
	expected := 20.0

	assert.Equal(t, DotProduct(testdata1, testdata2), expected, "should return the correct float")
}
func TestCross(t *testing.T) {
	a := Tuple{1, 2, 3, 0}
	b := Tuple{2, 3, 4, 0}
	crossAB := Tuple{-1, 2, -1, 0}
	crossBA := Tuple{1, -2, 1, 0}
	assert := assert.New(t)
	assert.Equal(CrossProduct(a, b), crossAB, "should be equal")
	assert.Equal(CrossProduct(b, a), crossBA, "should be equal")

}
