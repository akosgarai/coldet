package coldet

import (
	"testing"
)

var (
	aabb = AABB{[3]float32{1, 2, 3}, 4, 5, 6}
)

func TestX(t *testing.T) {
	if aabb.X() != 1 {
		t.Error("Invalid x coordinate.")
	}
}
func TestY(t *testing.T) {
	if aabb.Y() != 2 {
		t.Error("Invalid y coordinate.")
	}
}
func TestZ(t *testing.T) {
	if aabb.Z() != 3 {
		t.Error("Invalid z coordinate.")
	}
}
func TestWidth(t *testing.T) {
	if aabb.Width() != 4 {
		t.Error("Invalid width.")
	}
}
func TestLength(t *testing.T) {
	if aabb.Length() != 5 {
		t.Error("Invalid length.")
	}
}
func TestHeight(t *testing.T) {
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
