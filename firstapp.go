package main

import "fmt"

func main() {
	// var greetingText string
	// greetingText = "Hello World"

	// var greetingText string = "Hello World"

	greetingText := "Hello World"

	fmt.Println(greetingText)

	var firstFloat64 float64 = 1.123123123123123123123
	var firstFloat32 float32 = 1.123123123123123

	fmt.Println(firstFloat64)
	fmt.Println(firstFloat32)

	var firstRune rune = '&'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 'a'
	fmt.Println(firstByte)
}
