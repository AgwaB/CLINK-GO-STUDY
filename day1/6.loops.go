package main

// 반복문

func main6() {

	// 초기값 ; 조건식 ; 증감식 형태, 괄호() 없음
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	// while문 처럼 쓸 때
	n := 1
	for n < 100 {
		n *= 2
		//if n > 90 {
		//   break
		//}
	}

	// 무한루프
	for {
		println("Infinite loop")
	}

	// for-range 문
	// for 인덱스, 요소값 := range 컬렉션
	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names {
		println(index, name)
	}
}
