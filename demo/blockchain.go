package main

import (
	"firstProject/bolt"
	"log"
)

// 6 引入区块链
type BlockChain struct {
	//定义一个区块链数据
	//blocks []*Block
	//使用数据库代替数组
	db   *bolt.DB
	tail []byte //存储最后一个区块的哈希
}

const blockChainDb = "blockchain.db"
const blockBucket = "blockBucket"

// 7.定义一个区块链
func NewBlockChain() *BlockChain {
	//创建一个创世块，并作为第一个区块添加到区块链中
	//genesisBlock := GenesisBlock()

	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}
	//最后一个区块的hash,从数据库里读出来的
	var lastHash []byte
	//打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}

	//2 找到抽屉bucket（如果没有，就创建）
	db.Update(func(tx *bolt.Tx) error {
		//2 找到抽屉bucket（如果没有，就创建）
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//没有抽屉
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建失败")
			}

			//创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock()
			//写数据
			//hash体做为key,block字节流做为value
			bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

// 创世块
func GenesisBlock() *Block {
	return NewBlock("Go一期创世块，老牛逼了", []byte{})
}

// 8  添加区块
func (bc *BlockChain) AddBlock(data string) {
	/*
		//如何获取前区的哈希呢？

		//获取最后一个区块
		lastBlock := bc.blocks[len(bc.blocks)-1]
		prevHash := lastBlock.Hash

		// a.创建新的区块
		block := NewBlock(data, prevHash)
		//b 添加到区块链数组中
		bc.blocks = append(bc.blocks, block)
	*/
	
}

//9 重构代码
