package main

import (
	"fmt"
	"math"
)

type Punto struct {
	x, y float64
}

type Triangolo struct {
	P1, P2, P3 Punto
}

func lunghezzeLati(A, B, C Punto) [3]float64 {
	var l [3]float64
	l[0] = math.Sqrt((A.x-B.x)*(A.x-B.x) + (A.y-B.y)*(A.y-B.y))
	l[1] = math.Sqrt((A.x-C.x)*(A.x-C.x) + (A.y-C.y)*(A.y-C.y))
	l[2] = math.Sqrt((B.x-C.x)*(B.x-C.x) + (B.y-C.y)*(B.y-C.y))
	return l
}

func newTriangolo(A, B, C Punto) (Triangolo, error) {
	var t Triangolo

	l := lunghezzeLati(A, B, C)

	if l[0] >= l[1]+l[2] || l[1] >= l[0]+l[2] || l[2] >= l[0]+l[1] {
		return t, fmt.Errorf("non è un triangolo")
	}

	t.P1 = A
	t.P2 = B
	t.P3 = C

	return t, nil
}

func tipoTriangolo(tri Triangolo) int {
	l := lunghezzeLati(tri.P1, tri.P2, tri.P3)

	if (l[0] >= l[1] && l[0] <= l[1]+1e-6) || (l[1] >= l[2] && l[1] <= l[2]+1e-6) || (l[2] >= l[0] && l[2] <= l[0]+1e-6) {
		return 3
	}

	if l[0] == l[1] || l[1] == l[2] || l[0] == l[2] {
		return 2
	}

	return 0
}

func main() {
	var P1, P2, P3 Punto

	fmt.Scanf("%f %f\n%f %f\n%f %f", &P1.x, &P1.y, &P2.x, &P2.y, &P3.x, &P3.y)

	triangolo, err := newTriangolo(P1, P2, P3)
	if err != nil {
		fmt.Println(lunghezzeLati(P2, P1, P3))
		fmt.Println(err)
	} else {
		fmt.Println(lunghezzeLati(P2, P1, P3))
		fmt.Println("triangolo", triangolo)
		fmt.Println("tipo", tipoTriangolo(triangolo))
	}
}
