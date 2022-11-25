package sync

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func TestSync() {

	fmt.Println("sync...")
	// test1()
	test2()
	time.Sleep(time.Second * 10)
}

func test2() {
	go responeSize("https://www.baidu.com")
	go responeSize("https://jd.com")

}

func test1() {
	go showMsg("go") //通过，go 关键字开启了一个协程

	showMsg("java")
	/*
		运行结果举例：
		msg: java
		msg: go
		msg: go
		msg: java
		msg: java
		msg: go
		msg: java
		msg: go
	*/
}

func responeSize(url string) {
	fmt.Println("step1: ", url)
	respone, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("step2: ", url)
	defer respone.Body.Close()
	fmt.Println("step3: ", url)
	body, err := ioutil.ReadAll(respone.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("step4: ", len(body))
}

func showMsg(msg string) {
	for i := 0; i < 4; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}
