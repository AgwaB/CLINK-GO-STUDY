package main

import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

//Rect 정의
type Rect struct {
	width, height float64
}

//Circle 정의
type Circle struct {
	radius float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64      { return r.width * r.height }
func (r Rect) perimeter() float64 { return 2 * (r.width + r.height) }

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64      { return math.Pi * c.radius * c.radius }
func (c Circle) perimeter() float64 { return 2 * math.Pi * c.radius }

func main13() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}

func showArea(shapes ...Shape) { // 파라미터로도 사용 가능
	for _, s := range shapes {
		a := s.area() //인터페이스 메서드 호출
		println(a)
	}
}
