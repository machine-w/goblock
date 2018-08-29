package block

import (
	"testing"

	"github.com/ma2ma/gola"
)

func TestIntersectTriangle(t *testing.T) {
	orig := &gola.Vector3{0, 0, 0}
	orig_un := &gola.Vector3{10, 0, 0}
	orig_un2 := &gola.Vector3{0, 10, 10}
	dir := &gola.Vector3{1, 0, 0}
	v0 := &gola.Vector3{1, 1, -1}
	v1 := &gola.Vector3{1, -1, -1}
	v2 := &gola.Vector3{1, 0, 1}
	expected := true
	actual, _, _, _ := IntersectTriangle(orig, dir, v0, v1, v2)
	actual_un, _, _, _ := IntersectTriangle(orig_un, dir, v0, v1, v2)
	actual_un2, _, _, _ := IntersectTriangle(orig_un2, dir, v0, v1, v2)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	if actual_un == expected {
		t.Errorf("Expected %v, got %v", expected, actual_un)
	}
	if actual_un2 == expected {
		t.Errorf("Expected %v, got %v", expected, actual_un2)
	}
}
