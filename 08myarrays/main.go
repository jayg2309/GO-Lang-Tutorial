package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Go lang")
	//declare
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is :", fruitList)                // [Apple Orange  Banana]
	fmt.Println("Length of Fruit list is :", len(fruitList)) //4 // it is 4 because we have declared 4 elements in the array

	// directly putting values

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list is :", vegList)                // [potato beans mushroom]
	fmt.Println("Length of Veg list is :", len(vegList)) //3

}
