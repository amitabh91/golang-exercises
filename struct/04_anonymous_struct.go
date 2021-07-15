//Create an anonymous struct with fields type as string, map, and []string

package main

import "fmt"

func main() {
	s := struct {
		name      string
		contacts  map[string]int
		favDrinks []string
	}{
		name: "James",
		contacts: map[string]int{
			"emergency": 911,
			"ambulance": 108,
		},
		favDrinks: []string{"Diet-Coke", "ThumpsUp", "Ice-Tea"},
	}
	fmt.Printf("%+v", s)
}
