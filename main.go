package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/ramo798/golang-simpleapi/db"
)

//
// type Blogs []Blog

func main() {
	//dbにアクセス
	db.Init()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
		testread()
	})

	//urlから値を取り出す
	r.GET("/write", func(c *gin.Context) {

		title := c.DefaultQuery("title", "non")
		body := c.DefaultQuery("body", "non")
		come := c.DefaultQuery("come", "non")
		sakusya := c.DefaultQuery("sakusya", "guest")

		// writecon := Blog{}
		// writecon.Title = title
		// writecon.Body = body
		// writecon.Come = come
		// writecon.Sakusya = sakusya

		// db.Create(writecon)

		c.String(200, "%s %s %s %s", title, body, come, sakusya)
	})

	r.Run(":3000")
	// testwrite()
	// testread()
}

func testwrite() {
	//dbにアクセス
	db, err := gorm.Open("postgres", "user=root password=root dbname=DB79 sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	//テーブル作る
	// db.AutoMigrate(Blog{})

	//ここからテスト書き込み
	// testwrite := Blog{}
	// testwrite.Title = "test2"
	// testwrite.Body = "b"
	// testwrite.Come = "bbb"
	// testwrite.Sakusya = "bbbb"
	// db.Create(testwrite)
}

func testread() {
	//dbにアクセス
	// db.Init()
	//テーブル作る
	// db.AutoMigrate(Blog{})

	//読み込んだ
	// kizis := Blogs{}
	// &kizisはkizisに代入されるという意味 .find()は引数の変数に値を入れる
	// db.Debug().Find(&kizis) //SELECT * FROM "blogs"
	// fmt.Println(kizis)
}
