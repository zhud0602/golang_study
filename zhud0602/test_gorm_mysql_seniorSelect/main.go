package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	db = d
}

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type APIUser struct {
	ID   uint
	Name string
}

func create() {
	db.AutoMigrate(&User{}, &APIUser{})
}

func insert() {
	user := User{Name: "no1", Age: 10}
	db.Create(&user)
}

func selectSenior() {
	db.Model(&User{}).Limit(1).Find(&APIUser{})
}

func main() {
	//create()
	//insert()
	selectSenior()

}
