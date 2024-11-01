package gouant

import "math"

type Greeks struct {
	Delta float64
	Gamma float64
	Theta float64
	Rho   float64
	Vega  float64
}

// CND Cumulative Normal Distribution
func CND(x float64) float64 {
	return 0.5 * math.Erfc(-x/math.Sqrt2)
}

// PD Probability Density
func PD(x float64) float64 {
	return (1 / math.Sqrt(2*math.Pi)) * math.Exp(-0.5*math.Pow(x, 2))
}
