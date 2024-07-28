package main

import "fmt"

func main() {
	var a int = 10                    //변수 a선언
	var msg string = "Hello Variable" //변수 msg 선언

	a = 20               //변수 a의 값 변경
	msg = "Good Morning" //변수 msg의 값 변경
	fmt.Println(msg, a)  //msg와 a값 출력
}
