package sync

import (
	"fmt"
	"sync/atomic"
)

/*
原子操作
add 增加
load 加载 就是读
store 存储 就是写
swap:直接交换
SwapT系列函数实现的交换操作，在原子性上等价于：
```
old = *addr
*addr = new
return old
```
cas:比较交换
CompareAndSwapT系列函数实现的比较-交换操作，在原子性上等价于：
```

	if *addr == old {
		*addr = new
		return true
	}

return false
```
*/
func TestAtomic() {
	// add_sub()
	// read_write()
	cas()
	// swap()
}

var num int32 = 100

func add_sub() {
	fmt.Println("add_sub")
	atomic.AddInt32(&num, 1)
	fmt.Printf("num: %v\n", num)
	atomic.AddInt32(&num, -1)
	fmt.Printf("num: %v\n", num)

}

func read_write() {
	fmt.Println("read_write")
	r := atomic.LoadInt32(&num)
	fmt.Printf("r: %v\n", r)
	atomic.StoreInt32(&num, 200)
	fmt.Printf("num: %v\n", num)
}

func cas() {
	fmt.Println("cas")
	r := atomic.CompareAndSwapInt32(&num, 100, 300)
	fmt.Printf("num: %v\n", num)
	fmt.Printf("r: %v\n", r) //true or false

}

func swap() {
	fmt.Println("swap")
	r := atomic.SwapInt32(&num, 400)
	fmt.Printf("num: %v\n", num)
	fmt.Printf("r: %v\n", r) // 原来的值

}
