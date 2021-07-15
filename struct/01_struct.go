/* Create struct type person with fields first_name,last_name & flavour_icream's.
   Create two values of type person and print out the values ranging over the elements in the slice.
*/
package main

import "fmt"

type person struct {
	first_name     string
	last_name      string
	flavour_icream []string
}

func main() {
	p1 := person{
		first_name:     "James",
		last_name:      "Bond",
		flavour_icream: []string{"Chocolate", "Butterscotch"},
	}
	p2 := person{
		first_name:     "Taylor",
		last_name:      "Swift",
		flavour_icream: []string{"BlackCurrent", "Strawberry"},
	}
	fmt.Println("First Name:", p1.first_name)
	fmt.Println("Last Name: ", p1.last_name)

	for _, k := range p1.flavour_icream {
		fmt.Println("\tFav Icream Flavour's:", k)
	}

	fmt.Println("First Name:", p2.first_name)
	fmt.Println("Last Name: ", p2.last_name)

	for _, k := range p2.flavour_icream {
		fmt.Println("\tFav Icream Flavour's:", k)
	}
}
