package main

import "fmt"

func main10() {
	// 1. 슬라이스
	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
	a[1] = 10
	fmt.Println(a) // [1, 10, 3]출력

	// 2. make 사용
	s := make([]int, 5, 10)
	println(len(s), cap(s)) // len 5, cap 10

	// 3. 기본 값
	var s1 []int
	if s == nil {
		println("Nil Slice")
	}
	println(len(s1), cap(s1)) // 모두 0

	// 4. 부분 슬라이스
	s2 := []int{0, 1, 2, 3, 4, 5}
	s2 = s2[2:5]
	fmt.Println(s2) //2,3,4 출력

	// 5. 부분 슬라이스2
	s3 := []int{0, 1, 2, 3, 4, 5}
	s3 = s3[2:5]    // 2, 3, 4
	s3 = s3[1:]     // 3, 4
	fmt.Println(s3) // 3, 4 출력

	// 6. 슬라이스 추가, 병합(append), 복사(copy)
	s4 := []int{0, 1}
	// 하나 확장
	s4 = append(s4, 2) // 0, 1, 2
	// 복수 요소들 확장
	s4 = append(s4, 3, 4, 5) // 0,1,2,3,4,5
	fmt.Println(s4)

	/* 슬라이스에서 데이터를 추가할 때, 용량(capacity)이 부족하면 자동으로 현재 용량의 2배로 늘려준다. */

	// 7. 동적인 슬라이스
	// len=0, cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)
	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력

	// 8. 슬라이스 자체를 추가
	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}
	sliceB = append(sliceB, sliceC...)
	//sliceA = append(sliceA, 4, 5, 6)
	fmt.Println(sliceB) // [1 2 3 4 5 6] 출력

	// 9. 슬라이스 복사
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력

}
