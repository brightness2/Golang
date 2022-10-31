package main

import (
	"fmt"
	"test/extends"
)

func main(){
	fmt.Println("main2")
	extends.TestClass()
	fmt.Println("------------------------------------")
	extends.TestInterface()
	// fmt.Println(extends.o) //不可以访问小写开头的变量
	//模拟 OOP（对象） 的属性和方法
	fmt.Println(extends.Cat_obj.Name)//可以访问大写开头的属性
	// fmt.Println(extends.Cat_obj.color) //不可以访问小写开头的属性
	extends.Cat_obj.Jump()//可以调用大写开头的方法
	// extends.Cat_obj.eat()//不可以调用小写开头的方法
	c,err := extends.NewCat("kok","red")
	if(err != nil){
		fmt.Println("cat 创建失败")
	}
	fmt.Println(*c)
}