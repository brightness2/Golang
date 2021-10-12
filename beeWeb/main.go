/*
 * @Author: Brightness
 * @Date: 2021-10-09 08:57:27
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-12 08:51:46
 * @Description:
 */
package main

import (
	"beeWeb/controllers/admin"
	"beeWeb/controllers/common"
	"beeWeb/lib/mysql"
	_ "beeWeb/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	mysql.Init()
}
func main() {
	beego.ErrorController(&common.ErrorController{})

	//过滤器
	beego.InsertFilter("/admin/*", beego.BeforeRouter, admin.RouteFilter)

	beego.Run()

}
