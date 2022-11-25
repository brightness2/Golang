package sync

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/*
协程实现同步
*/

func TestWait() {
	// byWaitGroup()
	// byGosched()
	// byGoexit()
	// byGOMAXPROCS()
	// bySyncMutex()
	byAtomic()
}

/******* 通过waitGroup实现同步*****/
//不保证原子量一致性，如果要保证原子量，可以配和sync.Mutex互斥锁或sync/atomic原子操作实现
var wg sync.WaitGroup

func byWaitGroup() {
	for i := 0; i < 4; i++ {
		wg.Add(1) //协程登记 +1
		go hello(i)
	}
	wg.Wait() //等待所有登记的协程都接收
}

func hello(i int) {
	defer wg.Done() //协程结束，登记-1， 如 wg.Add(-1) 等效
	fmt.Println("Hello Brightness", i)
}

/************ 通过 runtime.Gosched 实现同步 *************/
//runtime.Gosched 方法作用是让出执行时间片给其它协程

func byGosched() {
	go show("golang")

	for i := 0; i < 10; i++ {
		runtime.Gosched() //切一下，再次分配任务
		fmt.Println("Goshed")
	}
	// runtime.Gosched() //切一下，再次分配任务
	// fmt.Println("Goshed")

}

func show(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, i)
	}
}

/************  runtime.Goexit  *************/
// runtime.Goexit 是退出协程

func byGoexit() {
	go show2()
	time.Sleep(time.Second)
}

func show2() {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit()
		}
		fmt.Println("show2", i)
	}
}

/************  runtime.GOMAXPROCS  *************/
//  runtime.GOMAXPROCS 设置运行时使用的核心个数

func byGOMAXPROCS() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go a()
	go b()

	time.Sleep(time.Second)
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A: ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B: ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

/************ 通过  sync.Mutex 互斥锁 *************/
// sync.Mutex 进行上锁和解锁
var wg2 sync.WaitGroup
var m int = 100 //通过锁保证原子量 m值一致性
var lock sync.Mutex

func bySyncMutex() {
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go add()
		wg2.Add(1)
		go sub()
	}
	wg2.Wait()
	fmt.Printf("m: %v\n", m) //如果没有锁，m的结果有可能不等于100
}

func add() {
	defer wg2.Done()
	lock.Lock()
	m += 1
	fmt.Println("i++:", m)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {
	defer wg2.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	m -= 1
	fmt.Println("i--:", m)

	lock.Unlock()

}

/************ 通过  sync.atomic 原子 *************/
//sync.atomic 保证原子量一致性,atomic 的实现是通过cas （compare and swap）比较并替换 实现的，就是进行操作前进行比较，判断是否是原值，是则交换新值，否则不操作
var m2 int32 = 50

func byAtomic() {
	for i := 0; i < 100; i++ {
		go add2()
		go sub2()
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("m2: %v\n", m2)
}

func add2() {
	atomic.AddInt32(&m2, 1)
	fmt.Println("m++", m2)
}

func sub2() {
	atomic.AddInt32(&m2, -1)
	fmt.Println("m--", m2)

}
