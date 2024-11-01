package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

// 0 定义结构
type Block struct {
	//版本号
	Version uint64
	//1 前区块哈希
	PrevHash []byte
	//Merkel根，梅克尔根，这就是一个哈希值，我们先不管，后面再介绍
	MerkelRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Difficulty uint64
	//随机数，也就是挖矿要找的数据
	Nonce uint64

	//2 当前区块哈希   正常比特比区块中没有当前区块的hash,我们为了方便做了简化
	Hash []byte
	//3 数据
	Data []byte
}

//1.补充区块字段
//2.更新计算哈希函数
//3.优化代码

// 实现一个辅助函数，功能是将uint64转成[]byte
func Uint64TOByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// 4 创建区块
func NewBlock(data string, prevBloackHash []byte) *Block {
	block := Block{
		Version:    0,
		PrevHash:   prevBloackHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,        //随便写的，无效值
		Nonce:      0,        //同上A
		Hash:       []byte{}, //先填空。再计算
		Data:       []byte(data),
	}
	//block.SetHash()
	//创建一个pow对象

	pow := NewProofOfWork(&block)

	//查找随机数，不停的进行哈希运算
	hash, nonce := pow.Run()

	//根据挖矿结果对区块数据进行更新(补充)
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

// 序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer

	//使用gob进行序列化（编码）得到字节流
	//1定义一个编码器
	//2使用编码器进行编码

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码出错，")
	}
	//fmt.Printf("编码后的小明: %v\n", buffer.Bytes())
	return buffer.Bytes()
}

// 反序列化
func DeSerialize(data []byte) Block {

	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block

	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错")
	}
	//fmt.Printf("解码后的：: %v\n", daMing)

	return block
}

// 5 生成哈希
func (block *Block) SetHash() {
	//	var blockInfo []byte
	//1。拼装数据
	/*
		blockInfo = append(blockInfo, Uint64TOByte(block.Version)...)
		blockInfo = append(blockInfo, block.PrevHash...)
		blockInfo = append(blockInfo, block.MerkelRoot...)
		blockInfo = append(blockInfo, Uint64TOByte(block.TimeStamp)...)
		blockInfo = append(blockInfo, Uint64TOByte(block.Difficulty)...)
		blockInfo = append(blockInfo, Uint64TOByte(block.Nonce)...)
		blockInfo = append(blockInfo, block.Data...)
	*/
	tmp := [][]byte{
		Uint64TOByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64TOByte(block.TimeStamp),
		Uint64TOByte(block.Difficulty),
		Uint64TOByte(block.Nonce),
		block.Data,
	}
	//将二维的切片数组链接起来，返回一个新的一维
	blockInfo := bytes.Join(tmp, []byte{})

	//2. sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
