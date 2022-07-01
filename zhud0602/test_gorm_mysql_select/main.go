package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

type Language struct {
	Code string
	Name string
}

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	db = d

}

func select1() {
	var user User
	//db.First(&user)
	//db.Take(&user)
	result := db.Last(&user)
	fmt.Printf("user.ID:%v\n", user.ID)
	fmt.Printf("affected:%v\n", result.RowsAffected)
	fmt.Printf("result.Error:%d\n", result.Error)
}

func select2() {
	//result := map[string]interface{}{}
	//db.Model(&User{}).First(&result)

	//result := map[string]interface{}{}
	//db.Table("users").Take(&result)

	result := map[string]interface{}{}
	db.Model(&User{}).Last(&result)
	fmt.Printf("result:%v\n", result["name"])
}

func select3() {
	//var language Language
	//db.AutoMigrate(&Language{})
	//db.Create(&Language{Code: "a01", Name: "JYJY"})
	//db.First(&Language{})

	//db.First(&language)
	//fmt.Println(language)

	var language Language
	db.AutoMigrate(&language)
	db.Create(&Language{Code: "a01", Name: "JYJY"})

}

func select4() {
	//var user User

	//db.First(&user, 10)
	//db.First(&user, "10")
	var users []User
	db.Find(&users, []int{1, 2, 3})
	fmt.Printf("users:%v\n", users)
}

func select5() {
	//var user User
	var users []User
	db.First(&users, "name = ?", "jinzhu4")
	//SELECT * FROM users WHERE name = "jinzhu4" limit 1;
	fmt.Printf("%v\n", users)
}

func select6() {
	var users []User
	db.Find(&users, "name=?", "jinzhu4")
	//fmt.Printf("%v\n", users)

	for _, user := range users {
		fmt.Printf("user.id:%v\n", user.ID)
	}
}

func select7() {
	var user = User{ID: 10}
	db.First(&user)
	fmt.Printf("user.name:%v\n", user.Name)

}

func select8() {
	var user = User{ID: 10}
	result := map[string]interface{}{}
	db.Model(user).First(&result)
	// SELECT * FROM users WHERE id = 10;
	fmt.Printf("result:%v\n", result["name"])
}

//检索全部对象
func selectAll() {
	var users []User
	// Get all records
	result := db.Find(&users)
	// SELECT * FROM users;

	for i, user := range users {
		fmt.Printf("id:%v  name:%v\n", i, user.Name)
	}

	fmt.Printf("RowsAffected:%v\n", result.RowsAffected)
	// returns found records count, equals `len(users)`
	fmt.Printf("result.Error:%v\n", result.Error)
	// returns error
}

//string 条件查询
func selectString() {
	//var user User
	var users []User

	// Get first matched record
	//db.Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	//fmt.Printf("created_at:%v\n", user.CreatedAt)

	// Get all matched records
	//db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	//db.Where("name IN ?", []string{"jinzhu", "jinzhu4"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	//db.Where("name LIKE ?", "%j%z%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	//db.Where("name like ? AND age >= ?", "%jinzhu%", "1").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	//db.Where("updated_at > ?", time.Now().AddDate(0, 0, -1)).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", time.Now().Add(time.Hour*(-24)), time.Now()).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	for _, user := range users {
		fmt.Printf(" ID:%v\n", user.ID)
	}
}

//Struct & Map 条件
func selectStruct() {
	var user User
	//var users []User

	// Struct
	//db.Where(&User{Name: "jinzhu4", Age: 0}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	//fmt.Printf("id:%v\n", user.ID)

	// Map
	//db.Where(map[string]interface{}{"name": "jinzhu4", "age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	//for _, user := range users {
	//	fmt.Printf("id:%v\n", user)
	//}

	// Slice of primary keys
	//db.Where([]int64{1, 4, 10}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);
	//for _, user := range users {
	//	fmt.Printf("user:%v\n", user)
	//}

	// Get by primary key if it were a non-integer type
	db.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

}

func selectNot() {
	var user User
	//var users []User

	// Not In
	//db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu4"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
	//for _, user := range users {
	//	fmt.Printf("%v\n", user)
	//}

	//db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

	// Struct
	//db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// Not In slice of primary keys
	db.Not([]int64{1, 2, 4}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;

	fmt.Printf("%v\n", user)
}

//Or 条件
func selectOr() {
	var users []User

	//db.Where("name = 'Jinzhu'").Or("name ='jinzhu4'").Find(&users)
	db.Where("name = ?", "jinzhu").Or("name =?", "jinzhu4").Find(&users)
	// SELECT * FROM users WHERE name = 'Jinzhu' OR name = 'jinzhu4';

	// Struct
	//db.Where("name = 'jinzhu4'").Or(User{Name: "jinzhu", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// Map
	//db.Where("name = 'jinzhu4'").Or(map[string]interface{}{"name": "jinzhu", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
}

//Order
func selectOrder() {
	var users []User
	//db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// Multiple orders
	//db.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)

	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
}

//Limit & Offset
func selectLimit() {
	var users []User
	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	// Cancel limit condition with -1
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	//db.Offset(3).Find(&users)
	// SELECT * FROM users OFFSET 3;

	//db.Limit(10).Offset(5).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// Cancel offset condition with -1
	//db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)

	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
}

func main() {
	//select1()
	//select2()
	//select3()
	//select4()
	//select5()
	//select6()
	//select7()
	//select8()
	//selectAll()
	//selectString()
	//selectStruct()
	//selectNot()
	//selectOr()
	//selectOrder()
	selectLimit()
}
