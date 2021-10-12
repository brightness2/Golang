/*
 * @Author: Brightness
 * @Date: 2021-10-12 09:07:27
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-12 09:14:12
 * @Description:
 */
package models

//定义 user模型

type User struct {
	// gorm.Model //基本模型定义gorm.Model，包括字段ID，CreatedAt，UpdatedAt，DeletedAt，你可以将它嵌入你的模型，或者只写你想要的字段
	Id   int
	Name string
	// CreditCard        CreditCard      // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	UserRoles []UserRole //One-To-Many (拥有多个 - UserRole表的UserID作外键)

}

type UserRole struct {
	Id     int
	UserId int
	RoleId int
}
