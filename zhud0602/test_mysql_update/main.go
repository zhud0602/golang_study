package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id       int
	username string
	password string
}

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 查询多条数据示例
func queryMultiRow() {
	sqlStr := "select id, username, password from user_tbl where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d username:%s password:%s\n", u.id, u.username, u.password)
	}
}

func insert() {
	s := "insert into user_tbl(username,password) values(?,?)"
	r, err := db.Exec(s, "zhangsan", "zs123")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		i, _ := r.LastInsertId()
		//l, _ := r.RowsAffected()
		fmt.Printf("i:%v\n", i)
		//fmt.Printf("l:%v\n", l)
	}
}

func insert2(username, password string) {
	s := "insert into user_tbl(username,password) values(?,?)"
	r, err := db.Exec(s, username, password)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		i, _ := r.LastInsertId()
		//l, _ := r.RowsAffected()
		fmt.Printf("i:%v\n", i)
		//fmt.Printf("l:%v\n", l)
	}
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功\n")
	}
	queryMultiRow()
	//insert()
	//insert2("lisi", "lisi123")
}
