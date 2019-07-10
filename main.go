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
	testwrite()
	testread()
}

func testwrite() {
	//dbにアクセス
	db, err := gorm.Open("postgres", "user=root password=root dbname=DB79 sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	//テーブル作る
	db.AutoMigrate(Blog{})

	//ここからテスト書き込み
	testwrite := Blog{}
	testwrite.Title = "test2"
	testwrite.Body = "b"
	testwrite.Come = "bbb"
	testwrite.Sakusya = "bbbb"
	db.Create(testwrite)
}

func testread() {
	//dbにアクセス
	db, err := gorm.Open("postgres", "user=root password=root dbname=DB79 sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	//テーブル作る
	db.AutoMigrate(Blog{})

	//読み込んだ
	kizis := Blogs{}
	// &kizisはkizisに代入されるという意味 .find()は引数の変数に値を入れる
	db.Debug().Find(&kizis) //SELECT * FROM "blogs"
	fmt.Println(kizis)
}
