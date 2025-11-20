package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//gorm demo05

// 1.定义模型
type User struct {
	gorm.Model //ID CreateAt UpdateAt DeleteAt
	Name       string
	Age        int64
	Active     bool
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
	// u1 := User{Name: "jackie", Age: 18, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "jackie2", Age: 20, Active: false}
	// db.Create(&u2)
	//5.查询
	var user User
	db.First(&user)
	//6.更新
	user.Name = "jackie666"
	user.Age = 99
	db.Debug().Save(&user) //默认会修改所有字段
	// db.Debug().Model(&user).Update("name", "jackie777")

	// m1 := map[string]interface{}{
	// 	"name":   "renhaokun",
	// 	"age":    24,
	// 	"active": true,
	// }

	// db.Debug().Model(&user).Updates(m1)                //m1列出来的所有字段都会更新
	// db.Debug().Model(&user).Select("age").Update(m1)   //只更新age字段
	// db.Debug().Model(&user).Omit("active").Updates(m1) //排除m1中的active字段更新其他的字段
	// db.Debug().Model(&user).UpdateColumn("age", 30)
	// rowsNum := db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected
	// fmt.Println(rowsNum)

	//让users表中所有的用户的年龄在原来的基础上+2
	db.Model(&User{}).Update("age", gorm.Expr("age + ?", 2))
}
