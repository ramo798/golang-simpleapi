package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Blog struct {
	Title string `json:"Title"`
	Body  string `json:"Body"`
	Come  string `json:"Come"`
}

type Blogs []Blog

func main() {
	db, err := gorm.Open("postgres", "user=root password=root dbname=now sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	db.AutoMigrate(Blog{})
	var testpost = Blog{
		Title: "今日は晴れの日",
		Body:  "晴れてたよ",
		Come:  "そーなんだ",
	}
	db.NewRecord(testpost)
	db.Create(&testpost)
	db.Save(&testpost)

	var Blogs = Blog{}
	db.Find(&Blogs)
	fmt.Println(Blogs)

}
