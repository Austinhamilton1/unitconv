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

func (n Nanometer) ToMicrometer() Micrometer {return Micrometer(1000*n)}
func (n Nanometer) ToMillimeter() Millimeter {return Millimeter(1000000*n)}
func (n Nanometer) ToCentimeter() Centimeter {return Centimeter(10000000*n)}
func (n Nanometer) ToDecimeter() Decimeter {return Decimeter(100000000*n)}
func (n Nanometer) ToMeter() Meter {return Meter(1000000000*n)}
func (n Nanometer) ToKilometer() Kilometer {return Kilometer(1000000000000*n)}
func (n Nanometer) ToInch() Inch {return Inch(n.ToCentimeter().ToInch())}
func (n Nanometer) ToFoot() Foot {return Foot(12 * n.ToInch())}
func (n Nanometer) ToYard() Yard {return Yard(3 * n.ToFoot())}
func (n Nanometer) ToMile() Mile {return Mile(5280 * n.ToFoot())}

func (m Micrometer) ToNanometer() Nanometer {return Nanometer(m/1000)}
func (m Micrometer) ToMillimeter() Millimeter {return Millimeter(1000*m)}
func (m Micrometer) ToCentimeter() Centimeter {return Centimeter(10000*m)}
func (m Micrometer) ToDecimeter() Decimeter {return Decimeter(100000*m)}
func (m Micrometer) ToMeter() Meter {return Meter(1000000*m)}
func (m Micrometer) ToKilometer() Kilometer {return Kilometer(1000000000*m)}
func (m Micrometer) ToInch() Inch {return Inch(m.ToCentimeter().ToInch())}
func (m Micrometer) ToFoot() Foot {return Foot(12 * m.ToInch())}
func (m Micrometer) ToYard() Yard {return Yard(3 * m.ToFoot())}
func (m Micrometer) ToMile() Mile {return Mile(5280 * m.ToFoot())}

func (m Millimeter) ToNanometer() Nanometer {return Nanometer(m/1000000)}
func (m Millimeter) ToMicrometer()  Micrometer {return Micrometer(m/1000)}
func (m Millimeter) ToCentimeter() Centimeter {return Centimeter(10*m)}
func (m Millimeter) ToDecimeter() Decimeter {return Decimeter(100*m)}
func (m Millimeter) ToMeter() Meter {return Meter(1000*m)}
func (m Millimeter) ToKilometer() Kilometer {return Kilometer(1000000*m)}
func (m Millimeter) ToInch() Inch {return Inch(m.ToCentimeter().ToInch())}
func (m Millimeter) ToFoot() Foot {return Foot(12 * m.ToInch())}
func (m Millimeter) ToYard() Yard {return Yard(3 * m.ToFoot())}
func (m Millimeter) ToMile() Mile {return Mile(5280 * m.ToFoot())}

func (c Centimeter) ToNanometer() Nanometer {return Nanometer(c/10000000)}
func (c Centimeter) ToMicrometer()  Micrometer {return Micrometer(c/10000)}
func (c Centimeter) ToMillimeter() Millimeter {return Millimeter(c/10)}
func (c Centimeter) ToDecimeter() Decimeter {return Decimeter(10*c)}
func (c Centimeter) ToMeter() Meter {return Meter(100*c)}
func (c Centimeter) ToKilometer() Kilometer {return Kilometer(100000*c)}
func (c Centimeter) ToInch() Inch {return Inch(c/2.54)}
func (c Centimeter) ToFoot() Foot {return Foot(12 * c.ToInch())}
func (c Centimeter) ToYard() Yard {return Yard(3 * c.ToFoot())}
func (c Centimeter) ToMile() Mile {return Mile(5280 * c.ToFoot())}

func (d Decimeter) ToNanometer() Nanometer {return Nanometer(d/100000000)}
func (d Decimeter) ToMicrometer()  Micrometer {return Micrometer(d/100000)}
func (d Decimeter) ToMillimeter() Millimeter {return Millimeter(d/100)}
func (d Decimeter) ToCentimeter() Centimeter {return Centimeter(d/10)}
func (d Decimeter) ToMeter() Meter {return Meter(10*d)}
func (d Decimeter) ToKilometer() Kilometer {return Kilometer(10000*d)}
func (d Decimeter) ToInch() Inch {return Inch(d.ToCentimeter().ToInch())}
func (d Decimeter) ToFoot() Foot {return Foot(12 * d.ToInch())}
func (d Decimeter) ToYard() Yard {return Yard(3 * d.ToFoot())}
func (d Decimeter) ToMile() Mile {return Mile(5280 * d.ToFoot())}

func (m Meter) ToNanometer() Nanometer {return Nanometer(m/1000000000)}
func (m Meter) ToMicrometer()  Micrometer {return Micrometer(m/1000000)}
func (m Meter) ToMillimeter() Millimeter {return Millimeter(m/1000)}
func (m Meter) ToCentimeter() Centimeter {return Centimeter(m/100)}
func (m Meter) ToDecimeter() Decimeter {return Decimeter(m/10)}
func (m Meter) ToKilometer() Kilometer {return Kilometer(1000*m)}
func (m Meter) ToInch() Inch {return Inch(m.ToCentimeter().ToInch())}
func (m Meter) ToFoot() Foot {return Foot(12 * m.ToInch())}
func (m Meter) ToYard() Yard {return Yard(3 * m.ToFoot())}
func (m Meter) ToMile() Mile {return Mile(5280 * m.ToFoot())}

