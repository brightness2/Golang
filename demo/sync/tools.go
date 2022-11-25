package sync

import (
	"fmt"
	"time"
)

/*
工具类的使用
*/

func TestTools() {
	// tools1()
	// tools2()
	tools3()
}

/*************Timer 定时器********************/
//类似JavaScript的setTimeout，是通过channel实现的
func tools1() {
	//创建定时器
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("start:%v\n", time.Now())

	t1 := <-timer.C //下面的代码会阻塞2秒，再执行，如果是单纯的等待可以用time.Sleep
	fmt.Printf("t1: %v\n", t1)

	time.Sleep(time.Second * 2)
	fmt.Println("又2秒后")

	<-time.After(time.Second * 2)
	fmt.Println("...")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired") //timer2定时器停止后，不会执行
	}()
	stop := timer2.Stop() //停止定时器,不会执行timer2定时器后面的代码
	if stop {
		fmt.Println("stop timer2")

	}

	timer3 := time.NewTimer(time.Second * 5) //原来定时5秒
	timer3.Reset(time.Second)                //重置为1秒
	<-timer3.C
	fmt.Println("timer3 ...")

	time.Sleep(time.Second * 3)
}

/*************Ticker 断续器*****************/
//类似JavaScript的setInterval，可以周期执行
func tools2() {
	ticker1 := time.NewTicker(time.Second * 2)
	counter := 0
	for _ = range ticker1.C {
		fmt.Println("ticker1")
		counter++
		if counter >= 5 {
			ticker1.Stop()
			break
		}
	}
}

var ch = make(chan int)

func tools3() {
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		defer close(ch)
		for _ = range ticker.C {
			select {
			case ch <- 1:
				fmt.Println("发送 1")
			case ch <- 2:
				fmt.Println("发送 2")
			case ch <- 3:
				fmt.Println("发送 3")
			default:
				fmt.Println("default")
			}
		}
	}()
	counter := 0
	for v := range ch {
		fmt.Println("收到 ", v)
		counter++
		if counter >= 5 {
			ticker.Stop()
			break
		}
	}
}
