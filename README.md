# Gouant: Black-Scholes Options Pricing Library

## Overview

Gouant is a Go library for calculating option pricing and Greeks using the Black-Scholes model. It provides comprehensive functionality for options pricing, including price calculation, implied volatility, and option Greeks.

## Features

- Calculate option prices for both call and put options
- Compute option Greeks (Delta, Gamma, Theta, Rho, Vega)
- Calculate implied volatility using Newton-Raphson method
- Support for dividend-paying underlying assets
- Precise mathematical calculations using standard financial formulas

## Installation

```bash
go get github.com/asargin-dev/gouant
```

## Usage Example

```go
package main

import (
    "fmt"
    "github.com/asargin-dev/gouant"
)

func main() {
    // Create a Black-Scholes option
    option := &gouant.BlackScholes{
        UnderlyingPrice: 100.0,
        StrikePrice:     105.0,
        TimeToMaturity:  1.0,
        RiskRate:        0.05,
        Dividend:        0.02,
        IsCall:          true,
        DerivativePrice: 5.0,
    }

    // Calculate option price
    price := option.Price(0.25)
    fmt.Printf("Option Price: %.2f\n", price)

    // Get option Greeks
    greeks := option.Greeks()
    fmt.Printf("Delta: %.4f\n", greeks.Delta)
    fmt.Printf("Gamma: %.4f\n", greeks.Gamma)
}
```

## Methods

### BlackScholes Struct Methods

- `Price(sigma float64)`: Calculate option price
- `IV()`: Calculate implied volatility
- `Delta(sigma float64)`: Calculate option delta
- `Gamma(sigma float64)`: Calculate option gamma
- `Theta(sigma float64)`: Calculate option theta
- `Rho(sigma float64)`: Calculate option rho
- `Vega(sigma float64)`: Calculate option vega
- `Greeks()`: Calculate all Greeks using implied volatility

## Mathematical Background

The library implements the Black-Scholes option pricing model, which uses the following key components:
- Cumulative Normal Distribution (CND)
- Probability Density Function (PDF)
- Volatility calculation
- Risk-free rate
- Dividend yield
- Time to maturity

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Ahmet Sargın - contact@ahmetsargin.com

Project Link: [https://github.com/asargin-dev/gouant](https://github.com/asargin-dev/gouant)
