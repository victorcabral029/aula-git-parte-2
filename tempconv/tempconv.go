package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	// AbsoluteZeroC representa a temperatura zero absoluto em Celsius
	AbsoluteZeroC Celsius = -273.15

	// FreezingC representa a temperatura de congelamento da água em Celsius
	FreezingC     Celsius = 0

	// BoilingC representa a temperatura de ebulição da água em Celsius
	BoilingC      Celsius = 100
)

// String imprime uma temperatura 'n' em Celsius no formato n°C
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// String imprime uma temperatura 'n' em Fahrenheit no formato n°F
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// CToF converte uma temperatura em Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converte uma temperatura em Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KtoF converte uma temperatura em Kelvin para Fahrenheit
func KtoF(k Kelvin) Fahrenheit { return Fahrenheit(32 + (k - 273) * 9/5 )}

// FtoK converte uma temperatura em Fahrenheit para Kelvin
func FtoK(f Fahrenheit) Kelvin { return Kelvin( (f - 32) * 5/9 + 273  )}
