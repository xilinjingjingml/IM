package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("root:123@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema(是数据库对象（表、索引、视图等）的集合。)
	//db.AutoMigrate(&Product{})
	db.AutoMigrate(&models.UserBasic{})

	// Create
	//db.Create(&Product{})
	user := &models.UserBasic{}
	user.Name = "申专"
	db.Create(user)

	//Read
	//var product Product
	fmt.Println(db.First(user, 1)) // 根据整型主建查找
	//db.First(&product, "code = ?", "D42") // 查找code字段值为D42的记录

	//Update - 将product的price更新为200
	db.Model(user).Update("PassWord", "123456")
	// Update- 更多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) //仅更新非零值字段
	//b.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//Delete - 删除 product
	//db.Delete(&product, 1)
}
