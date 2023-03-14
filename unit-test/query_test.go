package unit_test

import (
	"fmt"
	"golang-fiber/model/entity"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestQuerys(t *testing.T) {
	const MYSQL = "root:root@tcp(localhost:8889)/belajar_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	dbna, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var results []entity.User
	dbna.Where("1 = ?", 1).Find(&results)
	if len(results) < 1 {
		fmt.Println("gagal query")

	}
	fmt.Println("banyak data:", len(results))
	fmt.Println("berhasil query")
}
