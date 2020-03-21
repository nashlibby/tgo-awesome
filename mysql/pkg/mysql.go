/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package pkg

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

func NewMysql() *Mysql{
	return &Mysql{}
}

func (m *Mysql) Connect() (db *sql.DB, err error) {
	// 第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", UserName, PassWord, Host, Port, DataBase, Charset)
	db, err = sql.Open("mysql", dbDSN)
	if err != nil {
		return
	} else {
		fmt.Println("mysql connected")
	}

	// 最大连接数
	db.SetMaxOpenConns(100)
	// 闲置连接数
	db.SetMaxIdleConns(20)
	// 最大连接周期
	db.SetConnMaxLifetime(100*time.Second)

	if err = db.Ping(); err != nil {
		return
	}

	return
}
