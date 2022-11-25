package main

import (
	"fmt"
	// "demo/sync"
	// "demo/uos"
	// "demo/ulog"
	// "demo/ubuiltin"
	// "demo/encode"
	// "demo/mysql"
	"demo/ugorm"
)

func main() {
	// sync.TestSync()
	// sync.TestChan()
	// sync.TestWait()
	// sync.TestTools()
	// sync.TestAtomic()
	// uos.TestFile()
	// uos.TestFile2()
	// ulog.TestLog()
	// ubuiltin.TestBuiltin()
	// encode.TestJson()
	// encode.TestXml()
	// mysql.TestIndex()
	ugorm.TestIndex()
	fmt.Println("main end...") //主函数结束，其它协程也会结束
}
