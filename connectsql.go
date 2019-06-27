package main

import (
	"fmt"
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

	// var testpost = Blog{
	// 	Title:   "今日は晴れの日3",
	// 	Body:    "晴れてたよ2",
	// 	Come:    "そーなんだ１",
	// 	Sakusya: "sasasa",
	// }
	//これで書き込みできる
	// db.Save(&testpost)

	var kakikomi = Blog{
		Title:   "qqqqsasasa",
		Body:    "sa",
		Come:    "swq",
		Sakusya: "qw",
	}
	//2行で書けばこっちでも書き込みできる。primary　keyがどうとか
	db.Create(kakikomi)
	//インスタンスを作ってcreate()で書き込み
	kakikomimi := Blog{}
	kakikomimi.Title = "タイトル"
	kakikomimi.Body = "a"
	kakikomimi.Come = "aaa"
	kakikomimi.Sakusya = "sasasa"
	db.Create(kakikomimi)

	//読み込みインスタンス作ってからそこにfind()とかで入れる
	blogentry := []Blog{}
	db.Find(&blogentry)
	fmt.Println(blogentry)
}
