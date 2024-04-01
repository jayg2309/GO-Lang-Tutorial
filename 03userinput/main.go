// taking user input in golang
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//create a reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for your pizza: ")

	//comma ok syntax or comma error syntax
	//	input, err := reader.ReadString('\n')
	// if everything is ok, then the first value is stored in input and the second value is stored in _.
	//if anything goes wrong, then the error is stored in the second value and the first value is stored in _.
	//The _ is used after the input variable in line 24 to ignore the second return value of reader.ReadString('\n'). In this case, the second return value represents any error that may occur during the reading process.
	// By using _, we are indicating that we are intentionally ignoring the error and not handling it in our code.
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating :", input)
	fmt.Printf("Type of this rating is %T ", input)
}
