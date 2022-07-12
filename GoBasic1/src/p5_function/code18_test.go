package main

import "testing"

func GetArea(weight, height int) int {
	return weight * height
}

func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("test error")
	}
}
