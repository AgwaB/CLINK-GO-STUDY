/* 함수의 기본 형식
func + 함수명(파라미터) + return type{
	함수 내용
}
파라미터는 0개 이상
*/
package main

import "fmt"

func main7() {
	msg := "Hello"
	say(msg)

	msg2 := say2(msg)
	fmt.Println(msg2)

	msg3, msg4 := say3(msg)
	fmt.Println(msg3, msg4)

	total, count := sum(1, 2, 3, 4, 5)
	fmt.Println(total, count)

	total, count = sum2(1, 2, 3, 4, 5)
	fmt.Println(total, count)

}

// return 없음
func say(msg string) {
	println(msg)
}

// return 있음
func say2(msg string) string {
	msg += "CLINK"
	return msg
}

// return 복수개
func say3(msg string) (string, string) {
	msg += "CLINK"
	msg2 := "Success"
	return msg, msg2
}

// 파라미터 복수개
func sum(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

// Named Return Parameter 이용. 파라미터에 이름이 있다.
// total은 아래 코드에서 초기화가 안되어 있으니 zero value를 갖는다.
func sum2(nums ...int) (count int, total int) {
	/*	물론 이런식으로 초기화 가능
		total := 0
	*/
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
