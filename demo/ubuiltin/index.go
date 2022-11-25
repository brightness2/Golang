package ubuiltin

import "fmt"

/**
builtin 包
用于类型声明，变量和常量声明，一些便捷函数，这些包不用导入，就何以使用这些声明好的
比如 append函数 ，len 函数，int类型，string类型
*/
func TestBuiltin() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获panic：", err)
		}
	}()
	// testAppend()
	// testLen()
	testPanic()
}

func testAppend() {
	s1 := []int{1, 2, 3}
	i := append(s1, 100)
	fmt.Printf("i: %v\n", i)
	s2 := []int{4, 5, 6}
	i2 := append(s1, s2...)
	fmt.Printf("i2: %v\n", i2)
}

func testLen() {
	fmt.Printf("len(\"brightness\"): %v\n", len("brightness"))
}

func testPanic() {
	panic("danger ! .....")

}

/********make 和 new*******/
//make 只能初始化 slice 、map 、chan,返回的是引用，T
//new 可以初始化所有类型，返回的是指针 ， *T
