package main

// 조건문

func main5() {
	var k int = 1
	var max int = 3
	var a interface{} = 10

	// if

	if k == 1 { // 1. 시작 중괄호는 if문과 같은 라인에 두어야 한다.
		// something	  // 2. if 문의 조건식은 반드시 Boolean 식으로 표현되어야 한다. (1 or 0 같은 숫자 안됨)
	} else if k == 2 {
		// something
	} else {
		// something
	}

	if some := k * 2; some < max { // if 문에서 조건식 사용 전에 간단한 문장을 함께 실행시킬 수 있다. 'some' 변수는 if문 블럭 안에서만 사용 가능.
		//something
	}
	// if문을 빠져나온 여기서 some을 사용하면 에러

	// switch
	switch k { // golang 에서 switch는 컴파일러가 자동으로 break를 추가해 준다.
	case 1:
		// 1 something
	case 2:
		// 2 something
	case 3:
		// 3 something
	default:
		// 1, 2, 3 을 만족하지 않을 경우 default 실행
	}

	switch some := k * 2; some { // if와 마찬가지고 간단한 문장 실행 가능
	//...
	}

	switch a.(type) { // type 체크 가능
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}
}

func grade(score int) {
	switch {
	case score >= 90:
		println("A")
		//fallthrough -> 만약에 break 없이 계속 이어나가고 싶다면 fallthrough 추가
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}
