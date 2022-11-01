// package clause
package main

//import statements
import "fmt"

// main code-block
func main() {

	// assign and set decoupled syntax
	// var name, surname string = "Hasan", "SAHIN"

	//Assign and set couple syntax
	name, surname := "Hasan", "SAHIN"

	// assign and set syntax
	// var age int = 40
	age := 40

	//only set syntax
	age = 38

	fmt.Println("Hello, ", name, surname, ". You are ", age, ".")

	// var isMarried bool = true
	isMarried := true

	if isMarried {
		fmt.Println("You get married")
	}
}
