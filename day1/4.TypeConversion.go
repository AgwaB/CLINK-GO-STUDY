package main

func main4() {
	//앞서 언급한 것 처럼, golang은 명시적인 형변환을 해주어야 한다.
	var i int = 1
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
