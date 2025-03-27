package main

import "fmt"

func main(){
	defer fmt.Println("This is the first line")
	defer fmt.Println("This is the second line")
	defer fmt.Println("This is the third line")
	fmt.Println("Welcome to defer")
	// last in first out in defer format 
	// so first line got executed in the last 
}