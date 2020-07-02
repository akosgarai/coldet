package coldet

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

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

func NewBoundingPoint(pos [3]float32) *Point {
	return &Point{pos}
}
func NewBoundingSphere(pos [3]float32, radius float32) *Sphere {
	return &Sphere{pos, radius}
}
func NewBoundingBox(pos [3]float32, width, height, length float32) *AABB {
	return &AABB{pos, width, length, height}
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

// Distance returns the distance from the given point.
func (s *Sphere) Distance(to [3]float32) float32 {
	pointPos := mgl32.Vec3{s.position[0], s.position[1], s.position[2]}
	toPos := mgl32.Vec3{to[0], to[1], to[2]}
	return pointPos.Sub(toPos).Len() - s.radius
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

// Distance returns the distance from the given point.
func (p *Point) Distance(to [3]float32) float32 {
	pointPos := mgl32.Vec3{p.position[0], p.position[1], p.position[2]}
	toPos := mgl32.Vec3{to[0], to[1], to[2]}
	return pointPos.Sub(toPos).Len()
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

// Distance returns the distance from the given point.
func (a *AABB) Distance(to [3]float32) float32 {
	toPos := mgl32.Vec3{to[0], to[1], to[2]}
	// Get the closest point to the
	cX := clamp(toPos.X(), a.X()-a.Width()/2, a.X()+a.Width()/2)
	cY := clamp(toPos.Y(), a.Y()-a.Height()/2, a.Y()+a.Height()/2)
	cZ := clamp(toPos.Z(), a.Z()-a.Length()/2, a.Z()+a.Length()/2)
	closestPos := mgl32.Vec3{cX, cY, cZ}
	return closestPos.Sub(toPos).Len()
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

// CheckSphereVsAabb returns true if the given sphere intersects with the given bb.
func CheckSphereVsAabb(s Sphere, b AABB) bool {
	// Get the closest point to the sphere center
	cX := clamp(s.X(), b.X()-b.Width()/2, b.X()+b.Width()/2)
	cY := clamp(s.Y(), b.Y()-b.Height()/2, b.Y()+b.Height()/2)
	cZ := clamp(s.Z(), b.Z()-b.Length()/2, b.Z()+b.Length()/2)

	distanceSquare := (cX-s.X())*(cX-s.X()) + (cY-s.Y())*(cY-s.Y()) + (cZ-s.Z())*(cZ-s.Z())

	return distanceSquare < s.Radius()*s.Radius()
}

func clamp(value, min, max float32) float32 {
	return float32(math.Max(float64(min), math.Min(float64(value), float64(max))))
}
