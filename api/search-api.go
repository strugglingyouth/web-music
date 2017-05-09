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
func httpPostAlbum(query string) string {
	post_str := "s=" + query + "&limit=20&type=10&offset=0"
	resp, err := http.Post("http://music.163.com/api/search/get/",
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

// HTTP GET    Wangyiyun  专辑搜索  'type=10'
func GetWangyiyunAlbumHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//from := vars["from"]
	query := vars["query"]
	//fmt.Println(query)

	output := httpPostAlbum(query)

	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")

	fmt.Fprintln(w, string(output))
}

// search 百度音乐搜索
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
	query := vars["query"]
	output := httpPostBaidu(query)

	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	fmt.Fprintln(w, string(output))
}

// search 酷狗歌曲搜索
func httpPostKugou(query string) string {

	post_str := "page=1&pagesize=20&format=json&keyword=" + query
	resp, err := http.Get("http://mobilecdn.kugou.com/api/v3/search/song?" + post_str)
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
	return content
}

// HTTP GET 酷狗歌曲搜索
func GetKugouMusicHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	query := vars["query"]
	output := httpPostKugou(query)

	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	fmt.Fprintln(w, string(output))
}

// search 酷狗专辑搜索
func httpPostKugouAlbum(query string) string {
	post_str := "page=1&pagesize=20&keyword=" + query
	resp, err := http.Get("http://mobilecdn.kugou.com/api/v3/search/album?" + post_str)
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
	return content
}

// HTTP GET - /api/   酷狗专辑搜索
func GetKugouAlbumHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	query := vars["query"]

	output := httpPostKugouAlbum(query)

	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")

	fmt.Fprintln(w, string(output))
}

func main() {
	r := mux.NewRouter().StrictSlash(false)

	//网易云音乐搜索
	r.HandleFunc("/api/wangyiyunmusic/{query}", GetWangyiyunMusicHandler).Methods("GET")
	//网易云歌手搜索
	r.HandleFunc("/api/wangyiyunsinger/{query}", GetWangyiyunSingerHandler).Methods("GET")
	//网易云专辑搜索
	r.HandleFunc("/api/wangyiyunalbum/{query}", GetWangyiyunAlbumHandler).Methods("GET")

	//百度音乐搜索
	r.HandleFunc("/api/baidu/{query}", GetBaiduHandler).Methods("GET")

	//酷狗音乐搜索
	r.HandleFunc("/api/kugoumusic/{query}", GetKugouMusicHandler).Methods("GET")
	//酷狗歌手搜索
	//r.HandleFunc("/api/kugousinger/{query}", GetKugouSingerHandler).Methods("GET")
	//酷狗专辑搜索
	r.HandleFunc("/api/kugoualbum/{query}", GetKugouAlbumHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8884",
		Handler: r,
	}
	log.Printf("Listening at %s ...", server.Addr)
	server.ListenAndServe()
}
