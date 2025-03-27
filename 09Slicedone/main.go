package main

import (
	"fmt"
	"sort"
)


func main(){
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	//we need to initia,ise them when using this syntax 
	fmt.Printf("Fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Banana")
	fmt.Println("Fruit list is: ", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("Fruit list is: ", fruitList)
	  
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 777)
	sort.Ints(highScores)

	fmt.Println("High scores: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"react", "angular", "vue", "svelte"}
	fmt.Println("List of courses: ", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) 
	fmt.Println(courses)

}