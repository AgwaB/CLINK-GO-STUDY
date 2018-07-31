package main

import "fmt"

// struct 정의
type person struct {
	name string
	age  int
}

func main() {
	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	// 여러가지 초기화 방법
	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50} // 만약 지정 안 한 필드는 zero value

	p3 := new(person) // new로 초기화 시 모든 필드가 zero value
	p3.name = "Lee"   // p가 포인터라도 . 을 사용한다
}
