// Maps - Key Value Pair Data
package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	languages := make(map[string]string) // key is of type string and value is of type string
	// it can be of any datatype
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages :", languages)     // List of languages : map[JS:Javascript RB:Ruby PY:Python]
	fmt.Println("JS is short for :", languages["JS"]) // JS is short for : Javascript

	delete(languages, "RB")
	fmt.Println(languages) // map[JS:Javascript PY:Python]

	// loop in maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
		//For key JS, value is Javascript
		//For key PY, value is Python
	}

}
