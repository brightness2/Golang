/*
 * @Author: Brightness
 * @Date: 2021-10-11 15:58:42
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-11 16:50:54
 * @Description:
 */
package admin

import "beeWeb/controllers/common"

type LoginController struct {
	common.BaseController
}

func (c LoginController) Get() {
	//使用session 需要配置文件 session=true，开启session
	//假设登录成功，设置session,注意session必须在 beego.BeforeStatic 之后
	c.SetSession("userName", "Brightness")
	c.Ctx.WriteString("登录成功")
}
