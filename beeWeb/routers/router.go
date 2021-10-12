/*
 * @Author: Brightness
 * @Date: 2021-10-09 08:57:27
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-11 16:40:23
 * @Description:
 */
package routers

import (
	"beeWeb/controllers/admin"
	"beeWeb/controllers/api"
	"beeWeb/controllers/index"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	/***********index*********************/

	beego.Router("/", &index.IndexController{})
	beego.Router("/list", &index.IndexController{}, "get:List")

	beego.Router("adminLogin", &admin.LoginController{})
	/***********admin************/
	admin := beego.NewNamespace("/admin",

		beego.NSRouter("/", &admin.IndexController{}),
	)
	beego.AddNamespace(admin)

	/***********api***************/
	api := beego.NewNamespace("/api",
		beego.NSRouter("/", &api.IndexController{}),
	)
	beego.AddNamespace(api)

}
