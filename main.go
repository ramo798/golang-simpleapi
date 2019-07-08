package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Blog struct {
	Title   string `json:"Title"`
	Body    string `json:"Body"`
	Come    string `json:"Come"`
	Sakusya string `json:"Sakusya"`
}

type Blogs []Blog

func main() {
	Initdb()
}

//Initdb is テーブル作成
func Initdb() {
	//dbにアクセス
	db, err := gorm.Open("postgres", "user=root password=root dbname=DB79 sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	//テーブル作る
	db.AutoMigrate(Blog{})
}
