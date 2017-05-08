package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var database *sql.DB

// kugou
type Base struct {
	Status int     `json:"status"`
	From   string  `json:"from"`
	Name   string  `json:name`
	Base   []Kugou `json:"data"`
}

type Kugou struct {
	Music_id   string `json:"music_id"`
	Music_url  string `json:"music_url"`
	Music_name string `json:"music_name"`
}

// 百度音乐  music
type BaiduBase struct {
	Status    int     `json:"status"`
	From      string  `json:"from"`
	Name      string  `json:name`
	BaiduBase []Baidu `json:"data"`
}

type Baidu struct {
	Music_id         string `json:"music_id"`
	Music_url        string `json:"music_url"`
	Music_name       string `json:"music_name"`
	Music_artist     string `json:"music_artist"`
	Music_artist_url string `json:"music_artist_url"`
}

// 网易云 music
type WangyiyunBase struct {
	Status        int         `json:"status"`
	From          string      `json:"from"`
	Name          string      `json:name`
	WangyiyunBase []Wangyiyun `json:"data"`
}

type Wangyiyun struct {
	Music_id        string `json:"music_id"`
	Music_url       string `json:"music_url"`
	Music_name      string `json:"music_name"`
	Music_album     string `json:"music_album"`
	Music_artist    string `json:"music_artist"`
	Music_artist_id string `json:"music_artist_id"`
}

// HTTP GET - /api/kugou/kugou_net_hot    酷狗
func GetKugouHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["from"]
	table := vars["table"]
	fmt.Println(table)
	//fmt.Println(r)
	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	sql := "select * from " + table + " limit 10;"
	fmt.Println(sql)
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := Base{}
	for rows.Next() {
		kugou := Kugou{}
		rows.Scan(&kugou.Music_id, &kugou.Music_url, &kugou.Music_name)
		Response.Base = append(Response.Base, kugou)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// 添加 Status
	Response.From = "kugou"
	Response.Name = table
	Response.Status = http.StatusOK
	output, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}

// HTTP GET - /api/baidu/     百度音乐
func GetBaiduHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["from"]
	table := vars["table"]
	fmt.Println(table)
	//fmt.Println(r)
	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	sql := "select * from " + table + " limit 10;"
	fmt.Println(sql)
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := BaiduBase{}
	for rows.Next() {
		baidu := Baidu{}
		rows.Scan(&baidu.Music_id, &baidu.Music_url, &baidu.Music_name, &baidu.Music_artist, &baidu.Music_artist_url)
		Response.BaiduBase = append(Response.BaiduBase, baidu)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// 添加 Status
	Response.From = "baidu"
	Response.Name = table
	Response.Status = http.StatusOK
	output, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}

// HTTP GET - /api/wangyiyun/     网易云
func GetWangyiyunHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["from"]
	table := vars["table"]
	fmt.Println(table)
	//fmt.Println(r)
	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	sql := "select * from " + table + " limit 10;"
	fmt.Println(sql)
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := WangyiyunBase{}
	for rows.Next() {
		wangyiyun := Wangyiyun{}
		rows.Scan(&wangyiyun.Music_id, &wangyiyun.Music_url, &wangyiyun.Music_name, &wangyiyun.Music_album, &wangyiyun.Music_artist, &wangyiyun.Music_artist_id)
		Response.WangyiyunBase = append(Response.WangyiyunBase, wangyiyun)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// 添加 Status
	Response.From = "wangyiyun"
	Response.Name = table
	Response.Status = http.StatusOK
	output, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}

// HTTP Delete - /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	sql := "delete  from notes where id=" + id + ";"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	if err != nil {
		fmt.Println("Error:\n", err)
	} else {
		fmt.Println("exec delete sql success!")
	}
}
func main() {
	db, err := sql.Open("mysql", "root:123456@/api")
	if err != nil {
		log.Fatal("Error on initializing database connection: %s", err.Error())
	}
	defer db.Close()
	database = db
	r := mux.NewRouter().StrictSlash(false)

	//酷狗
	r.HandleFunc("/api/kugou/{table}", GetKugouHandler).Methods("GET")

	//百度音乐
	r.HandleFunc("/api/baidu/{table}", GetBaiduHandler).Methods("GET")

	//网易云
	r.HandleFunc("/api/wangyiyun/{table}", GetWangyiyunHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	log.Printf("Listening at %s ...", server.Addr)
	server.ListenAndServe()
}
