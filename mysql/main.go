/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package main

import (
	"mysql/usage"
)

func main() {
	mysql := usage.NewMysql()
	defer mysql.Close()

	// 查询单条记录
	mysql.QueryRow1()
	mysql.QueryRow2()
	// 查询多条记录
	mysql.QueryAll1()
	mysql.QueryAll2()
	// 插入数据
	mysql.Insert()
	// 更新数据
	mysql.Update()
	// 删除数据
	mysql.Del()
	// 事务处理
	mysql.Tx()
}
