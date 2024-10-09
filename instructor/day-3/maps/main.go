package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["delhi"] = 10
	m["london"] = 20

	fmt.Println(m["london"])

	// days := make(map[int]string)
	// days[1] = "Monday"
	// days[2] = "Tuesday"
	// days[3] = "Wednesday"
	// days[4] = "Thursday"
	// days[5] = "Friday"
	// days[6] = "Saturday"
	// days[7] = "Sunday"

	// Compact literal initalisation
	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	_ = days

	moons := map[string]int{
		"Mercury": 0,
		"Venus":   0,
		"Earth":   1,
		"Mars":    2,
		"Jupiter": 67,
	}

	// if n, ok := moons["Neptune"]; ok {
	// 	fmt.Println("Number of moons : ", n)
	// } else {
	// 	fmt.Println("No moons")
	// }
	delete(moons, "Mars1")
	fmt.Println(moons)

	cities := map[string]int{
		"Tokyo": 42100000, "Delhi": 19500000, "Sau Paulo": 2300000, "New York": 18600000, "Moscow": 2150000,
	}

	for _, val := range cities {
		fmt.Println(val)
	}
}
