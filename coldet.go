package coldet

// Axis aligned bounding box
type AABB struct {
	position [3]float32
	width    float32 // X axis
	length   float32 // Z axis
	height   float32 // Y axis
}

// X returns the x component of the position.
func (a *AABB) X() float32 {
	return a.position[0]
}

// Y returns the y component of the position.
func (a *AABB) Y() float32 {
	return a.position[1]
}

// Z returns the z component of the position/
func (a *AABB) Z() float32 {
	return a.position[2]
}

// Width returns the width of the bb.
func (a *AABB) Width() float32 {
	return a.width
}

// Length returns the length of the bb.
func (a *AABB) Length() float32 {
	return a.length
}

// Height returns the height of the bb.
func (a *AABB) Height() float32 {
	return a.height
}

// CheckAabbVsAabb returns true if the given two object has been collided.
func CheckAabbVsAabb(b1, b2 AABB) bool {
	colX := b1.X()+b1.Width()/2 >= b2.X()-b2.Width()/2 && b2.X()+b2.Width()/2 >= b1.X()-b1.Width()/2
	colY := b1.Y()+b1.Height()/2 >= b2.Y()-b2.Height()/2 && b2.Y()+b2.Height()/2 >= b1.Y()-b1.Height()/2
	colZ := b1.Z()+b1.Length()/2 >= b2.Z()-b2.Length()/2 && b2.Z()+b2.Length()/2 >= b1.Z()-b1.Length()/2

	return colX && colY && colZ
}
