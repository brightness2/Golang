/*
 * @Author: Brightness
 * @Date: 2021-10-08 14:26:59
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-08 15:11:35
 * @Description:
 */
package myPackage

import (
	"errors"
	"fmt"
)

func Test() {
	arr := [3]int{2, 4, 6}
	one := 6
	//if
	if one == 1 {
		fmt.Println("1111")
	} else if one == 2 {
		fmt.Println("222222")
	} else {
		// fmt.Println("3333")
	}

	//循环
	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[i])
	}

	// for key, value := range arr {
	// 	fmt.Println("key=", key, "value=", value)
	// }

	for _, value := range arr {
		fmt.Println(value)
	}

	//闭包调用
	f := fn()
	f()
	fmt.Println(f())

	//测试defer
	testDefer()

	//defer + recover 进行错误处理
	testRecover()

}

/*闭包*/
func fn() func() int {
	var count int = 0 //这个会变量一直存在内存中

	return func() int {
		count += 10
		return count
	}
}

/*defer 关键字*/
func testDefer() {
	//defer 把执行语句放进栈中，等其它语句执行完成后再从栈中取出语句执行，注意栈是先进后出的数据结构

	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("three")
	//执行结果
	// three
	// two
	// one
}

/*defer + recover 错误处理*/

func testRecover() {
	num1 := 10
	num2 := 0
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误已捕获")
			e := errors.New("除数不能为0")
			fmt.Println(e)
		}
	}()

	fmt.Println(num1 / num2) //panic: runtime error: integer divide by zero

}
