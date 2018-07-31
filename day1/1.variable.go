package main

import "fmt"

func main2() {

	var a int = 1        // a := 1
	var b float32 = 11.1 // b := 11.0
	var c string = "Hi"  // c := "Hi"
	fmt.Println("1.", a, b, c)
	// 초기화 하지 않았을 경우, zero Value를 기본적으로 할당
	// int : 0, bool : false, string : ""(빈문자열)

	a = 2 // 재할당
	b = 12.0
	c = "Hello"
	fmt.Println("2.", a, b, c)

	var ( // 묶어서 선언 가능
		f int    = 1
		h string = "Hi"
	)

	fmt.Println("3.", f, h)

	var i, j, k int = 1, 2, 3 // 만약 변수를 선언하고 사용하지 않으면 오류를 뿜는다.

	fmt.Println("4.", i, j, k)
}
