package public

import (
	"math/big"
)

func AddFloat64(x, y float64) float64 {
	bx := big.NewFloat(x)
	by := big.NewFloat(y)

	res := big.NewFloat(0)
	res.Add(bx, by)

	f, _ := res.Float64()
	return f
}

func SubFloat64(x, y float64) float64 {
	bx := big.NewFloat(x)
	by := big.NewFloat(y)

	res := big.NewFloat(0)
	res.Sub(bx, by)

	f, _ := res.Float64()
	return f
}

func MulFloat64(x, y float64) float64 {
	bx := big.NewFloat(x)
	by := big.NewFloat(y)

	res := big.NewFloat(0)
	res.Mul(bx, by)

	f, _ := res.Float64()
	return f
}

func DivFloat64(x, y float64) float64 {
	bx := big.NewFloat(x)
	by := big.NewFloat(y)

	res := big.NewFloat(0)
	res.Quo(bx, by)

	f, _ := res.Float64()
	return f
}
