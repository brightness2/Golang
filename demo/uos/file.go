package uos

import (
	"fmt"
	"os"
)

func TestFile() {
	fmt.Println("file")
	// createFile()
	// createDir()
	// doRemove()
	// wd()
	readFile()
}

func createFile() {
	f, err := os.Create("a.txt") //创建的文件，默认当前路径是go.mod所在目录
	//文件存在就会覆盖
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {

		fmt.Printf("f: %v\n", f) //f: &{0xc000004a00}
		fmt.Printf("f.Name(): %v\n", f.Name())

	}
}

func createDir() {
	err := os.Mkdir("test", os.ModePerm)
	//目录存在也会报错
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//创建多个文件夹
	os.MkdirAll("test/a/b", os.ModePerm)
}

func doRemove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	os.RemoveAll("test") //删除test目录及子目录
}

func wd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("dir: %v\n", dir)
	os.Chdir("g:/") //切换目录
	dir2, _ := os.Getwd()
	fmt.Printf("dir2: %v\n", dir2)

	tempDir := os.TempDir() //获取临时目录
	fmt.Printf("s: %v\n", tempDir)
}

func renameFile() {
	err := os.Rename("a.txt", "test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func readFile() {
	writeFile()
	b, err := os.ReadFile("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b[:]))
	}

}

func writeFile() {
	s := "brightness"
	os.WriteFile("a.txt", []byte(s), os.ModePerm) //文件不存在会自动创建,并且是覆盖写入
}
