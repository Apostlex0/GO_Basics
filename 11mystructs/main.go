package main

import "fmt"

func main(){
	fmt.Println("Welcome to structs")

	Sachin := User{"Sachin", "BzX3M@example.com", true, 16}
	fmt.Println(Sachin)
	fmt.Printf("Sachin details are: %+v\n", Sachin) 
	fmt.Printf("Name is %v and Email is %v.\n", Sachin.Name, Sachin.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}