package main

func main1() {

	const a int = 1 // const(상수) : 고정값, 재할당 불가능
	const b float32 = 11.1
	const c string = "Hi"

	const d = 10 // 이런식으로 Go에서 자동으로 추론하는 것도 가능
	const e = "I'm string"

	const (
		kim  = "kim"
		sung = "sung"
		jae  = "jae"
	)

	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)
}
