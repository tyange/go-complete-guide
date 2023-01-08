package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

//func main() {
//	var productNames [4]string
//	productNames = [4]string{"a Book"}
//	prices := [4]float64{10.99, 40.0, 12.99, 20.0}
//	fmt.Println(prices)
//
//	productNames[2] = "a Carpet"
//
//	fmt.Println(productNames)
//
//	fmt.Println(prices[2])
//
//	featuredPrices := prices[:3]
//	featuredPrices[0] = 199.99
//	highlightedPrices := featuredPrices[:1]
//	fmt.Println(highlightedPrices)
//	fmt.Println(prices)
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//
//	highlightedPrices = highlightedPrices[:3]
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//}
