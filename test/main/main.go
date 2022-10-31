/*
 * @Author: Brightness
 * @Date: 2021-10-08 10:37:15
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-08 14:28:22
 * @Description:
 */
/*包名称，与文件夹同名*/
package main

/*导入包，可以是自定义的包*/
import (
	"fmt"
	"test/myPackage"
)

var one = func() string {
	fmt.Println("----main 变量声明----")
	return "one"
}()

func init() {
	fmt.Println("----main init函数----")
}

func main() {
	fmt.Println("----main函数----")
	//变量短声明,必须在函数内部
	// age := 45
	fmt.Println("myPackage 对外变量Name=", myPackage.Name)
	myPackage.GetAge()
	myPackage.Test()
}
