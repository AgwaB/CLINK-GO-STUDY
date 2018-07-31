package main

import "fmt"

func main3() {
	rawLiteral := `1.CLINK \n CLINK "~!@#$%^&*()_+,./[]{}\t ` // ``로 둘러 쌓여있으면 Raw String Literal 이라고 부른다. 문자열을 별도로 해석하지 않는다.
	interLiteral := "2.CLINK \n CLINK"                        // 일반적인 string 선언. 문자열을 해석한다.

	fmt.Println(rawLiteral, "\n", interLiteral)
}
