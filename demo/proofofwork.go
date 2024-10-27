package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义一个工作量证明的结构ProofOfWork

type ProofOfWork struct {
	//a. block
	block *Block
	//b.  目标值
	//一个非常大的数，它有很多方法
	target *big.Int
}

//2.提供创建POW函数

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block}
	//我们指定的难度值
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"

	//引入辅助变量，目的是将上面的string类型进行转换
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow
}

//3 提供计算不断计算hash的函数

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//1 拼装数据  区块的数据 还有不断变化的随机数
	//2.做哈希运算
	//3 与pow中的target进行比较
	//a:找到了，退出返回
	//b 没有找到，继续找，随机数+1
	var nonce uint64
	block := pow.block
	var hash [32]byte
	fmt.Println("开始挖矿")
	for {
		//1 拼装数据  区块的数据 还有不断变化的随机数
		tmp := [][]byte{
			Uint64TOByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64TOByte(block.TimeStamp),
			Uint64TOByte(block.Difficulty),
			Uint64TOByte(nonce),
			block.Data,
		}
		//将二维的切片数组链接起来，返回一个新的一维
		blockInfo := bytes.Join(tmp, []byte{})
		//2.做哈希运算
		hash = sha256.Sum256(blockInfo)
		//3 与pow中的target进行比较
		tmpInt := big.Int{}

		//fmt.Printf("hash:%x\n,nonce: %d\n", hash, nonce)
		//将我们得到的hash值 转化为bigInt
		tmpInt.SetBytes(hash[:])
		//比较当前的hash值与目标的hash值，如果当前的小于目标的，就说明找到了，否则继续找
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功！ hash:%x,nonce: %d\n", hash, nonce)
			//break
			return hash[:], nonce
		} else {
			nonce++
		}

		//a:找到了，退出返回
		//b 没有找到，继续找，随机数+1

	}

	//return []byte("hello world"), 10
	//return hash[:], nonce
}
