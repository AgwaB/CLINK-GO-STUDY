package main

/* Pass by Value
func main() {
	msg := "Hello"
	value(msg) // 변수 msg의 "Hello"라는 값을 복사해서 넘김
}

func value(msg string) {
	println(msg)
}
*/

/*  Pass by Reference
func main() {
	msg := "Hello"
	Ref(&msg)
	println(msg) //변경된 메시지 출력
}

func Ref(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}
*/

func main8() {
	var ptr *int             // 포인터형 변수 선언 (*을 붙임) 초기화 값 : nil
	var ptr2 *int = new(int) // new 함수로 메모리를 할당
	/* Go는 가비지 컬렉션을 지원하므로, new 할당 후, 따로 해제하지 않아도 됨 */
	println(ptr)
	println(ptr2) // 메모리 주소, 실행 할 때 마다 값이 달라짐

	var num int = 1
	var numPtr *int = new(int)
	numPtr = &num
	// 35, 36 line은 var numPtr *int = &num 이렇게도 표현가능

	println(&num)    // 변수 num의 주소
	println(numPtr)  // 변수 num의 주소를 가리킴, (변수 num의 주소)
	println(*numPtr) // 변수 num의 주소에 할당된 값
	println(&numPtr) // 포인터 변수 numPtr의 주소
}
