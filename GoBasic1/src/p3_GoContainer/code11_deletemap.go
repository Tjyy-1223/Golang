package main

import "fmt"

func testDeleteMap() {
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	delete(scene, "brazil")

	for key, value := range scene {
		fmt.Println(key, value)
	}
}

//func main() {
//	testDeleteMap()
//}
