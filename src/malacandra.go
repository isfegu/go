package main

import("fmt")

func main() {
	var daysTraveling = 28
	const hoursPerDay = 24
	var distanceEarthMalacandra = 56000000 // Km
	

	fmt.Println(distanceEarthMalacandra / daysTraveling * hoursPerDay, " speed needed (km/h)")
}