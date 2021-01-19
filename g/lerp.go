package g

// Lerp Lerp
func Lerp(a, b, p float32) float32 { return (b-a)*p + a }

// LerpClamp lerp after clamped to 0..1
func LerpClamp(a, b, p float32) float32 { return Lerp(a, b, Clamp01(p)) }
