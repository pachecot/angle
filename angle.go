package angle

import "math"

// Angle is a generalized angular measurement type
type Angle float64

const (
	radian    = 1
	circle    = 2 * math.Pi
	degree    = 180 / math.Pi
	arcminute = degree * 60
	minute    = degree * 4
	hour      = minute / 60
	day       = hour / 24
)

// Degrees creates Angle from degrees
func Degrees(degrees float64) Angle {
	return Angle(degrees / degree)
}

// Radians creates Angle from radians
func Radians(radians float64) Angle {
	return Angle(radians / radian)
}

// Minutes creates Angle from minutes of time
func Minutes(minutes float64) Angle {
	return Angle(minutes / minute)
}

// Radians returns the radians of the angle
func (ang Angle) Radians() float64 {
	return float64(ang * radian)
}

// Degrees returns the degrees of the angle
func (ang Angle) Degrees() float64 {
	return float64(ang * degree)
}

// Arcminute returns the Arcminute of the angle
func (ang Angle) Arcminute() float64 {
	return float64(ang * arcminute)
}

// Minutes returns the Minutes (time) from the angle
func (ang Angle) Minutes() float64 {
	return float64(ang * minute)
}

// Hours returns the Hours (time) from the angle
func (ang Angle) Hours() float64 {
	return float64(ang * hour)
}

// Days returns the Days (time) from the angle
func (ang Angle) Days() float64 {
	return float64(ang * day)
}

// Mod returns a normalized angle from -360 to 360 degrees
//
//	Degrees(450).Mod() = Degrees(90)
//	Degrees(-450).Mod() = Degrees(-90)
func (ang Angle) Mod() Angle {
	return Angle(math.Mod(float64(ang), circle))
}

// Abs returns a positive of the angle
//
//	Degrees(-90).Abs() = Degrees(270)
func (ang Angle) Abs() Angle {
	if ang >= 0 {
		return ang
	}
	return Angle(math.Abs(float64(circle - ang)))
}

// Equal compares to another angle and returns true if
func (ang Angle) Equal(other Angle) bool {
	return math.Abs(float64(ang-other)) < 0.000_000_002
}

// Sin returns the sine of the ang argument.
func (ang Angle) Sin() float64 {
	return math.Sin(float64(ang * radian))
}

// Cos returns the cosine of the ang argument.
func (ang Angle) Cos() float64 {
	return math.Cos(float64(ang * radian))
}

// Tan returns the tangent of the ang argument.
func (ang Angle) Tan() float64 {
	return math.Tan(float64(ang * radian))
}

// Asin returns the Angle of the arcsine of x.
func Asin(x float64) Angle {
	return Angle(math.Asin(x) / radian)
}

// Acos returns the Angle of the arccosine of x.
func Acos(x float64) Angle {
	return Angle(math.Acos(x) / radian)
}

// Acos returns the Angle of the arctangent of x.
func Atan(x float64) Angle {
	return Angle(math.Atan(x) / radian)
}

// Acos returns the Angle of the arctangent of y/x, using the signs of the
// two to determine the quadrant of the return value.
func Atan2(y, x float64) Angle {
	return Angle(math.Atan2(y, x) / radian)
}
