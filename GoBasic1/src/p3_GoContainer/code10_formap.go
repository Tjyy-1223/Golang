package main

import (
	"fmt"
	"sort"
)

func testForMap() {
	scene := map[string]int{}

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	for i := range scene {
		fmt.Println(i)
	}

	for k, v := range scene {
		fmt.Println(k, v)
	}

	var sceneList []string
	for s := range scene {
		sceneList = append(sceneList, s)
	}
	sort.Strings(sceneList)
	fmt.Println(sceneList)
}

//func main() {
//	testForMap()
//}
