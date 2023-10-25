package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/todo?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Veritabanı bağlantısı başarısız: " + err.Error())
	}
	Db = db
}
