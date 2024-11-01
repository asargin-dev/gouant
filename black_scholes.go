package gouant

import "math"

type BlackScholes struct {
	DerivativePrice float64 // The price of the option.
	UnderlyingPrice float64 // The price of the underlying asset.
	StrikePrice     float64 // The strike price.
	TimeToMaturity  float64 // The annualized time to expiration. Must be positive.
	RiskRate        float64 // The Interest Free Rate.
	Dividend        float64 // The annualized continuous dividend yield.
	IsCall          bool    // For each contract, this should be specified as tru for a call option and false for a put option.
}

func (b *BlackScholes) d1(sigma float64) float64 {
	return (math.Log(b.UnderlyingPrice/b.StrikePrice) + (b.RiskRate-b.Dividend+0.5*math.Pow(sigma, 2))*b.TimeToMaturity) / (sigma * math.Sqrt(b.TimeToMaturity))
}

func (b *BlackScholes) d2(sigma float64) float64 {
	return b.d1(sigma) - sigma*math.Sqrt(b.TimeToMaturity)
}

func (b *BlackScholes) Price(sigma float64) float64 {
	d1 := b.d1(sigma)
	d2 := b.d2(sigma)
	if b.IsCall {
		return b.UnderlyingPrice*math.Exp(-b.Dividend*b.TimeToMaturity)*CND(d1) - b.StrikePrice*math.Exp(-b.RiskRate*b.TimeToMaturity)*CND(d2)
	}
	return b.StrikePrice*math.Exp(-b.RiskRate*b.TimeToMaturity)*CND(-d2) - b.UnderlyingPrice*math.Exp(-b.Dividend*b.TimeToMaturity)*CND(-d1)
}

// IV Implied Volatility calculation using Newton-Raphson method
func (b *BlackScholes) IV() float64 {
	sigma := 0.25 // Initial guess
	tolerance := 1e-8
	maxIterations := 1000

	for i := 0; i < maxIterations; i++ {
		price := b.Price(sigma)
		vega := b.UnderlyingPrice * math.Exp(-b.Dividend*b.TimeToMaturity) * math.Sqrt(b.TimeToMaturity) * CND(b.d1(sigma))

		priceDiff := price - b.DerivativePrice

		if math.Abs(priceDiff) < tolerance {
			return sigma
		}

		sigma -= priceDiff / vega
	}

	return sigma
}

// Delta Calculate Delta
func (b *BlackScholes) Delta(sigma float64) float64 {
	d1 := b.d1(sigma)
	if b.IsCall {
		return math.Exp(-b.Dividend*b.TimeToMaturity) * CND(d1)
	}
	return -math.Exp(-b.Dividend*b.TimeToMaturity) * CND(-d1)
}

// Gamma Calculate Gamma
func (b *BlackScholes) Gamma(sigma float64) float64 {
	d1 := b.d1(sigma)
	return (PD(d1) * math.Exp(-b.Dividend*b.TimeToMaturity)) / (b.UnderlyingPrice * sigma * math.Sqrt(b.TimeToMaturity))
}

// Theta Calculate Theta
func (b *BlackScholes) Theta(sigma float64) float64 {
	d1 := b.d1(sigma)
	d2 := b.d2(sigma)
	if b.IsCall {
		return (-b.UnderlyingPrice*PD(d1)*sigma/(2*math.Sqrt(b.TimeToMaturity)) - (b.RiskRate * b.StrikePrice * math.Exp(-b.RiskRate*b.TimeToMaturity) * CND(d2))) / 365.0
	}
	return (-b.UnderlyingPrice*PD(d1)*sigma*math.Exp(-b.Dividend*b.TimeToMaturity)/(2*math.Sqrt(b.TimeToMaturity)) + (b.RiskRate * b.StrikePrice * math.Exp(-b.RiskRate*b.TimeToMaturity) * CND(-d2)) - (b.Dividend * b.UnderlyingPrice * math.Exp(-b.Dividend*b.TimeToMaturity) * CND(-d1))) / 365.0
}

// Rho Calculate Rho
func (b *BlackScholes) Rho(sigma float64) float64 {
	d2 := b.d2(sigma)
	if b.IsCall {
		return b.StrikePrice * b.TimeToMaturity * math.Exp(-b.RiskRate*b.TimeToMaturity) * CND(d2) / 100.0
	}
	return -b.StrikePrice * b.TimeToMaturity * math.Exp(-b.RiskRate*b.TimeToMaturity) * CND(-d2) / 100.0
}

// Vega Calculate Vega
func (b *BlackScholes) Vega(sigma float64) float64 {
	d1 := b.d1(sigma)
	return b.UnderlyingPrice * math.Sqrt(b.TimeToMaturity) * PD(d1) * math.Exp(-b.Dividend*b.TimeToMaturity) / 100.0
}

func (b *BlackScholes) Greeks() *Greeks {
	impliedVolatility := b.IV()
	return &Greeks{
		Delta: b.Delta(impliedVolatility),
		Gamma: b.Gamma(impliedVolatility),
		Theta: b.Theta(impliedVolatility),
		Rho:   b.Rho(impliedVolatility),
		Vega:  b.Vega(impliedVolatility),
	}
}
