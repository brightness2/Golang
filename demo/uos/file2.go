package uos

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func TestFile2() {
	// writeAppend()
	// checkEFilexist()
	// listDirs("./", "")
	// readOps()
	// writeOps()
	// useIoutil()
	readByBuf()
	// writeByBuf()
}

func writeAppend() {

	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer f.Close()
	f.WriteString("\nhello world!")
}

func checkEFilexist() {
	_, err := os.Stat("test")
	if err != nil {
		fmt.Println("文件夹或文件不存在")
	} else {
		fmt.Println("文件夹或文件存在")
	}

}

/*
*
遍历文件夹
*/
func listDirs(name string, base string) {
	de, err := os.ReadDir(name) //只列出当前目录层级的文件和目录
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		if len(base) > 0 {
			base += "/"
		}
		for _, f := range de {

			fmt.Println(base + f.Name())
			if f.IsDir() {
				listDirs(f.Name(), f.Name())
			}
		}
	}

}

func readOps() {
	f, err := os.OpenFile("a.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer f.Close()
	for {
		buffer := make([]byte, 3) //需要在循环中声明
		_, err2 := f.Read(buffer)
		if err2 == io.EOF {
			break
		}
		// fmt.Printf("n: %v\n", n)
		fmt.Printf("%v", string(buffer))
	}
	buffer := make([]byte, 6)
	f.ReadAt(buffer, 10) //从某个偏移量开始读
	fmt.Println(string(buffer))

	f.Seek(10, 0) //定位光标
	f.Read(buffer)
	fmt.Println(string(buffer))
	fmt.Println("")
}

func writeOps() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, os.ModePerm) //第二个参数可选os.O_APPEND 追加，os.O_TRUNC 覆盖
	defer f.Close()

	f.Write([]byte("hello "))      //默认从首位开始写
	f.WriteAt([]byte("golang"), 6) //从某个偏移量开始写,不能和os.O_APPEND共用
	f.Seek(12, 0)
	f.WriteString("\nbrightness\n") //从当前偏移量写入字符串，比如现在是从12位置开始的

}

/***************** 下面是IO包的ioutil工具包 ********************************/
//此包已在Go1.16版本后弃用，相同的功能现在由包io或包os提供
/*
func useIoutil() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer f.Close()
	s, _ := ioutil.ReadAll(f) //读取所有内容

	fmt.Printf("%s", s)

	str, _ := ioutil.ReadFile("a.txt") //读取文件内容
	fmt.Printf("string: %s", str)

	files, _ := ioutil.ReadDir(".") //遍历所有文件

	for _, file := range files {
		fmt.Println(file.Name())
	}

	// ioutil.WriteFile("a.txt", []byte("ok"), 0644) //覆盖写入
}
*/

/*********** 结合bufio 进行文件内容操作提高性能且不需要自己管理buffer变量************/
//
func readByBuf() {
	f, _ := os.Open("a.txt")
	defer f.Close()
	r := bufio.NewReader(f)
	s, _ := r.ReadString('\r') //参数是分隔符
	fmt.Printf("s: %v\n", s)
	//
	f2, _ := os.Open("a.txt")
	defer f2.Close()
	r2 := bufio.NewReader(f2)
	line, isPrefix, _ := r2.ReadLine() //读一行
	fmt.Printf("line: %q\n", line)
	fmt.Printf("isPrefix: %v\n", isPrefix)
	line2, isPrefix2, _ := r2.ReadLine()
	fmt.Printf("line: %q\n", line2)
	fmt.Printf("isPrefix: %v\n", isPrefix2)
	//
	str := strings.NewReader("abc,efg,hij")
	br := bufio.NewReader(str)
	w, _ := br.ReadSlice(',') //以","分割的，包含“,”
	fmt.Printf("w: %q\n", w)
	w, _ = br.ReadSlice(',')
	fmt.Printf("w: %q\n", w)
	//

	str2 := strings.NewReader("abc efg hij")
	bs := bufio.NewScanner(str2)
	bs.Split(bufio.ScanWords) //以空格作为分割符
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}

func writeByBuf() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer f.Close()
	buf := bufio.NewWriter(f)
	buf.WriteString("\nhello lang")
	// buf.Reset(f) //重置清空缓冲区，这样hello lang 就不会写入了
	// buf.WriteString("\nok")
	buf.Flush() //必须刷新缓冲区才能写入

}
