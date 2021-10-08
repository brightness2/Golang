/*
 * @Author: Brightness
 * @Date: 2021-10-08 13:57:44
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-08 15:32:20
 * @Description:
 */
package myPackage

import "fmt"

/*
* 基本数据类型
* bool
* byte
* int8 int16 float32 float64
* string
*
 */
/*
* 复杂数据类型
	指针
	数组
	结构体
	管道
	函数
	切片
	接口
	map (集合)
*/
/*变量定义*/
var Flag bool = true
var char byte = 'c'
var size uint8 = 20
var name string = "Brightness"

var p *uint8 = &size                 //指针
var arr [2]int = [2]int{2, 3}        //一维数组
var arr2 = [2][2]int{{1, 2}, {3, 4}} //二维数组

//结构体，具有封装，继承，多态特殊性
type animal struct {
	Name string
	Sex  string
}

//结构体方法
func (a animal) say() {
	fmt.Println("animal say", a.Name)
}

//结构体实例
var an1 = animal{
	Name: "dog",
	Sex:  "m",
}

//结构体继承
type cat struct {
	animal //通过匿名属性
	foot   int
}

var ca1 = cat{animal: animal{Name: "tom", Sex: "m"}, foot: 4}

var slice = arr[0:1]        //切片,对数组的部分地址映射,修改切片同时修改数组的值
var slice2 = make([]int, 2) //长度为2的数组切片
var m = map[int]string{
	1: "one",
	2: "two",
}
var m2 = make(map[int]string) //集合，key唯一

func init() {
	//调用结构体方法
	an1.say()
	ca1.say()
}
