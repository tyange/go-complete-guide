package main

import (
	"fmt"
	"github.com/tyange/firstapp/greeting"
)

func main() {
	// var greetingText string
	// greetingText = "Hello World"

	// var greetingText string = "Hello World"

	// greetingText := "Hello World"

	fmt.Println(greeting.GreetingText)

	var firstFloat64 float64 = 1.123123123123123123123
	var firstFloat32 float32 = 1.123123123123123

	fmt.Println(firstFloat64)
	fmt.Println(firstFloat32)

	var firstRune rune = '&'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 'a'
	fmt.Println(firstByte)

	firstName := "Yu"
	lastName := "Tyange"
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 32

	fmt.Printf("Hi, I am %v and I am %v (Type: %T) years old.\n", fullName, age, age)
	// fmt.Println("9" +  1)
}
