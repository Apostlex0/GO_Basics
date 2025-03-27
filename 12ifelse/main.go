package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Welcome to if else")
	var result string

	loginCount := 10
	if loginCount < 10 {
		fmt.Println("Regular user")
	} else if loginCount > 10 {
		fmt.Println("Watch out")
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println("Result is: ", result)

	// switch now 

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice number is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1. You can open the app")
	case 2:
		fmt.Println("You got 2. You can close the app")
	case 3:
		fmt.Println("You got 3. You can delete the app")
		fallthrough
	case 4:
		fmt.Println("You got 4. You can uninstall the app")
	case 5:
		fmt.Println("You got 5. You can throw the phone")	
	case 6:
		fmt.Println("You got 6. You can run the app")	
	default:
		fmt.Println("You can do whatever you want")
	}

	//loop break and continue and goto 

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for d:=0; d<len(days); d++{
	// 	if days[d] == "Friday"{
	// 		fmt.Println("Go to temple")
	// 		break
	// 	}
	// 	fmt.Println(days[d])
	// }

	for index, day := range days {
		if day == "Friday"{
			fmt.Println("Go to temple")
			break
		}
		fmt.Println(index, day)
	}
	// here index is just key in ntegers if we wnteed to loop on maps with key as a string we will just use that and write key or if we dont then we can just use _ instead of index or key 

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	lco:
	    fmt .Println("Jumping to LCO")

		//here lco is a lable can use any lable and to reach a lable we use goto
}