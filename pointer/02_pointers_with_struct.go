/*
Create a person struct, create a function changeMe with *person as argument.
Change the value at *persons address.
Create a value of type person
In main func()
	1.) Print out the value of person.
	2.) call changeMe().
	3.) print out the value of person.
*/

package main

import "fmt"

type person struct {
	last    string
	first   string
	address string
}

func changeMe(p1 *person) {
	(*p1).address = "Bangalore"
}

func main() {
	p1 := person{
		last:    "Taylor",
		first:   "Swift",
		address: "New York",
	}
	fmt.Printf("%+v\n", p1)
	changeMe(&p1)
	fmt.Printf("%+v\n", p1)
}
