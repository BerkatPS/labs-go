package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-jwt"))

	if err != nil {
		fmt.Println("Gagal Koneksi Ke database...")
	}
	// akan menjalankan auto migrate
	// mengirim struct user
	db.AutoMigrate(&User{})

	// memasukan db dari connection ke db gorm nya.
	DB = db
}
