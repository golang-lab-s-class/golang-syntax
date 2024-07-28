package main

import "fmt"

var g int = 10 // 패키지 전역 변수 선언

func main() {
	var m int = 20 //지역 변수 선언

	{
		var s int = 50 //지역 변수 선언
		fmt.Println(m, s, g)
	} //지역 변수 s는 사라짐

	m = s + 20 //에러 발생
} //메인 함수 종료
