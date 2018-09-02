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
type Cube struct {
	center *gola.Vector3
	xSize  float64
	ySize  float64
	zSize  float64
}
type Object3D struct {
	vertices []float64
	faces    []int64
}

func Cube2Triangle(c Cube) [12]Triangle {
	tris := [12]Triangle{}
	xDet := c.xSize / 2
	yDet := c.ySize / 2
	zDet := c.zSize / 2
	points := [8]*gola.Vector3{}
	points[0] = &Vector3{c.center[0] - xDet, c.center[1] - yDet, c.center[2] - zDet}
	points[1] = &Vector3{c.center[0] + xDet, c.center[1] - yDet, c.center[2] - zDet}
	points[2] = &Vector3{c.center[0] + xDet, c.center[1] + yDet, c.center[2] - zDet}
	points[3] = &Vector3{c.center[0] + xDet, c.center[1] - yDet, c.center[2] - zDet}
	points[4] = &Vector3{c.center[0] - xDet, c.center[1] - yDet, c.center[2] + zDet}
	points[5] = &Vector3{c.center[0] + xDet, c.center[1] - yDet, c.center[2] + zDet}
	points[6] = &Vector3{c.center[0] + xDet, c.center[1] + yDet, c.center[2] + zDet}
	points[7] = &Vector3{c.center[0] + xDet, c.center[1] - yDet, c.center[2] + zDet}
	tris[0] = Triangle{v0: points[0], v1: points[1], v2: points[2]}
	tris[1] = Triangle{v0: points[0], v1: points[2], v2: points[3]}
	tris[2] = Triangle{v0: points[4], v1: points[5], v2: points[6]}
	tris[3] = Triangle{v0: points[4], v1: points[6], v2: points[7]}
	tris[4] = Triangle{v0: points[0], v1: points[4], v2: points[5]}
	tris[5] = Triangle{v0: points[0], v1: points[5], v2: points[1]}
	tris[6] = Triangle{v0: points[0], v1: points[3], v2: points[7]}
	tris[7] = Triangle{v0: points[0], v1: points[7], v2: points[4]}
	tris[8] = Triangle{v0: points[1], v1: points[2], v2: points[6]}
	tris[9] = Triangle{v0: points[1], v1: points[6], v2: points[5]}
	tris[10] = Triangle{v0: points[2], v1: points[6], v2: points[7]}
	tris[11] = Triangle{v0: points[2], v1: points[7], v2: points[3]}
	return tris
}
func Object3D2Segment(o Object3D) []*Segment {
	seg := []*Segment{}
	points := []*gola.Vector3{}
	strDict := make(map[string]int)
	for i := 0; i <= len(o.vertices)-3; i += 3 {
		points.append(points, &Vector3{o.vertices[i], o.vertices[i+1], o.vertices[i+2]})
	}
	lenFace := len(points)
	for i := 0; i <= len(o.faces)-4; i += 4 {
		if o.faces[i+1] < lenFace && o.faces[i+2] < lenFace && o.faces[i+3] < lenFace {
			k1 := fmt.Sprintf("%d-%d", o.faces[i+1], o.faces[i+2])
			k2 := fmt.Sprintf("%d-%d", o.faces[i+2], o.faces[i+1])
			if _, ok := strDict[k1]; !ok {
				strDict[k1] = 1
				strDict[k2] = 1
				seg.append(seg, &Segment{a: points[o.faces[i+1]], b: points[o.faces[i+2]]})
			}
			k3 := fmt.Sprintf("%d-%d", o.faces[i+3], o.faces[i+2])
			k4 := fmt.Sprintf("%d-%d", o.faces[i+2], o.faces[i+3])
			if _, ok := strDict[k3]; !ok {
				strDict[k3] = 1
				strDict[k4] = 1
				seg.append(seg, &Segment{a: points[o.faces[i+3]], b: points[o.faces[i+2]]})
			}
			k5 := fmt.Sprintf("%d-%d", o.faces[i+3], o.faces[i+1])
			k6 := fmt.Sprintf("%d-%d", o.faces[i+1], o.faces[i+3])
			if _, ok := strDict[k5]; !ok {
				strDict[k5] = 1
				strDict[k6] = 1
				seg.append(seg, &Segment{a: points[o.faces[i+3]], b: points[o.faces[i+1]]})
			}
		}
	}
	return seg
}
func Object3D2Faces(o Object3D) []Triangle {
	tris := []Triangle{}
	points := []*gola.Vector3{}
	for i := 0; i <= len(o.vertices)-3; i += 3 {
		points.append(points, &Vector3{o.vertices[i], o.vertices[i+1], o.vertices[i+2]})
	}
	lenFace := len(points)
	for i := 0; i <= len(o.faces)-4; i += 4 {
		if o.faces[i+1] < lenFace && o.faces[i+2] < lenFace && o.faces[i+3] < lenFace {
			tris.append(tris, Triangle{v0: points[o.faces[i+1]], v0: points[o.faces[i+2]], v0: points[o.faces[i+3]]})
		}
	}
	return tris
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
