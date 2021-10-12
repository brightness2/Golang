/*
 * @Author: Brightness
 * @Date: 2021-10-11 16:09:19
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-11 16:39:46
 * @Description:
 */
package admin

// import "github.com/beego/beego/v2/adapter/context"

import (
	"github.com/beego/beego/v2/server/web/context"
)

func RouteFilter(ctx *context.Context) {
	//使用session 需要配置文件 session=true，开启session
	username := ctx.Input.Session("userName")

	if username == nil {
		ctx.WriteString("你还没有登录")
	}
}
