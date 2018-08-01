package main

import "fmt"

func main() {
	/*
		var x interface{} // 어떠한 타입도 담을 수 있다.
		x = 1             // x : 1
		x = "Tom"         // x : "Tom"

		printIt(x)
	*/
	// type assertion
	var a interface{} = 1

	i := a       // a와 i 는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1, a.(string) 하면 런타임 에러 발생

	println(i) // 포인터주소 출력
	println(j) // 1 출력
}

func printIt(v interface{}) { // 파라미터의 type에 상관없이 다 들어갈 수 있다.
	fmt.Println(v) //Tom
}
