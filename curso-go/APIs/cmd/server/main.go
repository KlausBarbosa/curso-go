package main

import (
	"github.com/klausbarbosa/curso-go/curso-go/APIs/configs"
	"github.com/klausbarbosa/curso-go/curso-go/APIs/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	http.ListenAndServe(":8000", nil)
}
