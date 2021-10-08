/*
 * @Author: Brightness
 * @Date: 2021-10-08 12:31:24
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-08 13:53:58
 * @Description:
 */
package myPackage

import "fmt"

var Name string = "myPackage"
var age int = 18
var a = func() int {
	fmt.Println("----myPackage 变量声明----")
	return 1
}()

func init() {
	fmt.Println("----myPackage init函数----")
}

func GetAge() int {
	fmt.Println("myPackage GetAge函数被调用")
	return age
}
