package main

import "sync"

// 并发使用map 为防止竞争问题 使用mutex进行加锁
var (
	valueByKey      = make(map[string]int)
	valueByKeyGuard sync.Mutex
)

func readValue(key string) int {
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}
