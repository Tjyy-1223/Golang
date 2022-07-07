package main

var m = make(map[int]int)

func testSyncMap() {
	go func() {
		m[1] = 1
	}()

	go func() {
		_ = m[1]
	}()

	for {

	}
}

func main() {
	testSyncMap()
}
