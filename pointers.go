package main

import "fmt"

func main() {
	age := 32

	fmt.Println(age)

	myAge := &age
	// 아스터리크(*)는 myAge 변수가 int 타입의 데이터를 가진 메모리 주소의 '값'을 가리킬 것을 의미함.
	// var myAge *int
	// myAge = &age

	fmt.Println(*myAge)

	*myAge = 33

	fmt.Println(*myAge)
	fmt.Println(age)
}
