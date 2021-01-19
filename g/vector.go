package g

// Zero2 zero vec 2
var Zero2 = V2{}

// V2 vec2
type V2 struct{ X, Y float32 }

// RandomV2 random vec2
func RandomV2(r Rect) V2 {
	return V2{
		RandomBetween(r.Min.X, r.Max.X),
		RandomBetween(r.Min.Y, r.Max.Y),
	}
}

// RandomV2Circle random vec2 circle
func RandomV2Circle(radius float32) V2 {
	v := V2{RandomBetween(-1, 1), RandomBetween(-1, 1)}
	return v.Normalize().Scale(radius)
}

// XY returns both components
func (a V2) XY() (x, y float32) { return a.X, a.Y }

// XYZ returns x, y, 0
func (a V2) XYZ() (x, y, z float32) { return a.X, a.Y, 0 }

// Add adds two vectors and returns the result
func (a V2) Add(b V2) V2 { return V2{a.X + b.X, a.Y + b.Y} }

// AddScale adds b vector multiplied by scale to a and returns the result
func (a V2) AddScale(b V2, s float32) V2 { return V2{a.X + b.X*s, a.Y + b.Y*s} }

// Sub subtracts two vectors and returns the result
func (a V2) Sub(b V2) V2 { return V2{a.X - b.X, a.Y - b.Y} }

// Dot calculates the dot product
func (a V2) Dot(b V2) float32 { return a.X*b.X + a.Y*b.Y }

// Scale scales each component and returns the result
func (a V2) Scale(s float32) V2 { return V2{a.X * s, a.Y * s} }

// Length returns the length of the vector
func (a V2) Length() float32 { return Sqrt(a.X*a.X + a.Y*a.Y) }

// Length2 returns the squared length of the vector
func (a V2) Length2() float32 { return a.X*a.X + a.Y*a.Y }

// Distance returns the distance to vector b
func (a V2) Distance(b V2) float32 {
	dx, dy := a.X-b.X, a.Y-b.Y
	return Sqrt(dx*dx + dy*dy)
}

// Distance2 returns the squared distance to vector b
func (a V2) Distance2(b V2) float32 {
	dx, dy := a.X-b.X, a.Y-b.Y
	return dx*dx + dy*dy
}

// Normalize normalized vec2
func (a V2) Normalize() V2 {
	m := a.Length()
	if m < 1 {
		m = 1
	}
	return V2{a.X / m, a.Y / m}
}

// Negate Negate of a vec2
func (a V2) Negate() V2 { return V2{-a.X, -a.Y} }

// Cross product of a and b
func (a V2) Cross(b V2) float32 { return a.X*b.Y - a.Y*b.X }

// NearZero if vec2 is close to zero
func (a V2) NearZero() bool { return a.Length2() < 0.0001 }

// Rotate rotates a vec2 by an angle
func (a V2) Rotate(angle float32) V2 {
	cs, sn := Cos(angle), Sin(angle)
	return V2{a.X*cs - a.Y*sn, a.X*sn + a.Y*cs}
}

// Angle returns the angle for this absolute vector
func (a V2) Angle() float32 { return Atan2(a.Y, a.X) }

// Rotate90 rotates a vec2 by 90 degrees ccw
func (a V2) Rotate90() V2 { return V2{-a.Y, a.X} }

// Rotate90c rotates a vec2 by 90 degrees cw
func (a V2) Rotate90c() V2 { return V2{a.Y, -a.X} }

// Rotate180 rotates a vec2 by 180 degrees
func (a V2) Rotate180() V2 { return V2{-a.X, -a.Y} }
