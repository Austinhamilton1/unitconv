package distconv

import "fmt"

type Nanometer float64
type Micrometer float64
type Millimeter float64
type Centimeter float64
type Decimeter float64
type Meter float64
type Kilometer float64

type Inch float64
type Foot float64
type Yard float64
type Mile float64

func (n Nanometer) String() string {return fmt.Sprintf("%gnm", n)}
func (m Micrometer) String() string {return fmt.Sprintf("%gµm", m)}
func (m Millimeter) String() string {return fmt.Sprintf("%gmm", m)}
func (c Centimeter) String() string {return fmt.Sprintf("%gcm", c)}
func (d Decimeter) String() string {return fmt.Sprintf("%gdm", d)}
func (m Meter) String() string {return fmt.Sprintf("%gm", m)}
func (k Kilometer) String() string {return fmt.Sprintf("%gkm", k)}

func (i Inch) String() string {return fmt.Sprintf("%gin", i)}
func (f Foot) String() string {return fmt.Sprintf("%gft", f)}
func (y Yard) String() string {return fmt.Sprintf("%gyd", y)}
func (m Mile) String() string {return fmt.Sprintf("%gmi", m)}

func () ToNanometer() {return }
func () ToMicrometer() {return }
func () ToMillimeter() {return }
func () ToCentimeter() {return }
func () ToDecimeter() {return }
func () ToMeter() {return }
func () ToKilometer() {return }
func () ToInch() {return }
func () ToFoot() {return }
func () ToYard() {return }
func () ToMile() {return }