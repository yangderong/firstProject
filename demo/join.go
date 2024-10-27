package main

import (
	"fmt"
	"strings"
)

func main3() {

	str1 := []string{"hello", "world", "!"}
	res := strings.Join(str1, "+")
	fmt.Println("res:%s\n", res)
	//bytes.Join(s[][]{[]byte{"hello"},[]byte("world")})
}
