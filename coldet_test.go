package coldet

import (
	"testing"
)

var (
	aabb   = AABB{[3]float32{1, 2, 3}, 4, 5, 6}
	p      = Point{[3]float32{1, 2, 3}}
	sphere = Sphere{[3]float32{1, 2, 3}, 1}
)

func TestNewBoundingPoint(t *testing.T) {
	bp := NewBoundingPoint(p.position)
	if bp.position != p.position {
		t.Error("Invalid position")
	}
}
func TestNewBoundingSphere(t *testing.T) {
	bs := NewBoundingSphere(sphere.position, sphere.radius)
	if bs.position != sphere.position {
		t.Error("Invalid position")
	}
	if bs.radius != sphere.radius {
		t.Error("Invalid radius")
	}
}
func TestNewBoundingBox(t *testing.T) {
	bb := NewBoundingBox(aabb.position, aabb.width, aabb.height, aabb.length)
	if bb.position != aabb.position {
		t.Error("Invalid position")
	}
	if bb.width != aabb.width {
		t.Error("Invalid width")
	}
	if bb.length != aabb.length {
		t.Error("Invalid length")
	}
	if bb.height != aabb.height {
		t.Error("Invalid height")
	}
}
func TestSphereX(t *testing.T) {
	if sphere.X() != 1 {
		t.Error("Invalid x coordinate.")
	}
}
func TestSphereY(t *testing.T) {
	if sphere.Y() != 2 {
		t.Error("Invalid y coordinate.")
	}
}
func TestSphereZ(t *testing.T) {
	if sphere.Z() != 3 {
		t.Error("Invalid z coordinate.")
	}
}
func TestSphereRadius(t *testing.T) {
	if sphere.Radius() != 1 {
		t.Error("Invalid radius.")
	}
}
func TestSphereClosest(t *testing.T) {
	point := [3]float32{2, 2, 3}
	closest := sphere.ClosestPoint(point)
	if closest != point {
		t.Errorf("Invalid closest point. Instead of '%v', we have '%v'.", point, closest)
	}
	pointNew := [3]float32{4, 2, 3}
	closest = sphere.ClosestPoint(pointNew)
	if closest != point {
		t.Errorf("Invalid closest point. Instead of '%v', we have '%v'.", point, closest)
	}
}
func TestSphereDistance(t *testing.T) {
	distance := sphere.Distance([3]float32{2, 2, 3})
	if distance != 0.0 {
		t.Errorf("Invalid distance. Instead of '0.0', we have '%f'.", distance)
	}
	distance = sphere.Distance([3]float32{4, 2, 3})
	if distance != 2.0 {
		t.Errorf("Invalid distance. Instead of '2.0', we have '%f'.", distance)
	}
}
func TestPointX(t *testing.T) {
	if p.X() != 1 {
		t.Error("Invalid x coordinate.")
	}
}
func TestPointY(t *testing.T) {
	if p.Y() != 2 {
		t.Error("Invalid y coordinate.")
	}
}
func TestPointZ(t *testing.T) {
	if p.Z() != 3 {
		t.Error("Invalid z coordinate.")
	}
}
func TestPointClosest(t *testing.T) {
	testData := [][3]float32{
		[3]float32{0, 0, 0},
		[3]float32{1, 1, 1},
		[3]float32{1, 0, 0},
		[3]float32{1, 0, 1},
	}
	for _, tt := range testData {
		closest := p.ClosestPoint(tt)
		if closest != p.position {
			t.Errorf("Invalid closest point. Instead of '%v, we have '%v'", p.position, closest)
		}

	}
}
func TestPointDistance(t *testing.T) {
	distance := p.Distance([3]float32{p.X(), p.Y(), p.Z()})
	if distance != 0.0 {
		t.Errorf("Invalid distance. Instead of '0.0', we have '%f'.", distance)
	}
}
func TestAABBX(t *testing.T) {
	if aabb.X() != 1 {
		t.Error("Invalid x coordinate.")
	}
}
func TestAABBY(t *testing.T) {
	if aabb.Y() != 2 {
		t.Error("Invalid y coordinate.")
	}
}
func TestAABBZ(t *testing.T) {
	if aabb.Z() != 3 {
		t.Error("Invalid z coordinate.")
	}
}
func TestAABBWidth(t *testing.T) {
	if aabb.Width() != 4 {
		t.Error("Invalid width.")
	}
}
func TestAABBLength(t *testing.T) {
	if aabb.Length() != 5 {
		t.Error("Invalid length.")
	}
}
func TestAABBHeight(t *testing.T) {
	if aabb.Height() != 6 {
		t.Error("Invalid heigth.")
	}
}
func TestAABBClosest(t *testing.T) {
	point := [3]float32{3, 2, 3}
	closest := aabb.ClosestPoint(point)
	if closest != point {
		t.Errorf("Invalid closest point. Instead of '%v', we have '%v'.", point, closest)
	}
	otherPoint := [3]float32{4, 2, 3}
	closest = aabb.ClosestPoint(otherPoint)
	if closest != point {
		t.Errorf("Invalid closest point. Instead of '%v', we have '%v'.", point, closest)
	}
}
func TestAABBDistance(t *testing.T) {
	distance := aabb.Distance([3]float32{3, 2, 3})
	if distance != 0.0 {
		t.Errorf("Invalid distance. Instead of '0.0', we have '%f'.", distance)
	}
	distance = aabb.Distance([3]float32{4, 2, 3})
	if distance != 1.0 {
		t.Errorf("Invalid distance. Instead of '1.0', we have '%f'.", distance)
	}
}

