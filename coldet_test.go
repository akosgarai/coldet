package coldet

import (
	"testing"
)

var (
	aabb = AABB{[3]float32{1, 2, 3}, 4, 5, 6}
	p    = Point{[3]float32{1, 2, 3}}
)

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
func TestCheckPointVsAabb(t *testing.T) {
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
		result := CheckPointVsAabb(tt.p, tt.b)
		if result != tt.in {
			t.Error("Invalid collision result.")
		}
	}
}
