package ulog

import (
	"fmt"
	"log"
	"os"
)

/*
log 系列
print 单纯打印
panic 打印日志，抛出panic异常
fatal 打印日志，强制结束程序 os.Exit(1),defer 不会执行

*/

func TestLog() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered ", r)
		}
	}()
	// initLog()
	// printLog()
	// panicLog()
	// fatalLog()
	mylogger()
}

/*
 */

func printLog() {
	log.Print("my log")                       //2022/11/03 16:42:54 my log
	log.Printf("value err,value is %v", "oo") //2022/11/03 16:48:10 value err,value is oo
}

func panicLog() {
	defer fmt.Println("panic 发生后执行")
	log.Panic("my panic")
	fmt.Println("panic 发生后不执行 ")
}

func fatalLog() {
	defer fmt.Println("fatal 后不执行defer")
	log.Fatal("my fatal")
	fmt.Println("fatal 后不执行 ")
}

/****初始化log配置*****/
func initLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("My Log :") //设置前缀
	f, err := os.OpenFile("mylog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	// defer f.Close() //不用close文件，否则输出不了内容到文件
	if err != nil {
		log.Fatal("日志文件错误")
	} else {
		log.SetOutput(f)
	}

}

/*****创建自定义logger*****/
var logger *log.Logger

func mylogger() {
	f, err := os.OpenFile("mylog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	// defer f.Close() //不用close文件，否则输出不了内容到文件
	if err != nil {
		log.Fatal("日志文件错误")
	} else {
		logger = log.New(f, "Success:", log.Ldate|log.Ltime|log.Llongfile)
	}

	logger.Print("diy logger ok")
}
