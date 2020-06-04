package coldet

// Axis aligned bounding box
type AABB struct {
	position [3]float32
	width    float32 // X axis
	length   float32 // Z axis
	height   float32 // Y axis
}

type Point struct {
	position [3]float32
}
type Sphere struct {
	position [3]float32
	radius   float32
}

// X returns the x component of the position.
func (s *Sphere) X() float32 {
	return s.position[0]
}

// Y returns the y component of the position.
func (s *Sphere) Y() float32 {
	return s.position[1]
}

// Z returns the z component of the position.
func (s *Sphere) Z() float32 {
	return s.position[2]
}

// Radius returns the radius of the Sphere.
func (s *Sphere) Radius() float32 {
	return s.radius
}

// X returns the x component of the position.
func (p *Point) X() float32 {
	return p.position[0]
}

// Y returns the y component of the position.
func (p *Point) Y() float32 {
	return p.position[1]
}

// Z returns the z component of the position.
func (p *Point) Z() float32 {
	return p.position[2]
}

// X returns the x component of the position.
func (a *AABB) X() float32 {
	return a.position[0]
}

// Y returns the y component of the position.
func (a *AABB) Y() float32 {
	return a.position[1]
}

// Z returns the z component of the position.
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

// CheckPointVsAabb returns true if the given point is inside the AABB.
func CheckPointInAabb(p Point, b AABB) bool {
	inX := p.X() > b.X()-b.Width()/2 && p.X() < b.Width()/2
	inY := p.Y() > b.Y()-b.Height()/2 && p.Z() < b.Height()/2
	inZ := p.Z() > b.Z()-b.Length()/2 && p.Z() < b.Length()/2

	return inX && inY && inZ
}

// CheckPointInSphere returns true if the given point is inside the Sphere.
// Instead of the distance, the distance square is compared to the radius square.
func CheckPointInSphere(p Point, s Sphere) bool {
	distanceSquare := (p.X()-s.X())*(p.X()-s.X()) + (p.Y()-s.Y())*(p.Y()-s.Y()) + (p.Z()-s.Z())*(p.Z()-s.Z())

	return distanceSquare < s.Radius()*s.Radius()
}

// CheckSphereVsSphere returns true if the given spheres intersect.
// Instead of the distance, the distance square is compared to the sum of radius square.
func CheckSphereVsSphere(s1, s2 Sphere) bool {
	distanceSquare := (s1.X()-s2.X())*(s1.X()-s2.X()) + (s1.Y()-s2.Y())*(s1.Y()-s2.Y()) + (s1.Z()-s2.Z())*(s1.Z()-s2.Z())

	return distanceSquare < (s1.Radius()+s2.Radius())*(s1.Radius()+s2.Radius())
}
