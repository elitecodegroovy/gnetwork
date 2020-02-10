package main

import (
	"github.com/micro/go-micro/sync/lock"
	"github.com/micro/go-micro/sync/lock/etcd"
	"github.com/micro/go-micro/util/log"
	"time"
)

func main() {
	// 地址
	nodes := lock.Nodes("127.0.0.1:2379")

	resourceId := "id123456"
	go func() {
		lc := etcd.NewLock(nodes)
		log.Logf("协程1获取锁...")
		// 获取锁
		err := lc.Acquire(resourceId)
		if err != nil {
			log.Logf("[ERR] 协程1未得到锁")
			return
		}

		log.Logf("协程1得到锁，等一秒")
		time.Sleep(1 * time.Second)

		// 释放锁
		log.Logf("协程1释放锁")
		err = lc.Release(resourceId)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		lc := etcd.NewLock(nodes)
		log.Logf("协程2获取锁...")
		// 获取锁
		err := lc.Acquire(resourceId)
		if err != nil {
			log.Logf("[ERR] 协程2未得到锁")
			return
		}

		log.Logf("协程2得到锁，等一秒")
		time.Sleep(1 * time.Second)

		// 释放锁
		log.Logf("协程2释放锁")
		err = lc.Release(resourceId)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		lc := etcd.NewLock(nodes)
		log.Logf("协程3获取锁...")
		// 获取锁
		err := lc.Acquire(resourceId)
		if err != nil {
			log.Logf("[ERR] 协程3未得到锁")
			return
		}

		log.Logf("协程3得到锁，等一秒")
		time.Sleep(1 * time.Second)

		// 释放锁
		log.Logf("协程3释放锁")
		err = lc.Release(resourceId)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		lc := etcd.NewLock(nodes)
		log.Logf("协程4获取锁...")
		// 获取锁
		err := lc.Acquire(resourceId)
		if err != nil {
			log.Logf("[ERR] 协程3未得到锁")
			return
		}

		log.Logf("协程4得到锁，等一秒")
		time.Sleep(1 * time.Second)

		// 释放锁
		log.Logf("协程4释放锁")
		err = lc.Release(resourceId)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		lc := etcd.NewLock(nodes)
		log.Logf("协程5获取锁...")
		// 获取锁
		err := lc.Acquire(resourceId)
		if err != nil {
			log.Logf("[ERR] 协程5未得到锁")
			return
		}

		log.Logf("协程5得到锁，等一秒")
		time.Sleep(1 * time.Second)

		// 释放锁
		log.Logf("协程5释放锁")
		err = lc.Release(resourceId)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		lc := etcd.NewLock(nodes)
		log.Logf("协程6获取锁...")
		// 获取锁
		err := lc.Acquire(resourceId)
		if err != nil {
			log.Logf("[ERR] 协程6未得到锁")
			return
		}

		log.Logf("协程6得到锁，等一秒")
		time.Sleep(1 * time.Second)

		// 释放锁
		log.Logf("协程6释放锁")
		err = lc.Release(resourceId)
		if err != nil {
			log.Fatal(err)
		}
	}()
	// 等协程跑完
	time.Sleep(5 * time.Second)
}
