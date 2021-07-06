package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	city       string
	country    string
}

func main() {
	p1 := person{
		first_name: "Amitabh",
		last_name:  "Sinha",
		city:       "Bangalore",
		country:    "India",
	}
	p2 := person{
		first_name: "Vicky",
		last_name:  "Sinha",
		city:       "Auckland",
		country:    "New zealand",
	}

	fmt.Println("Data of person p1")
	fmt.Println("\tFirst Name:", p1.first_name, "\n", "\tLast Name:", p1.last_name, "\n", "\tCity:", p1.city, "\n", "\tCountry:", p1.country)
	fmt.Println("Data of person p2")
	fmt.Println("\tFirst Name:", p2.first_name, "\n", "\tLast Name:", p2.last_name, "\n", "\tCity:", p2.city, "\n", "\tCountry:", p2.country)
}
