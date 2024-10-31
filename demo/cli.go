package main

import (
	"fmt"
	"os"
)

//这是一个用来接收命令行参数并控制区块链操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `   
addBlock --data DATA    "add data to blockchain"
printChain             "print all blockchain data"
`

//接收参数的动作，我们放到一个函数中

func (cli *CLI) Run() {
	//1得到所有命令参数
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}

	//2分析命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		//添加区块
		//a获取数据； b使用bc添加区块
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Println("添加区块参数使用不当，请检查")
		}
	case "printChain":
		//打印
		fmt.Printf("打印区块\n")
		cli.PrintBlockChain()

	default:
		fmt.Println("无效命令")
		fmt.Printf(Usage)

	}
}
