// package tempconv offers tools for converting between temperature units
package tempconv

import "fmt"

type Kelvin float64
type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZero Kelvin = 0
	FreezingC Celcius = 0
	BoilingC Celcius = 100
	FreezingF Fahrenheit = 32
	BoilingF Fahrenheit = 212
)

func (k Kelvin) String() string {return fmt.Sprintf("%gK", k)}
func (c Celcius) String() string {return fmt.Sprintf("%g°C", c)}
func (f Fahrenheit) String() string {return fmt.Sprintf("%g°F", f)}

//KtoC converts a Kelvin value to its corresponding Celcius value
func (k Kelvin) ToCelcius() Celcius {return Celcius(k - 273.15)}

//KtoF converts a Kelvin value to its corresponding Fahrenheit value
func (k Kelvin) ToFahrenheit() Fahrenheit {return Fahrenheit(k.ToCelcius().ToFahrenheit())}

//CtoK converts a Celcius value to its corresponding Kelvin value
func (c Celcius) ToKelvin() Kelvin {return Kelvin(c + 273.15)}

//CtoF converts a Celcius value to its corresponding Fahrenheit value
func (c Celcius) ToFahrenheit() Fahrenheit {return Fahrenheit(c * 9/5 + 32)}

//FtoK converts a Fahrenheit value to its corresponding Kelvin value
func (f Fahrenheit) ToKelvin() Kelvin {return Kelvin(f.ToCelcius().ToKelvin())}

//FtoC converts a Fahrenheit value to its corresponding Celcius value
func (f Fahrenheit) ToCelcius() Celcius {return Celcius((f - 32) * 5/9)}