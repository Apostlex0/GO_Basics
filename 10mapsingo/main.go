package main

import "fmt"

func main(){
	languages := make(map[string]string)
	// firsty string is key and the next one is value

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println(languages)
	fmt.Println(languages["GO"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loops are interesting in golang 

	for key,value := range languages{
		fmt.Printf("For key %v, value is %v\n", key, value)
	} 
	for _,value := range languages{
		fmt.Printf("For key v, value is %v\n", value)
	} 
}