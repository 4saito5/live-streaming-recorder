package main

import (
	// "log"
	// "net/http"
	// "os"
	"fmt"
  "net/http"
	"io/ioutil"
	// jsoniter "github.com/json-iterator/go"
	"regexp"
	"os/exec"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
	// "github.com/4saito5/live-streaming-recorder/model"

	"time"
)

type CheckList struct {
	ID       int    `json:"id"`
	Group    string `json:"Group"`
	Name     string `json:"Name"`
	Site     string `json:"Site"`
	Key      string `json:"Key"`
	URL      string `json:"URL"`
	IsRecord int    `json:"IsRecord"`
	OnLive   int    `json:"OnLive"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("/data/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
		// return err
	}
	if db == nil {
		panic(err)
		// return err
	}
	var checklists []CheckList


	// Loop
	ticker := time.NewTicker(time.Millisecond * 30000)
	defer ticker.Stop()
	// count := 0
	for {
		select {
		case <-ticker.C:
			// fmt.Println("count=", count)
			// count++
			db.Find(&checklists)
			matchCheckList(db, &checklists)
		}
	}
}

func matchCheckList(db *gorm.DB, checklists *[]CheckList) {
	checkLists := *checklists
	for index, checklist := range checkLists {
		if checklist.IsRecord == 0 {
			continue
		}

		// Check live streaming
		url := "https://www.showroom-live.com/api/live/polling?room_id=" + checklist.Key
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: http.Get, ", checklist.Name, ", skip")
			continue
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		// fmt.Println(bodyString)

		// "is_login":true
		// TODO:上記の正規表現ができない
		reg := regexp.MustCompile(`is_login`)
		if reg.MatchString(bodyString) {
			if checkLists[index].OnLive == 0 {
				fmt.Println("Live Start: ", checklist.Name)
				checkLists[index].OnLive = 1
				db.Save(checkLists[index])
				// yyyy/mm/ ディレクトリに保存 TODO:オプションでカスタマイズできるようにする
				now := time.Now()
				const layout = "2006/01"
				exec.Command("youtube-dl", "-o", "/download/" + now.Format(layout) + "/%(title)s.%(ext)s", checklist.URL).Start()
			} else {
				fmt.Println("Live Streaming: ", checklist.Name)
			}
		} else {
			if checkLists[index].OnLive != 0 {
				fmt.Println("Live End: ", checklist.Name)
				checkLists[index].OnLive = 0
				db.Save(checkLists[index])
			}
		}
	}
}
