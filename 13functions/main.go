package main

import "fmt"

func main(){
	fmt.Println("Welcome to functions")
	greeter()
	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(1,2,3,4,5)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println("Pro myMessage is: ", myMessage)
}

func adder(valOne int, valTwo int) int{
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total :=0
	for _,value := range values{
		total += value
	}
	return total, "proadder message"
}

func greeter(){
	fmt.Println("Welcome to golang")
}

//we are not allowed to write a function inside another fucntion