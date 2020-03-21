/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package usage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	UserName = "root"
	PassWord = "root"
	Host     = "127.0.0.1"
	Port     = "3306"
	DataBase = "demo"
	Charset  = "utf8"
)

type Mysql struct {
	DB *sql.DB
}

// 用户表结构体
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 构造
func NewMysql() *Mysql{
	db, err := MysqlConnect()
	if err != nil {
		panic(err.Error())
	}
	return &Mysql{DB: db}
}

// mysql连接
func MysqlConnect() (db *sql.DB, err error) {
	// 第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", UserName, PassWord, Host, Port, DataBase, Charset)
	db, err = sql.Open("mysql", dbDSN)
	if err != nil {
		return
	} else {
		fmt.Println("mysql is connected")
	}

	// 最大连接数 实现连接池
	db.SetMaxOpenConns(100)
	// 闲置连接数
	db.SetMaxIdleConns(20)
	// 最大连接周期
	db.SetConnMaxLifetime(100 * time.Second)

	if err = db.Ping(); err != nil {
		return
	}

	return
}

// 关闭连接
func (m *Mysql) Close() {
	m.DB.Close()
}

// 查询单条记录方式一
func (m *Mysql) QueryRow1() {
	var user User
	row := m.DB.QueryRow("select id, name, age from users where id=?", 1)
	if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user.Id, user.Name, user.Age)
}

// 查询单条记录方式二
func (m *Mysql) QueryRow2() {
	rows, _ := m.DB.Query("select id, name, age from users where name=?", "nash")
	if rows == nil {
		return
	}
	var id int64
	var age int
	var name string
	for rows.Next() {
		rows.Scan(&id, &age, &name)
	}
	fmt.Println(id, age, name)
}

// 查询多条记录方式一
func (m *Mysql) QueryAll1() {
	var users []User
	rows, err := m.DB.Query("SELECT * FROM `users` limit ?", 100)
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

// 查下多条记录方式二
func (m *Mysql) QueryAll2() {
	// 查询数据，取所有字段
	rows, _ := m.DB.Query("select * from users")

	// 返回所有列
	cols, _ := rows.Columns()

	// 这里表示一行所有列的值，用[]byte表示
	data := make([][]byte, len(cols))

	// 这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	// 这里scans引用data，把数据填充到[]byte里
	for k, _ := range data {
		scans[k] = &data[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		// 填充数据
		rows.Scan(scans...)
		// 每行数据
		row := make(map[string]string)
		// 把data中的数据复制到row中
		for k, v := range data {
			key := cols[k]
			// 这里把[]byte数据转成string
			row[key] = string(v)
		}
		// 放入结果集
		result[i] = row
		i++
	}
	fmt.Println(result)
}

// 插入数据
func (m *Mysql) Insert() {
	ret, _ := m.DB.Exec("insert INTO users(name, age) values(?, ?)", "nash", 18)

	// 插入数据的主键id
	lastInsertID, _ := ret.LastInsertId()
	fmt.Println("LastInsertID: ", lastInsertID)

	// 影响行数
	rowsAffected, _ := ret.RowsAffected()
	fmt.Println("RowsAffected:", rowsAffected)
}

// 更新数据
func (m *Mysql) Update() {
	ret, _ := m.DB.Exec("UPDATE users set age=? where id=?", "20", 1)
	rowsAffected, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", rowsAffected)
}

// 删除数据
func (m *Mysql) Del() {
	ret, _ := m.DB.Exec("delete from users where id=?", 1)
	rowsAffected, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", rowsAffected)
}

// 事务处理
func (m *Mysql) Tx() {
	// 事务处理
	tx, _ := m.DB.Begin()

	// 新增
	userAddPre, _ := m.DB.Prepare("insert into users(name, age) values(?, ?)")
	addRet, _ := userAddPre.Exec("libby", 18)
	insertRowsAffected, _ := addRet.RowsAffected()

	// 更新
	userUpdatePre1, _ := tx.Exec("update users set name = 'nash'  where id=?", 1)
	updateRowsAffected, _ := userUpdatePre1.RowsAffected()

	fmt.Println(insertRowsAffected, updateRowsAffected)

	if insertRowsAffected > 0 && updateRowsAffected > 0 {
		tx.Commit()
	}else{
		tx.Rollback()
	}
}
