package block

import (
	"github.com/ma2ma/gola"
)

var Epsilon float64 = 0.0001

func IntersectTriangle(orig, dir, v0, v1, v2 *gola.Vector3) (intersect bool, t, u, v float64) {
	intersect = false
	t, u, v = 0, 0, 0
	e1 := v1.NewSub(v0)
	e2 := v2.NewSub(v0)
	p := dir.NewCross(e2)
	det := e1.Dot(p)

	var vt *gola.Vector3
	if det > 0 {
		vt = orig.NewSub(v0)
	} else {
		vt = v0.NewSub(orig)
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
	v = dir.Dot(q)
	// fmt.Println(v)
	if v < 0.0 || u+v > det {
		return
	}
	t = e2.Dot(q)
	// fmt.Println(t)
	if t < 0.0 {
		return
	}
	finvdet := 1.0 / det
	t *= finvdet
	u *= finvdet
	v *= finvdet
	intersect = true
	return
}
