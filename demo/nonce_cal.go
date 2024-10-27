package main

import (
	"crypto/sha256"
	"fmt"
)

func main1() {

	fmt.Println("hello")
	//交易数据
	data := "helloworld"
	for i := 0; i < 10000; i++ { //1000000
		hash := sha256.Sum256([]byte(data + string(1)))
		fmt.Println("hash: %x\n", hash[:])

	}
}
