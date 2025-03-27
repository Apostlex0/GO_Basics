package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	welcome := "welcome to usr input "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok || err err

	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

}