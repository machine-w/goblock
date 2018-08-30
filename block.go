package block

import (
	"fmt"

	"github.com/ma2ma/gola"
)

var Epsilon float64 = 0.0001

type Segment struct {
	a *gola.Vector3
	b *gola.Vector3
}
type Ray struct {
	orig *gola.Vector3
	dir  *gola.Vector3
}
type Triangle struct {
	v0 *gola.Vector3
	v1 *gola.Vector3
	v2 *gola.Vector3
}

func SegmentXTriangle(seg Segment, tri Triangle) (intersect bool, focus *gola.Vector3) {
	dir := seg.b.NewSub(seg.a)
	segmentLength := dir.Length()
	dir.Normalize()

	fmt.Println(seg.a, seg.b)
	intersect, t, _, _ := IntersectTriangle(Ray{orig: seg.a, dir: dir}, tri)
	if t > segmentLength {
		intersect = false
	}
	focus = seg.a.NewAdd(dir.Scale(t))
	return
}

func IntersectTriangle(ray Ray, tri Triangle) (intersect bool, t, u, v float64) {
	intersect = false
	t, u, v = 0, 0, 0
	e1 := tri.v1.NewSub(tri.v0)
	e2 := tri.v2.NewSub(tri.v0)
	p := ray.dir.NewCross(e2)
	det := e1.Dot(p)

	var vt *gola.Vector3
	if det > 0 {
		vt = ray.orig.NewSub(tri.v0)
	} else {
		vt = tri.v0.NewSub(ray.orig)
		det = -det
	}
	// fmt.Println(det)
	// fmt.Println(vt)
	// fmt.Println(p)
	if det < Epsilon {
		return
	}
	u = vt.Dot(p)
	// fmt.Println(u)
	if u < 0.0 || u > det {
		return
	}
	q := vt.NewCross(e1)
	v = ray.dir.Dot(q)
	// fmt.Println(v)
	if v < 0.0 || u+v > det {
		return
	}
	t = e2.Dot(q)
	fmt.Println(t)
	if t < 0.0 {
		return
	}
	finvdet := 1.0 / det
	t *= finvdet
	u *= finvdet
	v *= finvdet
	fmt.Println(t)
	intersect = true
	return
}
