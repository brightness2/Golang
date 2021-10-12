/*
 * @Author: Brightness
 * @Date: 2021-10-09 17:03:21
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-12 09:25:03
 * @Description:
 */
package index

import (
	"beeWeb/controllers/common"
)

type IndexController struct {
	common.BaseController
}

func (c *IndexController) Get() {

	// id, _ := c.GetInt("id")
	// user, err := models.ReadUser(id)
	// user, err := models.UpdateUser(models.User{Id: 2, Name: "test5555"})
	// user, err := models.CreateUser(models.User{Name: "demo1"})
	// user, err := models.DeleteUser(6)
	// user, err := models.ReadUserAll()

	// if err != nil {
	// 	c.Ctx.WriteString(err.Error())
	// } else {
	// 	c.Data["user"] = user
	// }
	/*********************************************/

	c.TplName = "index/index.html"
}

/*
自定义路由
*/
func (c *IndexController) List() {

	// users, _ := models.ReadUserAll()
	// c.Data["user"] = users
	c.TplName = "index/list.html"

}
