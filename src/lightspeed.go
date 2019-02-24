// Exercise 2.3

package main

import("fmt")

func main () {
	const lightSpeed = 299792 // Km/s
	var distance = 56000000 // Km

	fmt.Println(distance / lightSpeed, "seconds")

	distance = 401000000 // Km
	fmt.Println(distance / lightSpeed, "seconds")

	// SpaceX
	const hourPerDay = 24
	const spaceXSpeed = 100800 // Km/h
	var distanceEarthMars = 96300000 // Km
	fmt.Println(distanceEarthMars / spaceXSpeed / hourPerDay, "days")
}