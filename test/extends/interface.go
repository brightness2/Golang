package extends

import (
	"fmt"
)
/***********接口的使用******************************/
func TestInterface(){
	var phone I_Phone
	nokia := NokiaPhone{color:"red"}
	phone = nokia
	phone.call()
	mi := &MiPhone{size:14.2}//接受者是指针类型,所以要传指针值
	phone = mi
	phone.call()
	//mi 调用usb接口
	var usb I_Usb
	usb = mi
	usb.read()
	//usb = nokia //这里报错，因为没有实现usb接口，nokia NokiaPhone cannot use nokia (type NokiaPhone) as type I_Usb in assignment: NokiaPhone does not implement I_Usb (missing read method)

	//用新的接口
	fmt.Println("---------------用新的接口---------------------")
	var  new_phone I_New_Phone
	// new_phone = nokia //这里会报错，因为没有实现read方法
	new_phone = mi
	new_phone.call()
	new_phone.read()

	//开闭原则
	fmt.Println("---------------开闭原则---------------------")
	p := Person{name:"tom"}

	p.usePhone(nokia)
	p.usePhone(mi)
	
	wei := &WeiPhone{}
	p.usePhone(wei)

	/********************常见使用***********************/
	var phone2 I_Phone
	phone_name := "wei"
	switch phone_name {
		case "nokia":
			phone2 = NokiaPhone{}
		case "mi":
			phone2 = &MiPhone{}
		case "wei":
			phone2 = &WeiPhone{}
	}

	p.usePhone(phone2)

}

type I_Phone interface{
	call()
}


type NokiaPhone  struct {
	color string
}
//注意，接受者是值类型
func (nokiaPhone NokiaPhone) call(){
	fmt.Println(nokiaPhone.color,"NokiaPhone  call...")
}

type MiPhone  struct {
	size float32
}
//注意，接受者是指针类型
func (miPhone *MiPhone) call(){
	fmt.Println(miPhone.size,"MiPhone call...")
}
//定义了usb 接口
type I_Usb interface {
	read()
}
//miPhone 又实现了usb接口，也就是两个接口
func (miPhone *MiPhone) read(){
	fmt.Println(miPhone.size,"MiPhone read...")
}

//接口嵌套,得到一个新的接口
type I_New_Phone interface {
	I_Phone
	I_Usb
}

/**************************通过接口实现 OCP设计原则 （开-闭原则）,对扩展开放，对修改关闭************************************************************/

//定义person结构体
type Person struct {
	name string
}

func (person *Person) usePhone(phone I_Phone){
	fmt.Println("person use")
	phone.call()
}

//这时候扩展了一个 WeiPhone,person也可以调用
type WeiPhone struct {
		
}

func (weiPhone *WeiPhone) call(){
	fmt.Println("WeiPhone call...")
}