func (k Kilometer) ToNanometer() Nanometer {return Nanometer(k/1000000000000)}
func (k Kilometer) ToMicrometer()  Micrometer {return Micrometer(k/1000000000)}
func (k Kilometer) ToMillimeter() Millimeter {return Millimeter(k/1000000)}
func (k Kilometer) ToCentimeter() Centimeter {return Centimeter(k/100000)}
func (k Kilometer) ToDecimeter() Decimeter {return Decimeter(k/10000)}
func (k Kilometer) ToMeter() Meter {return Meter(k/1000)}
func (k Kilometer) ToInch() Inch {return Inch(k.ToCentimeter().ToInch())}
func (k Kilometer) ToFoot() Foot {return Foot(12 * k.ToInch())}
func (k Kilometer) ToYard() Yard {return Yard(3 * k.ToFoot())}
func (k Kilometer) ToMile() Mile {return Mile(5280 * k.ToFoot())}

func (i Inch) ToNanometer() Nanometer {return Nanometer(i.ToCentimeter().ToNanometer())}
func (i Inch) ToMicrometer()  Micrometer {return Micrometer(i.ToCentimeter().ToMicrometer())}
func (i Inch) ToMillimeter() Millimeter {return Millimeter(i.ToCentimeter().ToMillimeter())}
func (i Inch) ToCentimeter() Centimeter {return Centimeter(2.54*i)}
func (i Inch) ToDecimeter() Decimeter {return Decimeter(i.ToCentimeter().ToDecimeter())}
func (i Inch) ToMeter() Meter {return Meter(i.ToCentimeter().ToMeter())}
func (i Inch) ToKilometer() Kilometer {return Kilometer(i.ToCentimeter().ToKilometer())}
func (i Inch) ToFoot() Foot {return Foot(12*i)}
func (i Inch) ToYard() Yard {return Yard(36*i)}
func (i Inch) ToMile() Mile {return Mile(5280 * i.ToFoot())}

func (f Foot) ToNanometer() Nanometer {return Nanometer(f.ToInch().ToCentimeter().ToNanometer())}
func (f Foot) ToMicrometer()  Micrometer {return Micrometer(f.ToInch().ToCentimeter().ToMicrometer())}
func (f Foot) ToMillimeter() Millimeter {return Millimeter(f.ToInch().ToCentimeter().ToMillimeter())}
func (f Foot) ToCentimeter() Centimeter {return Centimeter(f.ToInch().ToCentimeter())}
func (f Foot) ToDecimeter() Decimeter {return Decimeter(f.ToInch().ToCentimeter().ToDecimeter())}
func (f Foot) ToMeter() Meter {return Meter(f.ToInch().ToCentimeter().ToMeter())}
func (f Foot) ToKilometer() Kilometer {return Kilometer(f.ToInch().ToCentimeter().ToKilometer())}
func (f Foot) ToInch() Inch {return Inch(f/12)}
func (f Foot) ToYard() Yard {return Yard(3*f)}
func (f Foot) ToMile() Mile {return Mile(5280*f)}

func (y Yard) ToNanometer() Nanometer {return Nanometer(y.ToInch().ToCentimeter().ToNanometer())}
func (y Yard) ToMicrometer()  Micrometer {return Micrometer(y.ToInch().ToCentimeter().ToMicrometer())}
func (y Yard) ToMillimeter() Millimeter {return Millimeter(y.ToInch().ToCentimeter().ToMillimeter())}
func (y Yard) ToCentimeter() Centimeter {return Centimeter(y.ToInch().ToCentimeter())}
func (y Yard) ToDecimeter() Decimeter {return Decimeter(y.ToInch().ToCentimeter().ToDecimeter())}
func (y Yard) ToMeter() Meter {return Meter(y.ToInch().ToCentimeter().ToMeter())}
func (y Yard) ToKilometer() Kilometer {return Kilometer(y.ToInch().ToCentimeter().ToKilometer())}
func (y Yard) ToInch() Inch {return Inch(y/36)}
func (y Yard) ToFoot() Foot {return Foot(y/3)}
func (y Yard) ToMile() Mile {return Mile(1760*y)}

func (m Mile) ToNanometer() Nanometer {return Nanometer(m.ToInch().ToCentimeter().ToNanometer())}
func (m Mile) ToMicrometer()  Micrometer {return Micrometer(m.ToInch().ToCentimeter().ToMicrometer())}
func (m Mile) ToMillimeter() Millimeter {return Millimeter(m.ToInch().ToCentimeter().ToMillimeter())}
func (m Mile) ToCentimeter() Centimeter {return Centimeter(m.ToInch().ToCentimeter())}
func (m Mile) ToDecimeter() Decimeter {return Decimeter(m.ToInch().ToCentimeter().ToDecimeter())}
func (m Mile) ToMeter() Meter {return Meter(m.ToInch().ToCentimeter().ToMeter())}
func (m Mile) ToKilometer() Kilometer {return Kilometer(m.ToInch().ToCentimeter().ToKilometer())}
func (m Mile) ToInch() Inch {return Inch(m.ToFoot().ToInch())}
func (m Mile) ToFoot() Foot {return Foot(m/5280)}
func (m Mile) ToYard() Yard {return Yard(m/1760)}