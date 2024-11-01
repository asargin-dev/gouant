package gouant

import (
	"math"
	"testing"
)

func TestBlackScholes_Price(t *testing.T) {
	bs := &BlackScholes{
		UnderlyingPrice: 112.5,
		StrikePrice:     190,
		TimeToMaturity:  0.21095890410958903,
		RiskRate:        0.5,
		Dividend:        0,
		DerivativePrice: 1.2,
		IsCall:          true,
	}

	// Test the price calculation
	sigma := bs.IV()
	expectedPrice := 1.2
	calculatedPrice := bs.Price(sigma)

	if math.Abs(calculatedPrice-expectedPrice) > 1e-4 {
		t.Errorf("Price() = %.4f; want %.4f", calculatedPrice, expectedPrice)
	}
}

func TestBlackScholes_IV(t *testing.T) {
	bs := &BlackScholes{
		UnderlyingPrice: 112.5,
		StrikePrice:     190,
		TimeToMaturity:  0.21095890410958903,
		RiskRate:        0.5,
		Dividend:        0,
		DerivativePrice: 1.2,
		IsCall:          true,
	}

	// Test the IV calculation
	expectedIV := 0.616449844803655
	calculatedIV := bs.IV()

	if math.Abs(calculatedIV-expectedIV) > 1e-4 {
		t.Errorf("IV() = %.4f; want %.4f", calculatedIV, expectedIV)
	}
}

func TestBlackScholes_Delta(t *testing.T) {
	bs := &BlackScholes{
		UnderlyingPrice: 112.5,
		StrikePrice:     190,
		TimeToMaturity:  0.21095890410958903,
		RiskRate:        0.5,
		Dividend:        0,
		DerivativePrice: 1.2,
		IsCall:          true,
	}

	// Test the Delta calculation
	sigma := bs.IV()
	expectedDelta := 0.09063789804591085
	calculatedDelta := bs.Delta(sigma)

	if math.Abs(calculatedDelta-expectedDelta) > 1e-4 {
		t.Errorf("Delta() = %.4f; want %.4f", calculatedDelta, expectedDelta)
	}
}

func TestBlackScholes_Gamma(t *testing.T) {
	bs := &BlackScholes{
		UnderlyingPrice: 112.5,
		StrikePrice:     190,
		TimeToMaturity:  0.21095890410958903,
		RiskRate:        0.5,
		Dividend:        0,
		DerivativePrice: 1.2,
		IsCall:          true,
	}

	// Test the Gamma calculation
	sigma := bs.IV()
	expectedGamma := 0.005124951216808061
	calculatedGamma := bs.Gamma(sigma)

	if math.Abs(calculatedGamma-expectedGamma) > 1e-4 {
		t.Errorf("Gamma() = %.4f; want %.4f", calculatedGamma, expectedGamma)
	}
}

func TestBlackScholes_Theta(t *testing.T) {
	bs := &BlackScholes{
		UnderlyingPrice: 112.5,
		StrikePrice:     190,
		TimeToMaturity:  0.21095890410958903,
		RiskRate:        0.5,
		Dividend:        0,
		DerivativePrice: 1.2,
		IsCall:          true,
	}

	// Test the Theta calculation
	sigma := bs.IV()
	expectedTheta := -0.04608938503448249
	calculatedTheta := bs.Theta(sigma)

	if math.Abs(calculatedTheta-expectedTheta) > 1e-4 {
		t.Errorf("Theta() = %.4f; want %.4f", calculatedTheta, expectedTheta)
	}
}

func TestBlackScholes_Rho(t *testing.T) {
	bs := &BlackScholes{
		UnderlyingPrice: 112.5,
		StrikePrice:     190,
		TimeToMaturity:  0.21095890410958903,
		RiskRate:        0.5,
		Dividend:        0,
		DerivativePrice: 1.2,
		IsCall:          true,
	}

	// Test the Rho calculation
	sigma := bs.IV()
	expectedRho := 0.01897947373000337
	calculatedRho := bs.Rho(sigma)

	if math.Abs(calculatedRho-expectedRho) > 1e-4 {
		t.Errorf("Rho() = %.4f; want %.4f", calculatedRho, expectedRho)
	}
}

func TestBlackScholes_Vega(t *testing.T) {
	bs := &BlackScholes{
		UnderlyingPrice: 112.5,
		StrikePrice:     190,
		TimeToMaturity:  0.21095890410958903,
		RiskRate:        0.5,
		Dividend:        0,
		DerivativePrice: 1.2,
		IsCall:          true,
	}

	// Test the Vega calculation
	sigma := bs.IV()
	expectedVega := 0.08435102979004167
	calculatedVega := bs.Vega(sigma)

	if math.Abs(calculatedVega-expectedVega) > 1e-4 {
		t.Errorf("Vega() = %.4f; want %.4f", calculatedVega, expectedVega)
	}
}

func TestBlackScholes_Greeks(t *testing.T) {
	bs := &BlackScholes{
		UnderlyingPrice: 112.5,
		StrikePrice:     190,
		TimeToMaturity:  0.21095890410958903,
		RiskRate:        0.5,
		Dividend:        0,
		DerivativePrice: 1.2,
		IsCall:          true,
	}
	greeks := bs.Greeks()

	expectedDelta := 0.09063789804591085
	expectedGamma := 0.005124951216808061
	expectedTheta := -0.04608938503448249
	expectedRho := 0.01897947373000337
	expectedVega := 0.08435102979004167

	tolerance := 1e-4

	if math.Abs(greeks.Delta-expectedDelta) > tolerance {
		t.Errorf("Delta is incorrect. Expected %.6f, got %.6f", expectedDelta, greeks.Delta)
	}

	if math.Abs(greeks.Gamma-expectedGamma) > tolerance {
		t.Errorf("Gamma is incorrect. Expected %.6f, got %.6f", expectedGamma, greeks.Gamma)
	}

	if math.Abs(greeks.Theta-expectedTheta) > tolerance {
		t.Errorf("Theta is incorrect. Expected %.6f, got %.6f", expectedTheta, greeks.Theta)
	}

	if math.Abs(greeks.Rho-expectedRho) > tolerance {
		t.Errorf("Rho is incorrect. Expected %.6f, got %.6f", expectedRho, greeks.Rho)
	}

	if math.Abs(greeks.Vega-expectedVega) > tolerance {
		t.Errorf("Vega is incorrect. Expected %.6f, got %.6f", expectedVega, greeks.Vega)
	}
}