func TestCheckAabbVsAabb(t *testing.T) {
	testData := []struct {
		b1       AABB
		b2       AABB
		collided bool
	}{
		{AABB{[3]float32{0, 0, 0}, 1, 1, 1}, AABB{[3]float32{0, 0, 0}, 1, 1, 1}, true},
		{AABB{[3]float32{0, 0, 0}, 1, 1, 1}, AABB{[3]float32{0, 0, 0}, 0.5, 0.5, 0.5}, true},
		{AABB{[3]float32{0, 0, 0}, 1, 1, 1}, AABB{[3]float32{2, 2, 2}, 1, 1, 1}, false},
		{AABB{[3]float32{0, 0, 0}, 1, 1, 1}, AABB{[3]float32{1, 1, 1}, 1, 1, 1}, true},
		{AABB{[3]float32{0, 0, 0}, 1, 1, 1}, AABB{[3]float32{1, 1, 1}, 0.99, 1, 1}, false},
	}

	for _, tt := range testData {
		result := CheckAabbVsAabb(tt.b1, tt.b2)
		if result != tt.collided {
			t.Error("Invalid collision result.")
		}
	}
}
func TestCheckPointInAabb(t *testing.T) {
	testData := []struct {
		p  Point
		b  AABB
		in bool
	}{
		{Point{[3]float32{0, 0, 0}}, AABB{[3]float32{0, 0, 0}, 1, 1, 1}, true},
		{Point{[3]float32{0, 0, 0}}, AABB{[3]float32{1, 1, 1}, 1, 1, 1}, false},
		{Point{[3]float32{0, 0, 0}}, AABB{[3]float32{1, 1, 1}, 2, 2, 2}, false},
		{Point{[3]float32{0.5, 0.5, 0.5}}, AABB{[3]float32{0, 0, 0}, 1, 1, 1}, false},
		{Point{[3]float32{0.49, 0.49, 0.5}}, AABB{[3]float32{0, 0, 0}, 1, 1, 1}, false},
		{Point{[3]float32{0.49, 0.49, 0.49}}, AABB{[3]float32{0, 0, 0}, 1, 1, 1}, true},
	}
	for _, tt := range testData {
		result := CheckPointInAabb(tt.p, tt.b)
		if result != tt.in {
			t.Error("Invalid collision result.")
		}
	}
}
func TestCheckPointInSphere(t *testing.T) {
	testData := []struct {
		p  Point
		s  Sphere
		in bool
	}{
		{Point{[3]float32{0, 0, 0}}, Sphere{[3]float32{0, 0, 0}, 1}, true},
		{Point{[3]float32{0, 0, 0}}, Sphere{[3]float32{1, 1, 1}, 1}, false},
		{Point{[3]float32{0, 0, 0}}, Sphere{[3]float32{1, 1, 1}, 2}, true},
		{Point{[3]float32{0.5, 0.5, 0.5}}, Sphere{[3]float32{0, 0, 0}, 1}, true},
		{Point{[3]float32{0.49, 0.49, 0.5}}, Sphere{[3]float32{0, 0, 0}, 1}, true},
		{Point{[3]float32{0.49, 0.49, 0.49}}, Sphere{[3]float32{0, 0, 0}, 1}, true},
	}
	for _, tt := range testData {
		result := CheckPointInSphere(tt.p, tt.s)
		if result != tt.in {
			t.Error("Invalid collision result.")
			t.Log(tt)
		}
	}
}
func TestCheckSphereVsSphere(t *testing.T) {
	testData := []struct {
		s1       Sphere
		s2       Sphere
		collided bool
	}{
		{Sphere{[3]float32{0, 0, 0}, 1}, Sphere{[3]float32{0, 0, 0}, 1}, true},
		{Sphere{[3]float32{0, 0, 0}, 1}, Sphere{[3]float32{0, 0, 0}, 0.5}, true},
		{Sphere{[3]float32{0, 0, 0}, 1}, Sphere{[3]float32{2, 2, 2}, 1}, false},
		{Sphere{[3]float32{0, 0, 0}, 1}, Sphere{[3]float32{1, 1, 1}, 1}, true},
		{Sphere{[3]float32{0, 0, 0}, 1}, Sphere{[3]float32{1, 1, 1}, 0.5}, false},
		{Sphere{[3]float32{0, 0, 0}, 1}, Sphere{[3]float32{1, 1, 1}, 0.8}, true},
	}

	for _, tt := range testData {
		result := CheckSphereVsSphere(tt.s1, tt.s2)
		if result != tt.collided {
			t.Error("Invalid collision result.")
			t.Log(tt)
		}
	}
}
func TestCheckSphereVsAabb(t *testing.T) {
	testData := []struct {
		s        Sphere
		b        AABB
		collided bool
	}{
		{Sphere{[3]float32{0, 0, 0}, 1}, AABB{[3]float32{2, 2, 2}, 1, 1, 1}, false},
		{Sphere{[3]float32{0, 0, 0}, 1}, AABB{[3]float32{0, 0, 0}, 1, 1, 1}, true},
		{Sphere{[3]float32{0, 0, 0}, 1}, AABB{[3]float32{0, 0, 0}, 0.5, 0.5, 0.5}, true},
		{Sphere{[3]float32{0, 0, 0}, 1}, AABB{[3]float32{1, 1, 1}, 1, 1, 1}, true},
		{Sphere{[3]float32{0, 0, 0}, 1}, AABB{[3]float32{1, 1, 1}, 0.5, 0.5, 0.5}, false},
	}

	for _, tt := range testData {
		result := CheckSphereVsAabb(tt.s, tt.b)
		if result != tt.collided {
			t.Error("Invalid collision result.")
			t.Log(tt)
		}
	}
}
