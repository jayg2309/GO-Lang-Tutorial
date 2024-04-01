// Handling strings and converting it to numbers in Go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my pizza app:")
	fmt.Println("Please rate the pizza between 1 and 5:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // \n implies keep on reading till you get \n
	//fmt.Println(input)
	fmt.Println("Pizza rating : " + input)

	//let's say we wanna add one to our rating

	//numRating = input + 1 //error
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // string.TrimSpace removes leading and trailing spaces
	// hence if user enters 3 and then space, it will remove space and convert 3 to float64 3.0 and store in numRating

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("Added 1 to your rating :", numRating+1) // now no error as numRating
	}

}
