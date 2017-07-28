package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := make(chan int, 1000) //使用信道缓存，在缓存未存满之前不会挂起线程
	o := make(chan bool) //线程结束判断信道

	//给 cs 通道写入数据
	for i := 0; i < 10; i++ {
		cs <- i
	}

	close(cs)

	//创建信道信号接收线程
	go func(){
		var e int
		ok := true

		for{
			select {
			case e, ok = <- cs:
				fmt.Println(ok)
				if !ok {
					fmt.Println("End.")
					break
				}
				fmt.Printf("cs -> %d\n",e)
			case <- time.After(300 * time.Millisecond): //超时
				fmt.Println("Timeout.")
				ok = false
				break
			}
			//终止for循环
			if !ok {
				o <- true
				break
			}
		}

	}()
	//开启线程
	<- o
}
func Go(c chan int, index int) {
	//累加方法
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	//测试输出
	fmt.Printf("%6s: %d, %s: %d \n", "index", index, "number", a)
	//信道输入
	c <- index
}
func GoW(wg *sync.WaitGroup, index int) {
	//累加方法
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	//测试输出
	fmt.Printf("%6s: %d, %s: %d \n", "index", index, "number", a)
	//线程结束
	wg.Done()
}
