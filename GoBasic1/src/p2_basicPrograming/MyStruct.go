package main

import (
	"fmt"
	"reflect"
)

type Brand struct {
}

// 为商标结构添加show()方法
func (t Brand) show() {

}

// FakeBrand Brand起一个别名FakeBrand
type FakeBrand = Brand

// Vehicle 定义车辆结构
type Vehicle struct {
	FakeBrand
	Brand
}

func myStruct() {
	var a Vehicle

	a.FakeBrand.show()

	ta := reflect.TypeOf(a)
	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		fmt.Printf("FiledName : %v , FiledType : %v\n", f.Name, f.Type)
	}
}

//func main() {
//	myStruct()
//}
