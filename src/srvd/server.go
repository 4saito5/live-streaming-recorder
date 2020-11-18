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
)

type CheckList struct {
	ID       int    `json:"id"`
	Group    string `json:"Group"`
	Name     string `json:"Name"`
	Site     string `json:"Site"`
	Key      string `json:"Key"`
	URL      string `json:"URL"`
	IsRecord int    `json:"IsRecord"`
}


// type OnLives struct {
// 	CornerImagePath string `json:"corner_image_path"`
// 	Onlives         []struct {
// 		GenreID     int           `json:"genre_id"`
// 		Banners     []interface{} `json:"banners"`
// 		HasUpcoming bool          `json:"has_upcoming"`
// 		GenreName   string        `json:"genre_name"`
// 		Lives       []struct {
// 			RoomURLKey       string `json:"room_url_key"`
// 			OfficialLv       int    `json:"official_lv"`
// 			FollowerNum      int    `json:"follower_num"`
// 			StartedAt        int    `json:"started_at"`
// 			LiveID           int    `json:"live_id"`
// 			IsFollow         bool   `json:"is_follow"`
// 			StreamingURLList []struct {
// 				IsDefault bool   `json:"is_default"`
// 				URL       string `json:"url"`
// 				Label     string `json:"label"`
// 				ID        int    `json:"id"`
// 				Type      string `json:"type"`
// 				Quality   int    `json:"quality"`
// 			} `json:"streaming_url_list"`
// 			LiveType        int           `json:"live_type"`
// 			Tags            []interface{} `json:"tags"`
// 			Image           string        `json:"image"`
// 			ViewNum         int           `json:"view_num"`
// 			GenreID         int           `json:"genre_id"`
// 			MainName        string        `json:"main_name"`
// 			PremiumRoomType int           `json:"premium_room_type"`
// 			CellType        int           `json:"cell_type"`
// 			BcsvrKey        string        `json:"bcsvr_key"`
// 			RoomID          int           `json:"room_id"`
// 			Telop           string        `json:"telop,omitempty"`
// 		} `json:"lives"`
// 	} `json:"onlives"`
// 	BcsvrPort int    `json:"bcsvr_port"`
// 	BcsvrHost string `json:"bcsvr_host"`
// }

func main() {
	var checklists []CheckList
	getCheckList(&checklists)
	// // fmt.Println()
	// for _, checklist := range checklists {
	// 	fmt.Println(checklist.Site)
	// }

	matchCheckList(&checklists)

}

func getCheckList(checklists *[]CheckList) {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("/data/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
		// return err
	}
	if db == nil {
		panic(err)
		// return err
	}


	// var checklists []model.CheckList
	db.Find(checklists)

	// fmt.Println(checklists)

	// return checklists

}



func matchCheckList(checklists *[]CheckList) {

	for _, checklist := range *checklists {
		// fmt.Println(checklist.Name)
		if checklist.IsRecord == 0 {
			continue
		}

		// 配信中かチェック
		url := "https://www.showroom-live.com/api/live/polling?room_id=" + checklist.Key
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		fmt.Println(bodyString)

		// "is_login":true
		// TODO:上記の正規表現ができない
		reg := regexp.MustCompile(`is_login`)
		if reg.MatchString(bodyString) {
			fmt.Println("配信中：", checklist.Name)
			// err := exec.Command("youtube-dl", checklist.URL).Run()
			// if err != nil {
      //   return err
			// }
			exec.Command("youtube-dl", checklist.URL).Run()
		}

	}

}



// func matchCheckList(checklists *[]CheckList) {

// 	// for _, checklist := range *checklists {
// 	// 		fmt.Println(checklist.Name)

// 	// }

// 	// const response = await axios.get('/api/live/onlives');

//   url := "https://www.showroom-live.com/api/live/onlives"

//   resp, _ := http.Get(url)
//   defer resp.Body.Close()

//   body, _ := ioutil.ReadAll(resp.Body)
//   // fmt.Println(string(byteArray)) // htmlをstringで取得

// 	bytes := []byte(body)

// 	var onlives OnLives
// 	var json = jsoniter.ConfigCompatibleWithStandardLibrary
// 	json.Unmarshal(bytes, &onlives)

// 	// fmt.Printf("%d: %s", resp.StatusCode, onlives.CornerImagePath)
// 	for _, onlive := range onlives.Onlives {
// 		if onlive.GenreName == "Popularity" {
// 			for _, live := range onlive.Lives {
// 				// fmt.Println(live.RoomURLKey)
// 				for _, checklist := range *checklists {
// 					if live.RoomURLKey == checklist.Site {
// 						fmt.Println(checklist.Name, live.RoomID)
// 					}
// 				}
// 			}
// 		} else {
// 			break
// 		}
// 	}

// }

// func matchCheckList(checklists *[]CheckList) {
// 	// const response = await axios.get('/api/live/onlives');

//   url := "https://www.showroom-live.com/api/live/onlives"

//   resp, _ := http.Get(url)
//   defer resp.Body.Close()

//   body, _ := ioutil.ReadAll(resp.Body)
//   // fmt.Println(string(byteArray)) // htmlをstringで取得

// 	bytes := []byte(body)

// 	var onlives OnLives
// 	var json = jsoniter.ConfigCompatibleWithStandardLibrary
// 	json.Unmarshal(bytes, &onlives)

// 	// fmt.Printf("%d: %s", resp.StatusCode, onlives.CornerImagePath)
// 	for _, onlive := range onlives.Onlives {
// 		if onlive.GenreName == "Popularity" {
// 			for _, live := range onlive.Lives {

// 				fmt.Println(live.RoomURLKey)
// 			}

// 		} else {
// 			break
// 		}
// 	}

// }
