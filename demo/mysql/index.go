package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/**
mysql 操作
1、安装配置mysql驱动
go get -u github.com/go-sql-driver/mysql
2、
*/

var db *sql.DB

func init() {
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
}
func TestIndex() {
	// insert("Brightness", 20)
	// find(4)
	// selects()
	// update(1, "test", 18)
	delete(2)
}

func initDB() (err error) {
	user := "root"
	pass := "root"
	database := "demo"
	dsn := user + ":" + pass + "@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True"
	//不会校验账号密码是否正确
	db, err = sql.Open("mysql", dsn) //赋值给全局 db变量
	if err != nil {
		return err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	//尝试连接数据库，校验dsn是否正确
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil

}

func insert(name string, age int) {
	sqlStr := "INSERT INTO `test`(`name`,`age`) VALUES(?,?)"
	r, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("insert ID: %v\n", i)
	}
}

type User struct {
	id   int
	name string
	age  int
}

func find(id int) {
	sqlStr := "SELECT * FROM `test` WHERE `id` = ?"
	r := db.QueryRow(sqlStr, id)
	var user User
	err := r.Scan(&user.id, &user.name, &user.age)
	if err != nil {
		fmt.Printf("err: %v\n", err) //比如sql: no rows in result set
		return
	}
	fmt.Printf("user: %v\n", user)
}

func selects() {
	sqlStr := "SELECT * FROM `test`"
	r, err := db.Query(sqlStr)
	defer r.Close() //注意要关闭r
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	var u User
	for r.Next() {
		r.Scan(&u.id, &u.name, &u.age)
		fmt.Printf("u: %v\n", u)
	}

}

func update(id int, name string, age int) {
	sqlStr := "UPDATE `test` set `name`=?,`age`=? WHERE `id`=?"
	r, err := db.Exec(sqlStr, name, age, id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	i, _ := r.RowsAffected()
	fmt.Println("影响的行数:", i)

}

func delete(id int) {
	sqlStr := "DELETE FROM `test` WHERE `id`=?"
	r, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	i, _ := r.RowsAffected()
	fmt.Println("影响的行数:", i)

}
