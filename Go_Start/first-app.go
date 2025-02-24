package main

import "fmt"

func main() {
	//var greetingText string
	// greetingText = "Hello world"
	// var greetingText string = "Hello world"
	greetingText := "Hello world" // this is the common way of declaring a variable
	luckyNumber := 7              // this is the common way of declaring a variable
	evenMoreLuckyNumber := luckyNumber + 4

	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)
	var newNumber float64 = float64(luckyNumber) / 3
	evenMoreLuckyNumber = luckyNumber / 2
	fmt.Println(evenMoreLuckyNumber)
	fmt.Println(newNumber)

	var firstRune rune = '€'

	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	firstName := "Malik"
	lastName := "TABISH"
	fullName := firstName + " " + lastName
	age := 33

	fmt.Printf("Hi, I'm ")

}
