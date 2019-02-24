// Exercise 2.4

package main

import("fmt")

func main() {
	var (
		weight = 149.0
		age = 41
	)

	var count, price = 10, 100

	weight *= 0.3783 // The same as weight = weight * 0.3783

	fmt.Println("My weight on the surface of Mars is", weight)

	age++
	fmt.Println("Age:", age)
	age += 1
	fmt.Println("Age:", age)

	count--
	fmt.Println("Count:", count)
	price /= 2
	fmt.Println("Price:", price)

	weight -= 2
	fmt.Println("Weight:", weight)
}