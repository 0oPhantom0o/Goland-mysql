package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// CREATE USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
// CREATE USER 'phantom'@'localhost' IDENTIFIED BY 'bookstore';
// Create table book(ID int,Name varchar(64),author varchar(64),publication varchar(64));

func Connect() {
	d, err := gorm.Open("mysql", "phantom:bookstore/BookStore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
