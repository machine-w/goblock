package block

import (
	"fmt"
	"testing"

	"github.com/ma2ma/gola"
)

func TestSegmentXTriangle(t *testing.T) {
	a := &gola.Vector3{0, 0, 0}
	b := &gola.Vector3{10, 0, 0}
	b_un := &gola.Vector3{0.5, 0, 0}
	v0 := &gola.Vector3{1, 1, -1}
	v1 := &gola.Vector3{1, -1, -1}
	v2 := &gola.Vector3{1, 0, 1}
	expected := true
	actual, focus := SegmentXTriangle(Segment{a: a, b: b}, Triangle{v0: v0, v1: v1, v2: v2})
	actual_un, _ := SegmentXTriangle(Segment{a: a, b: b_un}, Triangle{v0: v0, v1: v1, v2: v2})

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	if !focus.Equals(&gola.Vector3{1, 0, 0}) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	if actual_un == expected {
		t.Errorf("Expected %v, got %v", expected, actual_un)
	}
}
func TestIntersectTriangle(t *testing.T) {
	orig := &gola.Vector3{0, 0, 0}
	orig_un := &gola.Vector3{10, 0, 0}
	orig_un2 := &gola.Vector3{0, 10, 10}
	dir := &gola.Vector3{1, 0, 0}
	v0 := &gola.Vector3{1, 1, -1}
	v1 := &gola.Vector3{1, -1, -1}
	v2 := &gola.Vector3{1, 0, 1}
	tri := Triangle{v0: v0, v1: v1, v2: v2}
	test := Ray{orig: orig, dir: dir}
	test1 := Ray{orig: orig_un, dir: dir}
	test2 := Ray{orig: orig_un2, dir: dir}
	expected := true
	actual, _, _, _ := IntersectTriangle(test, tri)
	actual_un, _, _, _ := IntersectTriangle(test1, tri)
	actual_un2, _, _, _ := IntersectTriangle(test2, tri)
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

func TestCube2Triangle(t *testing.T) {
	actual := Cube2Triangle(Cube{center: &gola.Vector3{1, 1, 1}, xSize: 2, ySize: 2, zSize: 2})
	for _, v := range actual {
		fmt.Println(v)
	}
}
func TestNewObject3D(t *testing.T) {
	vs := []float64{0, 0, 0, 2, 0, 0, 2, 2, 0, 0, 2, 0, 1, 1, 1}
	fs := []int64{0, 0, 1, 3, 0, 1, 2, 3, 0, 0, 1, 4, 0, 1, 2, 4, 0, 2, 3, 4, 0, 0, 3, 4}
	actual := NewObject3D(vs, fs)
	fmt.Println("points:")
	for _, v := range actual.points {
		fmt.Println(v)
	}
	fmt.Println("points_len:", actual.points_len)
	fmt.Println("maxx:", actual.MaxX)
	fmt.Println("maxy:", actual.MaxY)
	fmt.Println("maxZ:", actual.MaxZ)
	fmt.Println("faces:")
	for _, v := range Object3D2Faces(actual) {
		fmt.Println(v)
	}
	fmt.Println("segments:")
	for _, v := range Object3D2Segment(actual) {
		fmt.Println(v)
	}
}
func TestPointInsideObject3D(t *testing.T) {
	vs := []float64{0, 0, 0, 2, 0, 0, 2, 2, 0, 0, 2, 0, 1, 1, 1}
	fs := []int64{0, 0, 1, 3, 0, 1, 2, 3, 0, 0, 1, 4, 0, 1, 2, 4, 0, 2, 3, 4, 0, 0, 3, 4}
	obj := NewObject3D(vs, fs)
	actual := PointInsideObject(&gola.Vector3{1, 1, 0.5}, Object3D2Faces(obj))
	fmt.Println(actual)
}
func TestMakeOriBlock(t *testing.T) {
	actual := MakeOriBlock(100, 100, 100, -100, -100, -100, 3, 5, 5)
	fmt.Println(actual)
}
