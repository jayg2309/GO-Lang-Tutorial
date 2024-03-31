package main

import "fmt"

//jwtToken := 3000.00 // this will throw an error
//var jwtToken = 3000.00 // this will work

// constants
const LoginToken string = "jgongsofnd" // public token

// first letter of LoginToken is capital meaning it is public

func main() {
	//string
	var username string = "Jay"
	fmt.Println("Username :" + username)
	fmt.Printf("Variable is of type : %T \n", username)
	//boolean
	var IsLoggedIn bool = false
	fmt.Println(IsLoggedIn)
	fmt.Printf("Variable is of type : %T \n", IsLoggedIn)
	// small int
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)
	// small float
	var smallFloat float64 = 255.5443345465733234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable string                                 // if directly declared value is 0 by default
	fmt.Println(anotherVariable)                               // for string it is empty string
	fmt.Printf("Variable is of type : %T \n", anotherVariable) // hence no garbage values

	//implicit type
	var website = "google.com" // lexer automatically detects the type(string here)
	fmt.Println(website)
	//website = 3 // this will throw an error as website is string type

	// no var style
	numberOfUsers := 300000 // no use of var keyword
	// this is only allowed inside a function and not outside look at line 5
	// := called as walrus operator
	fmt.Println(numberOfUsers)

	//calling public constant
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
