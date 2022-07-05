package main

import "fmt"

type Weapon int

const (
	Arrow Weapon = iota // 从0开始枚举生成
	Shuriken
	SniperRifle
	Rigle
	Blower
)

func myEnum() {
	fmt.Println(Arrow, Shuriken, SniperRifle, Rigle, Blower)

	var weapen Weapon = Blower
	fmt.Println(weapen)
}

type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

//func main() {
//	fmt.Printf("%s,%d\n", CPU, GPU)
//	fmt.Printf("%T,%T\n", CPU, GPU)
//}
