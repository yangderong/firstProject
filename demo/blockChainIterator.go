package main

import (
	"firstProject/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	//游标用于不断索引
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {

	return &BlockChainIterator{
		bc.db,
		//最初指向区块的最后一个区块，随着next的调用，会不断变化
		bc.tail,
	}
}

// 迭代器属于区块链的
// next方未能属于迭代器的
// 1返回当前区块 2指针前移
func (it *BlockChainIterator) Next() *Block {

	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket应不为空，请检查")
		}

		blockTmp := bucket.Get(it.currentHashPointer)
		//解码动作
		block = DeSerialize(blockTmp)
		it.currentHashPointer = block.PrevHash //哈希左移
		return nil
	})

	return &block
}
