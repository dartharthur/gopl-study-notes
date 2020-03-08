// Package tempconv peforms Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	AbsoluteZeroF Fahrenheit = -459.7
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212

	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.16
	BoilingK      Kelvin = 373.16
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
