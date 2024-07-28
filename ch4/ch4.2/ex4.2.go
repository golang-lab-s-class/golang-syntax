package main

import "fmt"

func main() {
	var mininumWage int = 10 //변수 minumWage 선언 및 초기화
	var workingHour int = 20 //변수 workingHour 선언 및 초기화

	//변수 income 선언 및 초기화
	var income int = mininumWage + workingHour

	//변수 minumWage,workingHour, income 출력
	fmt.Println(mininumWage, workingHour, income)
}
