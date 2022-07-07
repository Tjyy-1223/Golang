package main

import "fmt"

type Profile struct {
	Name    string
	Age     int
	Married bool
}

/*
func simpleHash(str string) (ret int) {
	for i := 0; i < len(str); i++ {
		c := str[i]
		ret += int(c)
	}
	return ret
}

type classicQueryKey struct {
	Name string
	Age  int
}

func (c *classicQueryKey) hash() int {
	return simpleHash(c.Name) + c.Age*1000000
}
*/

type queryKey struct {
	Name string
	Age  int
}

// 创建哈希值到数据的索引关系
var mapper = make(map[interface{}]*Profile)

func buildIndex(list []*Profile) {
	for _, profile := range list {
		key := queryKey{profile.Name, profile.Age}
		mapper[key] = profile
	}
}

func queryData(name string, age int) {
	key := queryKey{name, age}
	result, ok := mapper[key]

	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("no found")
	}
}

func testMapKey() {
	list := []*Profile{
		{Name: "zhangsan", Age: 30, Married: true},
		{Name: "lisi", Age: 21},
		{Name: "wangmazi", Age: 21},
	}
	buildIndex(list)
	queryData("zhangsan", 30)
	//fmt.Println(list)
}

func main() {
	testMapKey()
}
