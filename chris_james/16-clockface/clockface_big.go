package clockface

import (
	"math"
	"math/big"
	"time"
)

type BigPoint struct {
	X, Y *big.Float
}

func bigFloatSecondsInRadian(t time.Time) *big.Float {
	seconds := big.NewRat(int64(t.Second()), 1)
	secondsInMinute := big.NewRat(60, 1)
	twoPi := new(big.Rat).SetFloat64(2 * math.Pi)

	angle := new(big.Rat).Quo(seconds, secondsInMinute)
	angle.Mul(angle, twoPi)

	angleFloat := new(big.Float).SetRat(angle)
	return angleFloat
}

func bigFloatSecondHandPoint(t time.Time) BigPoint {
	angle := bigFloatSecondsInRadian(t)
	sin := bigFloatSin(angle)
	cos := bigFloatCos(angle)
	return BigPoint{X: sin, Y: cos}
}

func bigFloatSin(angle *big.Float) *big.Float {
	sin := new(big.Float).Set(angle)
	term := new(big.Float).Set(angle)
	angleSquared := new(big.Float).Mul(angle, angle)
	factorial := new(big.Float).SetInt64(1)
	sign := new(big.Float).SetInt64(-1)

	for i := 3; i < 50; i += 2 {
		factorial.Mul(factorial, big.NewFloat(float64(i*(i-1))))
		term.Mul(term, angleSquared)
		term.Quo(term, factorial)
		term.Mul(term, sign)
		sin.Add(sin, term)
		sign.Neg(sign)
	}

	return sin
}

func bigFloatCos(angle *big.Float) *big.Float {
	cos := new(big.Float).SetInt64(1)
	term := new(big.Float).SetInt64(1)
	angleSquared := new(big.Float).Mul(angle, angle)
	factorial := new(big.Float).SetInt64(1)
	sign := new(big.Float).SetInt64(-1)

	for i := 2; i < 50; i += 2 {
		factorial.Mul(factorial, big.NewFloat(float64(i*(i-1))))
		term.Mul(term, angleSquared)
		term.Quo(term, factorial)
		term.Mul(term, sign)
		cos.Add(cos, term)
		sign.Neg(sign)
	}

	return cos
}
