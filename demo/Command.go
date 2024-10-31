package main

import "fmt"

func (cli *CLI) PrintBlockChain() {
	bc := cli.bc
	it := bc.NewIterator()

	for {
		block := it.Next()
		//	fmt.Printf("============当前区块高度：%d============\n", i)
		//block := NewBlock("老师转班长一个比特币", []byte{})
		fmt.Println("==========================================\n\n")

		fmt.Printf("版本号：%d\n", block.Version)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("梅克尔根：%x\n", block.MerkelRoot)
		fmt.Printf("时间戳：%d\n", block.TimeStamp)
		fmt.Printf("难度值（随便写的）：%d\n", block.Difficulty)
		fmt.Printf("随机数：%d\n", block.Nonce)
		fmt.Printf("当前区块哈希值）：%d\n", block.Difficulty)
		fmt.Printf("区块数据：%s\n", block.Data)
		//fmt.Println("hello")
		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束")
			break
		}
	}

}

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块成功!")
}
