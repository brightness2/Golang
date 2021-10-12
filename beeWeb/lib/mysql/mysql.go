/*
 * @Author: Brightness
 * @Date: 2021-10-12 08:56:26
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-12 09:03:02
 * @Description:
 */
package mysql

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() {
	user, _ := beego.AppConfig.String("db_user")
	pass, _ := beego.AppConfig.String("db_password")
	host, _ := beego.AppConfig.String("db_host")
	port, _ := beego.AppConfig.String("db_port")
	database, _ := beego.AppConfig.String("db_database")
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, database)
	db, err := gorm.Open("mysql", source)
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
