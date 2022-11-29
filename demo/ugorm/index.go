package ugorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
需要：
安装 mysql
安装 gorm
go get -u gorm.io/driver/mysql
go get -u gorm.io/gorm
*/

func TestIndex() {
	initDB()
	// createData()
	readData()
	// updateData()
	// deleteData()
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func initDB() (err error) {
	user := "root"
	pass := "root"
	database := "demo"
	dsn := user + ":" + pass + "@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建表
	db.AutoMigrate(&Product{}) //会创建数据表product
	/*创建的表结构
	CREATE TABLE `products` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`created_at` datetime(3) DEFAULT NULL,
	`updated_at` datetime(3) DEFAULT NULL,
	`deleted_at` datetime(3) DEFAULT NULL,
	`code` longtext,
	`price` bigint(20) unsigned DEFAULT NULL,
	PRIMARY KEY (`id`),
	KEY `idx_products_deleted_at` (`deleted_at`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	*/
	return err
}

func createData() {
	db.Create(&Product{
		Code:  "1001",
		Price: 100,
	})
}

func readData() {
	var product Product
	// db.First(&product, 1) // 根据整型主键查找
	// fmt.Printf("product: %v\n", product)
	db.First(&product, "code = ?", "1001") // 查找 code 字段值为 1001 的记录
	fmt.Printf("product: %v\n", product)

}

func updateData() {
	var product Product
	db.First(&product, 1) // 根据整型主键查找
	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
}

func deleteData() {
	var product Product

	// Delete - 删除 product,一般是软删除
	// db.Delete(&product, 1)//根据主键删除
	db.First(&product, 2) // 根据整型主键查找
	db.Delete(&product)   //根据查询的记录进行删除

}
