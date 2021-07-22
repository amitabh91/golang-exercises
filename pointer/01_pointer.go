// Create a value and assign it to a variable then print the address of that variable
package main

import "fmt"

func main() {
	a := 10
	fmt.Println("The value of a is", a, "and it stored at", &a)

}
