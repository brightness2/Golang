package extends

import (
	"errors"
	"fmt"
)

/**************************class 模拟**************************************/

func TestClass(){
	
	dog := Dog{
		Animal{
			name:"tom",
			age:4,
		},
		"yellow",
	}
	dog.say()
}


/**
定义 Animal 
*/
type Animal struct {
	name string
	age int
	
}
//Animal 方法
func (animal *Animal) say()  {
	fmt.Println("hello,my name is",animal.name)
}

func (animal *Animal) run() {
	fmt.Println("run...")
}

/**
通过结构体嵌套实现继承
*/
type Dog struct {
	Animal
	color	string
}

func (dog *Dog) run(){
	fmt.Println("dog run...")
}

/******************模拟 OOP（对象） 的属性和方法*****************************/

type Cat struct {
	Name string //这个公开的属性，外部包可以访问
	color	string //这个是私有属性，外部包不可以访问
}

//模拟构造函数
func NewCat(name string,color string) (*Cat,error){
	if name=="" {
		return nil,errors.New("name 不能为空")
	}
	return &Cat{Name:name,color:color},nil
}

//这个是私有方法
func (cat *Cat) eat(){
	fmt.Println("cat eat...")
}
//这个是公有方法,外部包可以调用
func (cat *Cat) Jump(){
	fmt.Println("cat jump...")
}
/***************这样的就是模拟的OOP******************************/
var Cat_obj = Cat{Name:"tom",color:"white"} //变量需要大写开头才能被外部包访问
var o = 1