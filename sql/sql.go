package sql

import (
	"log"

	"github.com/jinzhu/gorm"
)

//Blogmodel is model
type Blogmodel struct {
	Title  string `json:"Title"`
	Body   string `json:"Body"`
	Author string `json:"Author"`
}

//Init is initialize db from main function
func Init() {
	db, err := gorm.Open("postgres", "user=root password=root dbname=now sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	db.AutoMigrate(Blogmodel{})
}

//TestInit is test
func TestInit() {
	db, err := gorm.Open("postgres", "user=root password=root dbname=nanana sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	db.AutoMigrate(Blogmodel{})

	var momo = Blogmodel{
		Title:  "taitoru",
		Body:   "bodi-",
		Author: "sakusya",
	}

	db.Create(momo)
}
