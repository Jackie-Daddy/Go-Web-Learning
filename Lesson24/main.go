package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1.定义模型
type User struct {
	// gorm.Model // ID CreateAt UpdateAt DeleteAt
	ID     int64
	Name   string
	Age    int64
	Active bool
}

func main() {
	//2.连接MySQL数据库
	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//3.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	//4.创建
	// u1 := User{Name: "renhaokun2", Age: 24, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "jackie2", Age: 25, Active: false}
	// db.Create(&u2)

	//5.删除
	// var u = User{}
	// // u.ID = 1
	// u.Name = "jackie2"
	// db.Debug().Delete(&u)
	db.Debug().Where("name = ?", "jackie2").Delete(User{})
	// db.Delete(User{}, "age = ?", 24)

	// var u1 []User
	// db.Debug().Where("name = ?", "jackie2").Find(&u1)
	// fmt.Println(u1)

	//物理删除
	// db.Debug().Unscoped().Where("name = ?", "jackie2").Delete(User{})
}
