package main

import (
	// "log"
	// "net/http"
	// "os"
	"fmt"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
	// "github.com/4saito5/live-streaming-recorder/model"
)

type CheckList struct {
	ID       int    `json:"id"`
	Group    string `json:"Group"`
	Name     string `json:"Name"`
	Site     string `json:"Site"`
	URL      string `json:"URL"`
	IsRecord int    `json:"IsRecord"`
}

func main() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("/data/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}


  var checklists []CheckList
  // var checklists []model.CheckList
	db.Find(&checklists)

	fmt.Println(checklists)
}
