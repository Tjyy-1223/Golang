package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(header string, channel chan<- string) {
	//无限循环 不停地产生数据
	for {
		//将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())
		//等待1s
		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string) {
	for {
		// 从通道中取出数据
		message := <-channel
		fmt.Println(message)
	}
}

func main() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()函数兵法的goroutine
	go producer("cat", channel)
	go producer("dog", channel)
	customer(channel)

}
