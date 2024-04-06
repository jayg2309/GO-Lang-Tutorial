// Use slices for most cases when you need a collection of
// elements as they are more flexible and efficient for memory usage.
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	//syntax
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList) //Type is []string
	//appending
	fruitList = append(fruitList, "Banana", "Mango")
	fmt.Println(fruitList) // [Apple Tomato Peach Banana Mango]

	fruitList = append(fruitList[1:3]) // range are always non inclusive
	fmt.Println(fruitList)             // [Tomato Peach]

	//using make keyword to create a slice
	highScore := make([]int, 4)
	highScore[0] = 345
	highScore[1] = 978
	highScore[2] = 545
	highScore[3] = 767
	fmt.Println(highScore) // [345 978 545 767]
	//highScore[4] = 111     // error here
	//appending
	//entire memory allocation happens again and hence it stores it without errors
	highScore = append(highScore, 222, 333, 666)
	fmt.Println(highScore) // [345 978 545 767 222 333 666] no error here
	// sort package
	sort.Ints(highScore)
	fmt.Println(highScore)                     //sorts  // [222 333 345 545 666 767 978]
	fmt.Println(sort.IntsAreSorted(highScore)) // true

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	//	courses = append(courses[:index], courses[index+1:])    // gives error
	courses = append(courses[:index], courses[index+1:]...) // gives correct output
	fmt.Println(courses)                                    //[reactjs javascript python ruby]
}
