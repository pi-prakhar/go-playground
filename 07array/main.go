package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to array in golangs")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", len(vegList))

	var fruitListSlice = []string{"Apple", "Tomato"}
	fmt.Printf("Type of fruitlist is %T \n", fruitListSlice)
	fruitListSlice = append(fruitListSlice, "Mango", "Banana")
	fmt.Println(fruitListSlice)

	fruitListSlice = append(fruitListSlice[1:3])
	fmt.Println(fruitListSlice)

	highscores := make([]int, 4)
	highscores[0] = 123
	highscores[1] = 456
	highscores[2] = 678
	highscores[3] = 890
	//highscores[4] = 909 Error: index out of bound

	highscores = append(highscores, 555, 666, 999) //now new memory is allocated and size increases

	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

}
