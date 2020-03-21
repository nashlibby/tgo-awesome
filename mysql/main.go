/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package main

import (
	"database/sql"
	"fmt"
	"mysql/pkg"
)

// 用户表结构体
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 查询单条记录
func QueryRow(db *sql.DB) {
	var user User
	row := db.QueryRow("select id, name, age from users where id=?", 1)
	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user.Id, user.Name, user.Age)
}

// 查询多条记录
func QueryAll(db *sql.DB) {
	var users []User
	rows, err := db.Query("SELECT * FROM `users` limit ?", 100)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 遍历
	var user User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Age)
		users = append(users, user)
	}
	fmt.Println(users)
}

// 插入数据
func Insert(db *sql.DB) {
	ret, _ := db.Exec("insert INTO users(name, age) values(?, ?)", "nash", 18)

	// 插入数据的主键id
	lastInsertID, _ := ret.LastInsertId()
	fmt.Println("LastInsertID: ", lastInsertID)

	// 影响行数
	rowsAffected, _ := ret.RowsAffected()
	fmt.Println("RowsAffected:", rowsAffected)
}

// 更新数据
func Update(db *sql.DB) {
	ret, _ := db.Exec("UPDATE users set age=? where id=?", "20", 1)
	rowsAffected, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", rowsAffected)
}

// 删除数据
func Del(db *sql.DB) {
	ret, _ := db.Exec("delete from users where id=?", 1)
	rowsAffected, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", rowsAffected)
}


func main() {
	mysql := pkg.NewMysql()
	db, err := mysql.Connect()
	// 关闭数据库
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	// 查询单条记录
	QueryRow(db)
	// 查询多条记录
	QueryAll(db)
	// 插入数据
	Insert(db)
	// 更新数据
	Update(db)
	// 删除数据
	Del(db)
}
