package main

import "fmt"

func main(){
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = ""
	fruitList[3] = "peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "brinjal"}
	fmt.Println("Veg list is: ", vegList)
}	