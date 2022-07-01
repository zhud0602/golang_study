package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	db = d

}

//创建表
func createtable() {
	db.AutoMigrate(&User{})
}

//创建记录
func insert1() {
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}

//创建记录并更新给出的字段。
func insert2() {
	user2 := User{
		Name:     "lisi",
		Age:      20,
		Birthday: time.Now(),
	}
	db.Select("Name", "Age", "CreatedAt").Create(&user2)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

}

//创建一个记录且一同忽略传递给略去的字段值。 给定的字段不传值，其余的都传递
func insert3() {
	user3 := User{Name: "wangwu", Age: 20, Birthday: time.Now()}
	db.Omit("Name", "Age", "CreatedAt").Create(&user3)
}

//批量插入
func muchInsert() {
	var users = []User{{Name: "jinzhu4"}, {Name: "jinzhu5"}, {Name: "jinzhu6"}}
	db.Create(&users)
	for _, user := range users {
		fmt.Printf("user.id:%v\n", user.ID)
	}
}

//创建钩子
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	return
}

//根据主键删除
func del() {
	db.Delete(&User{}, []int{14, 15, 16})
}

func main() {
	//createtable()
	//insert1()
	//insert2()
	//insert3()
	muchInsert()
	//del()
}
