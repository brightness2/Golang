package sync

import (
	"fmt"
	"math/rand"
	"time"
)

/**
channel 通道是用于协程之间进行通信的
channel 有无缓冲和有缓冲的两种，无缓冲是同步的，有缓冲可以实现异步的
比如：一条线表示一个协程，口表示通道，协程的通信通过口传递的
	|           |
	| -> 口 ->  |
	|           |
	|           |
注意：数据发送到通道后必须关闭通道，避免死锁
*/

func TestChan() {
	fmt.Println("chan...")
	// chan2()
	// chan3()
	chan4()
}

func chan1() {
	/***创建通道***/
	// unbufered := make(chan int) //整形无缓冲通道
	// bufered := make(chan string, 10) //字符串形有缓冲通道
	// unbufered <- 20 //发送数据，把整形20发送到unbufered通道
	// res := <-unbufered //接收数据，从unbufered通道接收数据
}

var bufered = make(chan int)

func chan2() {
	defer close(bufered)
	go send()
	fmt.Println("chan2 wait...")
	value := <-bufered
	fmt.Println("receive: ", value)
	fmt.Println("chan2 end...")

}

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Println("send: ", value)
	bufered <- value
}

/***************channel 的遍历********************/
var c = make(chan int) //此时是无缓冲的通道，以下运行时就会同步的

func chan3() {
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c) //不关闭通道，通道被多读就会出现死锁情况，关闭了就会读到所属类型的默认值
	}()

	//遍历方式1
	// for i := 0; i < 2; i++ {
	// 	r := <-c
	// 	fmt.Printf("r: %v\n", r)
	// }

	//遍历方式2 推荐
	for v := range c {
		fmt.Printf("v: %v\n", v)
	}

	//遍历方式3
	// for {
	// 	v, ok := <-c
	// 	if ok {
	// 		fmt.Printf("v: %v\n", v)
	// 	} else {
	// 		break
	// 	}

	// }

}

/******************** 用于channel 的select 语句 *****************************/
//select 的case 语句必须是一个channel操作，default语句总是可运行的,如果多个case可以执行，会随机选择一个执行
//注意：没有可执行的case且没有default语句时，会阻塞
var c_string = make(chan string)
var c_int = make(chan int)

func chan4() {
	go func() {
		c_int <- 100
		c_string <- "hello"
		// close(c_int)
		// close(c_string)
	}()

	for {
		select {
		case r := <-c_int:
			fmt.Printf("c_int: %v\n", r)
		case r := <-c_string:
			fmt.Printf("c_string: %v\n", r)
		default:
			fmt.Println("chan4 default")
		}
		time.Sleep(time.Second)
	}

}
