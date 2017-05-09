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

// 歌手信息
type Base struct {
	Status int      `json:"status"`
	From   string   `json:"from"`
	Name   string   `json:name`
	Base   []Singer `json:"data"`
}

type Singer struct {
	Singer_id   string `json:"singer_id"`
	Singer_name string `json:"singer_name"`
	Singer_pic  string `json:"singer_pic"`
}

// 歌曲评论表
type CommentBase struct {
	Status      int       `json:"status"`
	From        string    `json:"from"`
	Name        string    `json:name`
	CommentBase []Comment `json:"data"`
}

type Comment struct {
	Id                 string `json:"id"`
	Music_id           string `json:"music_id"`
	Comment_content    string `json:"comment_content"`
	Comment_user_id    string `json:"comment_user_id"`
	Comment_username   string `json:"comment_username"`
	Comment_like_count string `json:"comment_like_count"`
	Comment_id         string `json:"comment_id"`
}

// 歌曲详细信息
type MusicBase struct {
	Status    int     `json:"status"`
	From      string  `json:"from"`
	Name      string  `json:name`
	MusicBase []Music `json:"data"`
}

type Music struct {
	Music_id        string `json:"music_id"`
	Music_url       string `json:"music_url"`
	Music_name      string `json:"music_name"`
	Music_album     string `json:"music_album"`
	Music_mv        string `json:"music_mv"`
	Music_artist    string `json:"music_artist"`
	Music_artist_id string `json:"music_artist_id"`
}

// HTTP GET - /api/wangyiyun/app1_singer     歌手信息
func GetWangyiyunsingerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["from"]
	table := vars["table"]
	fmt.Println(table)
	//fmt.Println(r)
	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	sql := "select * from " + table + ";"
	fmt.Println(sql)
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := Base{}
	for rows.Next() {
		singer := Singer{}
		rows.Scan(&singer.Singer_id, &singer.Singer_name, &singer.Singer_pic)
		Response.Base = append(Response.Base, singer)
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

// HTTP GET - /api/music/app1_musiccomment     歌曲信息
func GetMusicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["from"]
	table := "app1_music_items"
	music_artist_id := vars["id"]
	fmt.Println(music_artist_id)
	//fmt.Println(r)
	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	sql := "select * from " + table + " where music_artist_id=" + music_artist_id + ";"
	fmt.Println(sql)
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := MusicBase{}
	for rows.Next() {
		music := Music{}
		rows.Scan(&music.Music_id, &music.Music_url, &music.Music_name, &music.Music_album, &music.Music_mv, &music.Music_artist, &music.Music_artist_id)
		Response.MusicBase = append(Response.MusicBase, music)
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

// HTTP GET - /api/comment/app1_musiccomment     歌曲评论信息
func GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["from"]
	table := "app1_musiccomment"
	music_id := vars["id"]
	fmt.Println(music_id)
	//fmt.Println(r)
	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	sql := "select * from " + table + " where music_id='" + music_id + "';"
	fmt.Println(sql)
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := CommentBase{}
	for rows.Next() {
		comment := Comment{}
		rows.Scan(&comment.Id, &comment.Music_id, &comment.Comment_content, &comment.Comment_user_id, &comment.Comment_username, &comment.Comment_like_count, &comment.Comment_id)
		Response.CommentBase = append(Response.CommentBase, comment)
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
func main() {
	db, err := sql.Open("mysql", "root:123456@/demo")
	if err != nil {
		log.Fatal("Error on initializing database connection: %s", err.Error())
	}
	defer db.Close()
	database = db
	r := mux.NewRouter().StrictSlash(false)

	//酷狗
	//r.HandleFunc("/api/kugou/{table}", GetKugouHandler).Methods("GET")

	//百度音乐
	//r.HandleFunc("/api/baidu/{table}", GetBaiduHandler).Methods("GET")

	//歌手信息
	r.HandleFunc("/api/singer/{table}", GetWangyiyunsingerHandler).Methods("GET")
	//歌曲信息
	r.HandleFunc("/api/music/{id}", GetMusicHandler).Methods("GET")
	//歌曲评论信息
	r.HandleFunc("/api/comment/{id}", GetCommentHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8886",
		Handler: r,
	}
	log.Printf("Listening at %s ...", server.Addr)
	server.ListenAndServe()
}
