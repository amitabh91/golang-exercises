/* Create a struct type vehicle, truck, and sedan with below-mentioned fields. And embed struct truck and sedan with struct type vehicle
vehicle fields:
	door
	color
truck fields:
	sixWheel  should be type bool
sedan fields:
	luxury should be type bool

Print each fields of struct truck and sedan.
*/

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	sixWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		sixWheel: true,
		vehicle: vehicle{
			doors: 2,
			color: "Black",
		},
	}
	s := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
	}
	// printing eaching fields values
	fmt.Println("vechile is type truck:", t.sixWheel)
	fmt.Println("doors:", t.doors)
	fmt.Println("color:", t.color)
	fmt.Println("===========")
	fmt.Println("vechile is type Sedan:", s.luxury)
	fmt.Println("doors:", s.doors)
	fmt.Println("color:", s.color)
}
