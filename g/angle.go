package g

import "math"

// constants
const (
	Pi  = math.Pi
	Tau = Pi * 2
	Phi = math.Phi
)

// Radian angle in radians
type Radian = float32

// DegToRad degree to radians
func DegToRad(deg float32) Radian { return deg * Pi / 180 }

// RadToDeg radians to degree
func RadToDeg(rad float32) float32 { return rad * 180 / Pi }

// Cos cosine
func Cos(v Radian) float32 { return float32(math.Cos(float64(v))) }

// Sin sine
func Sin(v Radian) float32 { return float32(math.Sin(float64(v))) }

// Atan returns the arctangent, in radians, of x.
func Atan(v float32) float32 { return float32(math.Atan(float64(v))) }

// Atan2 returns the arc tangent of y/x, using the signs of the two to determine the quadrant of the return value.
func Atan2(y, x float32) float32 { return float32(math.Atan2(float64(y), float64(x))) }
