package main

/*
//Rect - struct 정의
type Rect struct {
	width, height int
}

// value receiver
//Rect의 area() 메소드
func (r Rect) area() int { // (r Rect) 이 부분이 Receiver
	return r.width * r.height
}

func main12() {
	rect := Rect{10, 20} // Rect 구조체의 '객체' 선언
	area := rect.area()  // 객체 내부의 메서드 호출
	println(area)
}
*/
/* // 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area2()      //메서드 호출
	println(rect.width, area) // 11 220 출력
}
*/
