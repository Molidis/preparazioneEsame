package main

import "fmt"

type Point struct {
	x, y float64
}

type Rectangle struct {
	pLL, pUR Point
}

func NewPoint(x, y float64) Point {
	var p Point
	p.x = x
	p.y = y
	return p
}

func NewRectangle(p1, p2 Point) Rectangle {
	var r Rectangle
	if p1.x < p2.x {
		if p1.y < p2.y {
			r.pLL = p1
			r.pUR = p2
		} else {
			r.pLL = p2
			r.pUR = p1
		}
	} else {
		if p1.y < p2.y {
			r.pLL = p2
			r.pUR = p1
		} else {
			r.pLL = p1
			r.pUR = p2
		}
	}
	return r
}

func Rotate(r *Rectangle, dir byte) {
	if dir == 'R' {
		r.pLL.x, r.pLL.y = r.pLL.y, r.pLL.x
		r.pUR.x, r.pUR.y = r.pUR.y, r.pUR.x
	} else if dir == 'L' {
		r.pLL.x, r.pLL.y = r.pLL.y, r.pLL.x
		r.pUR.x, r.pUR.y = r.pUR.y, r.pUR.x
	} else {
		return
	}
}

func (r Rectangle) String() string {
	return fmt.Sprintf("(%v %v)", r.pLL, r.pUR)
}

func main() {
	var p1, p2 Point
	var r Rectangle
	p1 = NewPoint(10.5, 6.9)
	p2 = NewPoint(4, 0)
	r = NewRectangle(p1, p2)
	fmt.Println(r)
	Rotate(&r, 'R')
	fmt.Println(r)
}
