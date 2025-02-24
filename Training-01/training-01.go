package main

import ("fmt")

func main() {

	/*
 var firstName string = "paul"
	var lastName string = "McCartney"
	// firstName := "Paul"
	// lastName := "McCartney"

	currentYear := 2025
	myBirthYear := 1991
	currentYear = currentYear + 1
	age := currentYear - myBirthYear


	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

*/
	luckyNumber := 10
	fmt.Println(luckyNumber)

	moreLuckyNumber := luckyNumber + 5
	fmt.Println(moreLuckyNumber)

	var newNum float64 = float64(luckyNumber) / 3 
	

	fmt.Println(newNum)


	var firstRune rune = '€'

	fmt.Println(firstRune)
	fmt.Println(string (firstRune))

	var firstByte byte = 'a'

	fmt.Println(firstByte)
	fmt.Println(string(firstByte))


	firstName := "Malik"
	lastName := "Tabish"
	age:= 33

	fmt.Println(firstName, lastName)
	// fullName := "Malik Tabish"
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	
	// fmt.Println("Hi, my full name is " + fullName + " and i am " + age + " years old") //  so here it couldnt work so we use another 
	// version of command key. 

	//fmt.println("9" + 1) This operation dosent work in GO ; this is a mix operation of String and int.

	fmt.Printf("Hi, I am %v and I am %v (Type: %T) years old. ", fullName, age, age) // special identifier ( %v ) 



}





/*
											import :
	The standrd libary is a "built-in module" full of core 
	packages and functionalities.
	* we have the same concept ("Standard Libary") in Python
	or other language. 

*/