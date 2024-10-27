package main

import (
	"firstProject/bolt"
	"fmt"
	"log"
)

func mainT() {
	fmt.Println("Hello world")
	//1 打开数据库
	db, err := bolt.Open("test.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}

	//2 找到抽屉bucket（如果没有，就创建）
	db.Update(func(tx *bolt.Tx) error {
		//2 找到抽屉bucket（如果没有，就创建）
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			//没有抽屉
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建失败")
			}

		}

		bucket.Put([]byte("111111"), []byte("hello1"))
		bucket.Put([]byte("222222"), []byte("world2"))
		return nil
	})

	//3 写数据

	//4 读数据
	db.View(func(tx *bolt.Tx) error {
		//1找到数据库，没有就退出
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket b1不应为空")
		}

		//2直接读数据
		v1 := bucket.Get([]byte("111111"))
		v2 := bucket.Get([]byte("222222"))

		fmt.Printf("v1:%s\n", v1)
		fmt.Printf("v2:%s\n", v2)
		return nil

	})

}
