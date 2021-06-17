package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func init()  {
	connectDB()
}

func main() {
	//add()
	//find()
	//update()
	delete()
}

var db *gorm.DB

func connectDB()  {
	var err error
	db,err = gorm.Open("mysql", "root:123456@/txtv?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connect fail" + err.Error())
	}
	fmt.Println("connect successful")
	db.DB().SetMaxIdleConns(1000)
	db.DB().SetConnMaxLifetime(time.Second * 10)
	db.DB().SetMaxOpenConns(1000)
}

type User struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(20) column:'name'"`
	UserId string `gorm:"type:varchar(20); unique; notnull column:'user_id'"`
	Pwd string `gorm:"type:varchar(32) notnull column:'pwd'"`
	Email string `gorm:"type:varchar(40) unique; notnull column:'email'"`
	CreationDate time.Time `gorm:"autoCreateTime"`
	ModificationDate time.Time `gorm:"autoUpdateTime"`
	Version int `gorm:"type:int column: version"`
}

func (User) TableName() string {
	return "user"
}

func add()  {
	db.Debug().Create(&User{
		Name: "momoo",
		UserId: "sssssssss",
		Pwd: "123456",
		Email: "123@163.com",
	})
}

func find()  {
	uu := []User{}
	db.Debug().Find(&uu)
	//db.Row() 原生sql
	fmt.Println(uu)
}

func update()  {
	db.Debug().Model(&User{}).Where("pwd = ?", "123456").Update("pwd","111222")
}

func delete()  {
	db.Debug().Where("pwd = ? ", "111222").Delete(&User{})
}
