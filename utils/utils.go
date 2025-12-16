package utils

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func AbsDiff[T Number](i, j T) int {
	fi, fj := float64(i), float64(j)

	return int(math.Abs(fi - fj))
}
