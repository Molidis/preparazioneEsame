package main

import "fmt"

type Point struct {
	x, y float64
}

type Rectangle struct {
	pLL, pUr Point
}

func NewPoint(x, y float64) Point {
	return Point(x, y)
}

func NewRectangle(p1, p2 Point) Rectangle {
	var r Rectangle
	if p1.x < p2.x {
		if p1.y < p2.y {
			r.pLL = p1
			r.pUr = p2
		} else {
			r.pLL = p2
			r.pUr = p1
		}
	} else {
		if p1.y < p2.y {
			r.pLL = p2
			r.pUr = p1
		} else {
			r.pLL = p1
			r.pUr = p2
		}
	}
	return r
}

func Rotate(r *Rectangle, dir byte) {
	if dir == 'R' {
		r.pLL.x, r.pLL.y = r.pLL.y, r.pLL.x
		r.pUr.x, r.pUr.y = r.pUr.y, r.pUr.x
	} else if dir == 'L' {
		r.pLL.x, r.pLL.y = r.pLL.y, r.pLL.x
		r.pUr.x, r.pUr.y = r.pUr.y, r.pUr.x
	} else {
		return
	}
}

func (r Rectangle) String() string {
	return fmt.Sprintf("(%v %v)", r.pLL, r.pUr)
}
