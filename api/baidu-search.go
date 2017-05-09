package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var database *sql.DB

func httpPostMusic(query string) string {
	post_str := "s=" + query + "&limit=20&type=1&offset=0"
	resp, err := http.Post("http://music.163.com/api/search/get/",
		"application/x-www-form-urlencoded",
		strings.NewReader(post_str))
	//fmt.Println(post_str)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	content := string(body)
	//fmt.Println(content)
	return content
}

// HTTP GET - /api/notes   Wangyiyun
func GetWangyiyunMusicHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//from := vars["from"]
	query := vars["query"]
	//fmt.Println(query)

	output := httpPostMusic(query)

	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")

	fmt.Fprintln(w, string(output))
}

// search singer 'type=100'
func httpPostSinger(query string) string {
	post_str := "s=" + query + "&limit=20&type=100&offset=0"
	resp, err := http.Post("http://music.163.com/api/search/get/",
		"application/x-www-form-urlencoded",
		strings.NewReader(post_str))
	//fmt.Println(post_str)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	content := string(body)
	//fmt.Println(content)
	return content
}

// HTTP GET - /api/notes   Wangyiyun  歌手搜索
func GetWangyiyunSingerHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//from := vars["from"]
	query := vars["query"]
	//fmt.Println(query)

	output := httpPostSinger(query)

	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")

	fmt.Fprintln(w, string(output))
}

// search singer 'type=10'
func httpPostBaidu(query string) string {

	post_str := "from=webapp_music&version=5.6.5.0&method=baidu.ting.search.catalogSug&format=json&query=" + query
	resp, err := http.Post("http://tingapi.ting.baidu.com/v1/restserver/ting",
		"application/x-www-form-urlencoded",
		strings.NewReader(post_str))
	fmt.Println(post_str)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	content := string(body)
	//fmt.Println(content)
	return content
}

// HTTP GET baidu 百度音乐搜索
func GetBaiduHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//from := vars["from"]
	query := vars["query"]
	//fmt.Println(query)

	output := httpPostBaidu(query)

	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")

	fmt.Fprintln(w, string(output))
}
func main() {
	r := mux.NewRouter().StrictSlash(false)

	//百度音乐搜索
	r.HandleFunc("/api/baidu/{query}", GetBaiduHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8882",
		Handler: r,
	}
	log.Printf("Listening at %s ...", server.Addr)
	server.ListenAndServe()
}
