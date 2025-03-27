package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcome to files in go")
	content := "this needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myLcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of the file is: ", length)
	defer file.Close()
	readFile("./myLcogofile.txt")

}

func readFile(filname string){
	databyte, err := ioutil.ReadFile(filname)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text in the file is:\n ", string(databyte))
}