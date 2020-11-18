package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"regexp"
	"github.com/4saito5/live-streaming-recorder/graph/model"

	"github.com/4saito5/live-streaming-recorder/graph"
	"github.com/4saito5/live-streaming-recorder/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("/data/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	// defer func() {
	// 	if db != nil {
	// 		if err := db.Close(); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }()


	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	// http.Handle("/fuga", &f)
	// http.Handle("/setURL/:url", methodHandler{"GET": http.HandlerFunc(handler.GetUser)})
	http.HandleFunc("/add_url", addUrlHandler)


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type addUrlJSON struct {
	Url string `json:"url"`
}

func addUrlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Credentials", "false")
	switch r.Method {
	// case http.MethodGet:
	// 	w.WriteHeader(http.StatusOK)
	// 	fmt.Fprint(w, "GET hello!\n")
	case http.MethodPost:
	// Request body
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		var addUrl addUrlJSON
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(buf.Bytes(), &addUrl)

		// SHOWROOM
		reg := regexp.MustCompile(`https://www\.showroom-live\.com/`)
		if reg.MatchString(addUrl.Url) {
			addShowroom(addUrl.Url)
		}

		// ...



		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Add URL Success!\n")
	case http.MethodOptions:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "OPTIONS Success!\n")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed.\n")
	}
	// fmt.Println("debug 99")

}

func addShowroom(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	// get name & growp
	reg := regexp.MustCompile(`<title>(.*) - SHOWROOM`)
	findString := reg.FindStringSubmatch(bodyString)
	// fmt.Println(findString[1])

	reg = regexp.MustCompile(`^(.*)（`)
	findString2 := reg.FindStringSubmatch(findString[1])
	// fmt.Println(findString2[1])
	name := findString2[1]

	reg = regexp.MustCompile(`（(.*)）`)
	findString2 = reg.FindStringSubmatch(findString[1])
	// fmt.Println(findString2[1])
	growp := findString2[1]

	// get room_id
	reg = regexp.MustCompile(`room_id=(\d+)`)
	findString = reg.FindStringSubmatch(bodyString)
	// fmt.Println(findString[1])
	key := findString[1]

	db, _ := gorm.Open(sqlite.Open("/data/sqlite.db"), &gorm.Config{})


	checkList := &model.CheckList{
		Group:    growp,
		Name:     name,
		Site:     "SHOWROOM",
		Key:      key,
		URL:      url,
		IsRecord: 1,
	}

	result := db.Create(&checkList)
	if result.Error != nil {
		panic(result.Error)
	}
}