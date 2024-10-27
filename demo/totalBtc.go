package main

import "fmt"

func main2() {
	fmt.Println("hello")
	//1.每21万个减半
	//2.最初50
	//3.用一个循环
	total := 0.0
	blockInterval := 21.0 //万
	currentReward := 50.0
	for currentReward > 0 {
		amount1 := blockInterval * currentReward
		currentReward *= 0.5
		total += amount1
	}
	fmt.Println("比特币总量：", total, "万")

}
