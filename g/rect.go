package g

// Rect rectangle
type Rect struct{ Min, Max V2 }

// NewRect a new instance of Rect
func NewRect(w, h float32) Rect {
	return Rect{
		V2{-w / 2, -h / 2},
		V2{w / 2, h / 2},
	}
}

// NewCircleRect a new instance of Rect bounding circle of radius r
func NewCircleRect(r float32) Rect {
	return Rect{V2{-r, -r}, V2{r, r}}
}

// Size size of a rect
func (r Rect) Size() V2 { return r.Max.Sub(r.Min) }

// Offset offsets the rect by delta
func (r Rect) Offset(delta V2) Rect {
	return Rect{
		r.Min.Add(delta),
		r.Max.Add(delta),
	}
}

// ScaleInv scale inv
func (r Rect) ScaleInv(v V2) Rect {
	return Rect{
		V2{
			r.Min.X / v.X,
			r.Min.Y / v.Y,
		},
		V2{
			r.Max.X / v.X,
			r.Max.Y / v.Y,
		},
	}
}

// Contains if rect contains point
func (r Rect) Contains(p V2) bool {
	return (r.Min.X <= p.X) && (p.X <= r.Max.X) &&
		(r.Min.Y <= p.Y) && (p.Y <= r.Max.Y)
}

// EnforceInside enforces inside
func EnforceInside(pos, vel *V2, bounds Rect, dampening float32) {
	minx, maxx := MinMax(bounds.Min.X, bounds.Max.X)
	if pos.X < minx {
		pos.X = minx
		vel.X = +Abs(vel.X) * dampening
	}
	if pos.X > maxx {
		pos.X = maxx
		vel.X = -Abs(vel.X) * dampening
	}

	miny, maxy := MinMax(bounds.Min.Y, bounds.Max.Y)
	if pos.Y < miny {
		pos.Y = miny
		vel.Y = +Abs(vel.Y) * dampening
	}
	if pos.Y > maxy {
		pos.Y = maxy
		vel.Y = -Abs(vel.Y) * dampening
	}
}
