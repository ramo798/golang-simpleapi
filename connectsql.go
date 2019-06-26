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
	db, err := gorm.Open("postgres", "user=root password=root dbname=now sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	//テーブル作る
	db.AutoMigrate(Blog{})

	var testpost = Blog{
		Title:   "今日は晴れの日3",
		Body:    "晴れてたよ2",
		Come:    "そーなんだ１",
		Sakusya: "sasasa",
	}
	//これで書き込みできる
	db.Save(&testpost)

	//この2行要調査
	// db.NewRecord(testpost)
	// db.Create(&testpost)

	// var Blogs = Blog{}
	// db.Find(&Blogs)
	// fmt.Println(Blogs)

}
