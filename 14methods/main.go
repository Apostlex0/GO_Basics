package main

import "fmt"

func main(){
	fmt.Println("Welcome to structs")

	Sachin := User{"Sachin", "BzX3M@example.com", true, 16}
	fmt.Println(Sachin)
	fmt.Printf("Sachin details are: %+v\n", Sachin) 
	fmt.Printf("Name is %v and Email is %v.\n", Sachin.Name, Sachin.Email)
	Sachin.GetStatus()
	Sachin.NewMail()
	fmt.Printf("Name is %v and Email is %v.\n", Sachin.Name, Sachin.Email)
	// when passing onto a function , a copy of it is passed along hence no chagne to actual value so to change the actual value we need to pass it as pointer
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}
// a method is just a function with a struct called into it 

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
