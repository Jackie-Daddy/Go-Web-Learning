package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model //ID CreateAt UpdateAt DeleteAt
	Name       string
	Age        int64
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	//3.创建
	// u1 := User{Name: "jackie", Age: 18}
	// db.Create(&u1)
	// u2 := User{Name: "jackie2", Age: 20}
	// db.Create(&u2)

	//4.查询
	// var user User   //声明结构体类型变量user
	// user := new(User) //new和make的区别
	// db.First(&user)   //查询表中第一条数据保存到user中
	// fmt.Printf("user: %#v\n", user)

	// var users []User
	// db.Debug().Find(&users)
	// fmt.Printf("users: %#v\n", users)

	//FirstOrInit
	var user User
	// db.Attrs(User{Age: 99}).FirstOrInit(&user, User{Name: "aaa"})
	db.Assign(User{Age: 99}).FirstOrInit(&user, User{Name: "jackie"})
	fmt.Printf("user: %#v\n", user)
}
