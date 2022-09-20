package main

import (
	"fmt"
)

func main () {
	var number1 int
	var number2 int
	var choiseUser int
	// number1=5
	// number2=2
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&number2)
	fmt.Printf("Here are the two integers that we will be working on, %d & %d", number1,number2)
	fmt.Println("\n")
	fmt.Println("Great, now enter 1 to perform Addition, 2 to Multiplication and 3 to Division ")
	fmt.Scanln(&choiseUser)
	if choiseUser == 1 {
		performAdd(number1,number2)
	}else if choiseUser == 2 {
		performMultiplication(number1,number2)
	}else {
		performDivision(number1,number2)
	}
}

 func performAdd(number1 int,number2 int){
 	result:=(number1+number2)
 	fmt.Printf("The result of addition is %d",result)
	fmt.Println("\n")
 }
func performMultiplication(number1 int,number2 int){
	result:=(number1*number2)
	fmt.Printf("The result of Multiplication is %d",result)
	fmt.Println("\n")
 }
func performDivision(number1 int,number2 int){
	result:=(number1/number2)
	fmt.Printf("The result of Division is %d",result)
  	fmt.Println("\n")
}	
