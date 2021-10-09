/*
 * @Author: Brightness
 * @Date: 2021-10-08 16:03:59
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-08 16:24:36
 * @Description:
 */

/*
 *http包执行流程：

创建Listen Socket, 监听指定的端口, 等待客户端请求到来。

Listen Socket接受客户端的请求, 得到Client Socket, 接下来通过Client Socket与客户端通信。

处理客户端的请求, 首先从Client Socket读取HTTP请求的协议头, 如果是POST方法, 还可能要读取客户端提交的数据, 然后交给相应的handler处理请求, handler处理完毕准备好客户端需要的数据, 通过Client Socket写给客户端。

*/
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func webHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息,map[age:[11] name:[Brightness]]
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["name"]) //[Brightness]
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	/*key: age
	val: 11
	key: name
	val: Brightness
	*/
	fmt.Fprint(w, "hello world")

}

func main() {
	http.HandleFunc("/", webHandle)          //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe error:", err)
	}
}
