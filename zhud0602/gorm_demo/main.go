package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//模型 Model
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// 在数据库中创建Product表
func create(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

//在表中添加数据
func insert(db *gorm.DB) {
	db.Create(&Product{Code: "D42", Price: 100})
}

//根据主键查询1条记录
func slct(db *gorm.DB) {
	var product Product
	db.First(&product, 1)
	fmt.Printf("product:%v\n", product)
}

//根据条件查询数据，code=D42的记录
func slctmuch(db *gorm.DB) {
	var product Product
	db.First(&product, "code=?", "D42")
	fmt.Printf("product: %v\n", product)
}

//更新1个字段
func update(db *gorm.DB) {
	var product Product
	db.First(&product, 1)
	db.Model(&product).Update("price", 200)
	fmt.Printf("product:%v", product)
}

//更新多个字段
func updates(db *gorm.DB) {
	var product Product
	db.First(&product, 1)
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	fmt.Printf("product:%v", product)
}

//根据主键删除1条数据
func del(db *gorm.DB) {
	var product Product
	db.Delete(&product, 4)
	fmt.Printf("product:%v", product)

}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	//update(db)
	//updates(db)
	del(db)
}
