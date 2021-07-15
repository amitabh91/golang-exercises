/* Create struct type person with fields first_name,last_name,flavour_icream's, store multiple ice flavour's. Use map and take last_name as key and person details as a value.
Print the person's details and range over ice flavour.
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
	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}
	for _, v := range m {
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for j, k := range v.flavour_icream {
			fmt.Println(j, k)
		}
		fmt.Println("===============")
	}
}
