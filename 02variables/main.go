package main


import "fmt"

const LoginToken string = "asgashgas"

//login token here is public as L is capital

func main(){
	var username string = "John"
	fmt.Println("Hello, ", username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println("Hello, ", isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	
	fmt.Println("Hello, ", smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	var website = "learncodeonline.in"
	fmt.Println(website)  


	//no var mode

	nuberOfuser := 30000
	fmt.Println(nuberOfuser)

}