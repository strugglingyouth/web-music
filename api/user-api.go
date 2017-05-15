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

// user
type Base struct {
	Status int    `json:"status"`
	From   string `json:"from"`
	Name   string `json:name`
}
type PopularUser struct {
	Base
	UserList []User `json:"userlist"`
}

type HotList struct {
	Base
	MusicListBase []Music_list `json:"my_music_list"`
}
type UserBase struct {
	Base
	UserBase      []User       `json:"user_info"`
	UploadBase    []Upload     `json:"upload_info"`
	MusicListBase []Music_list `json:"my_music_list"`
	//MusicListBase  []Upload     `json:"my_music_list"`
	//CollectionBase []Collection `json:"collection"`
	CollectionBase []Music_list `json:"collection"`
	FollowBase     []User       `json:"follow"`
	FansBase       []User       `json:"fans"`
}

type User struct {
	User_name      string `json:"user_name"`
	User_password  string `json:"user_password"`
	User_nick_name string `json:"user_nick_name"`
	User_birth     string `json:"user_birth"`
	User_sex       string `json:"user_sex"`
	User_intro     string `json:"user_intro"`
	User_open      string `json:"user_open"`
	User_list_open string `json:"user_list_open"`
	Fans_count     string `json:"fans_count"`
	Follow_count   string `json:"follow_count"`
	List_count     string `json:"list_count"`
	User_avatar    string `json:"user_avatar"`
}

type Upload struct {
	Id                    string `json:"upload_id"`
	Upload_user_name      string `json:"upload_user_name"`
	Upload_music_name     string `json:"upload_music_name"`
	Upload_open           string `json:"upload_open"`
	Upload_date           string `json:"upload_date"`
	Upload_music_file_url string `json:"upload_music_file_url"`
	From_self             string `json:"from_self"`
}

//歌单
type Music_list struct {
	Id string `json:"my_list_id"`
	//My_list_id        string `json:"my_list_id"`
	My_list_name      string `json:"my_list_name"`
	My_list_count     string `json:"my_list_count"`
	My_list_open      string `json:"my_list_open"`
	Create_date       string `json:"create_date"`
	My_list_user_name string `json:"my_list_user_name"`
	From_self         string `json:"from_self"`
	Collection        string `json:"collection"`
	//Music_info        []Music `json:"music"`
	Music_info []Upload `json:"music"`
}

// 歌单歌曲表
type Music struct {
	Music_id        string `json:"music_id"`
	Music_url       string `json:"music_url"`
	Music_name      string `json:"music_name"`
	Music_album     string `json:"music_album"`
	Music_mv        string `json:"music_mv"`
	Music_artist    string `json:"music_artist"`
	Music_artist_id string `json:"music_artist_id"`
	Id              string `json:"id"`
	My_list_id      string `json:"my_list_id"`
	Song_id         string `json:"music_id"`
}

type Collection struct {
	User_name     string       `json:"user_name"`
	My_list_id    string       `json:"my_list_id"`
	MusicListBase []Music_list `json:"collection_list"`
}

type Collect struct {
	Collect []Collection `json:"collect"`
}

type AuthUser struct {
	User_name     string `json:"user_name"`
	User_password string `json:"user_password"`
}

//type Follow struct {
//Followed_user_name string `json:"followed_user_name"`
//Follow_user_name   string `json:"follow_user_name"`
//}
//数据库操作出错
func DatabaseFailReturn(err error) []byte {
	Response := Base{}
	if err != nil {
		fmt.Println("insert failed!")
		Response.Status = http.StatusRequestTimeout
	} else {
		fmt.Println("insert success!")
		Response.Status = http.StatusOK
	}
	output, _ := json.Marshal(Response)
	return output
}

//接收数据有问题
func FormatErrorReturn(err error) []byte {

	Response := Base{}
	fmt.Println("Error: the data format is wrong!\n", err)
	Response.Status = http.StatusNotAcceptable //406
	output, _ := json.Marshal(Response)
	return output
}

// HTTP POST - 验证用户并获取用户数据 user api
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	base_url := "http://222.24.63.118:8080/upload/"
	table := "app1_user"
	authuser := AuthUser{}
	r.ParseForm()
	authuser.User_name = r.FormValue("user_name")
	authuser.User_password = r.FormValue("user_password")
	output, err := json.Marshal(authuser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Errorn\n:", err)
	}
	sql := "select * from " + table + " where  user_name='" + authuser.User_name + "' and user_password='" + authuser.User_password + "';"
	fmt.Println("sql:", sql)
	user_name := authuser.User_name

	//验证用户
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := UserBase{}
	fmt.Println("-----------------rows:", rows)
	ct := 0
	for rows.Next() {
		ct++
		user := User{}
		rows.Scan(&user.User_name, &user.User_password, &user.User_nick_name, &user.User_birth, &user.User_sex, &user.User_intro, &user.User_open, &user.User_list_open, &user.Fans_count, &user.Follow_count, &user.List_count, &user.User_avatar)

		//fmt.Println("-----------------authuser user_name:", user.User_name)
		if len(user.User_avatar) != 0 {
			user.User_avatar = base_url + user.User_avatar
		}
		Response.UserBase = append(Response.UserBase, user)
	}

	if ct == 0 {
		//fmt.Println("-------------------------->>>>>>>>>>>>>>>>user_name is null")
		Response := Base{}
		Response.From = "user"
		Response.Name = table
		Response.Status = http.StatusBadRequest
		output, _ := json.Marshal(Response)
		fmt.Fprintln(w, string(output))
		return
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// fans 查询
	sql = "select * from " + table + " where user_name=(select user_name from app1_follow where follow_user_name='" + user_name + "');"
	fmt.Println(sql)
	rows, err = database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		rows.Scan(&user.User_name, &user.User_password, &user.User_nick_name, &user.User_birth, &user.User_sex, &user.User_intro, &user.User_open, &user.User_list_open, &user.Fans_count, &user.Follow_count, &user.List_count, &user.User_avatar)
		if len(user.User_avatar) != 0 {
			user.User_avatar = base_url + user.User_avatar
		}
		Response.FansBase = append(Response.FansBase, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 获取 username 对应的 upload 数据
	upload_table := "app1_upload"
	sql = "select * from " + upload_table + " where upload_user_name='" + user_name + "';"
	fmt.Println(sql)
	rows, err = database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		upload := Upload{}
		rows.Scan(&upload.Id, &upload.Upload_user_name, &upload.Upload_music_name, &upload.Upload_open, &upload.Upload_date, &upload.Upload_music_file_url, &upload.From_self)
		if len(upload.Upload_music_file_url) != 0 {
			upload.Upload_music_file_url = base_url + upload.Upload_music_file_url
		}
		Response.UploadBase = append(Response.UploadBase, upload)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 获取 username 对应的歌单表
	music_list_table := "app1_my_list"

	sql = "select * from " + music_list_table + " where my_list_user_name='" + user_name + "';"
	fmt.Println(sql)
	rows, err = database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		music_list := Music_list{}
		rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
		Response.MusicListBase = append(Response.MusicListBase, music_list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//获取歌单中的歌曲
	//fmt.Println(len(Response.MusicListBase))
	for i := 0; i < len(Response.MusicListBase); i++ {
		id := Response.MusicListBase[i].Id

		sql = "select app1_upload.id,app1_upload.upload_user_name,app1_upload.upload_music_name,app1_upload.upload_open,app1_upload.upload_date,app1_upload.upload_music_file_url,app1_upload.from_myself from app1_upload inner join app1_my_list_to_music on app1_upload.id=app1_my_list_to_music.music_id  where   app1_my_list_to_music.my_list_id='" + id + "';"
		fmt.Println(sql)
		rows, err = database.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			upload := Upload{}
			rows.Scan(&upload.Id, &upload.Upload_user_name, &upload.Upload_music_name, &upload.Upload_open, &upload.Upload_date, &upload.Upload_music_file_url, &upload.From_self)
			//rows.Scan(&music.Music_id, &music.Music_url, &music.Music_name, &music.Music_album, &music.Music_mv, &music.Music_artist, &music.Music_artist_id, &music.Id, &music.My_list_id, &music.Song_id)
			Response.MusicListBase[i].Music_info = append(Response.MusicListBase[i].Music_info, upload)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

	}

	// 获取 username 对应的收藏表
	collection_table := "app1_collection"
	sql = "select * from " + collection_table + " where user_name='" + user_name + "';"
	fmt.Println(sql)
	rows, err = database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	coll := Collect{}
	for rows.Next() {
		collection := Collection{}
		rows.Scan(&collection.User_name, &collection.My_list_id)
		coll.Collect = append(coll.Collect, collection)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 收藏的歌单信息
	for i := 0; i < len(coll.Collect); i++ {
		id := coll.Collect[i].My_list_id
		fmt.Println("收藏的歌单信息--------------->", id)
		sql = "select app1_my_list.id,app1_my_list.my_list_name,app1_my_list.my_list_count,app1_my_list.my_list_open,app1_my_list.create_date,app1_my_list.my_list_user_name,app1_my_list.from_myself,app1_my_list.collection from app1_my_list inner join app1_collection  on app1_my_list.id=app1_collection.my_list_id  where my_list_id=" + id + ";"

		fmt.Println("collection:", sql)
		rows, err = database.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			music_list := Music_list{}
			rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
			//	rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
			fmt.Println("music_list:", music_list)
			Response.CollectionBase = append(Response.CollectionBase, music_list)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	//获取收藏歌单中的歌曲
	//fmt.Println(len(Response.MusicListBase))
	for i := 0; i < len(Response.CollectionBase); i++ {
		id := Response.CollectionBase[i].Id

		sql = "select app1_upload.id,app1_upload.upload_user_name,app1_upload.upload_music_name,app1_upload.upload_open,app1_upload.upload_date,app1_upload.upload_music_file_url,app1_upload.from_myself from app1_upload inner join app1_my_list_to_music on app1_upload.id=app1_my_list_to_music.music_id  where   app1_my_list_to_music.my_list_id='" + id + "';"
		fmt.Println(sql)
		rows, err = database.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			upload := Upload{}
			rows.Scan(&upload.Id, &upload.Upload_user_name, &upload.Upload_music_name, &upload.Upload_open, &upload.Upload_date, &upload.Upload_music_file_url, &upload.From_self)
			//rows.Scan(&music.Music_id, &music.Music_url, &music.Music_name, &music.Music_album, &music.Music_mv, &music.Music_artist, &music.Music_artist_id, &music.Id, &music.My_list_id, &music.Song_id)
			Response.CollectionBase[i].Music_info = append(Response.CollectionBase[i].Music_info, upload)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

	}

	// 获取 username 的关注者 follow
	//follow_table := "app1_follow"
	sql = "select * from " + table + " where user_name=(select follow_user_name from app1_follow where user_name='" + user_name + "');"

	fmt.Println("follow:", sql)
	rows, err = database.Query(sql)
	fmt.Println(rows)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		rows.Scan(&user.User_name, &user.User_password, &user.User_nick_name, &user.User_birth, &user.User_sex, &user.User_intro, &user.User_open, &user.User_list_open, &user.Fans_count, &user.Follow_count, &user.List_count, &user.User_avatar)
		if len(user.User_avatar) != 0 {
			user.User_avatar = base_url + user.User_avatar
		}
		Response.FollowBase = append(Response.FollowBase, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 添加 Status
	Response.From = "user"
	Response.Name = table
	Response.Status = http.StatusOK
	output, _ = json.Marshal(Response)
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

// HTTP GET - 直接 get 用户数据api
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	base_url := "http://222.24.63.118:8080/upload/"
	table := "app1_user"
	authuser := AuthUser{}
	//r.ParseForm()
	vars := mux.Vars(r)
	authuser.User_name = vars["user_name"]
	//authuser.User_password = r.FormValue("user_password")
	//output, err := json.Marshal(authuser)
	//fmt.Println(string(output))
	//if err != nil {
	//fmt.Println("Errorn\n:", err)
	//}
	sql := "select * from " + table + " where  user_name='" + authuser.User_name + "';"
	fmt.Println("sql:", sql)

	user_name := authuser.User_name

	//验证用户
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := UserBase{}
	fmt.Println("-----------------rows:", rows)
	ct := 0
	for rows.Next() {
		ct++
		user := User{}
		rows.Scan(&user.User_name, &user.User_password, &user.User_nick_name, &user.User_birth, &user.User_sex, &user.User_intro, &user.User_open, &user.User_list_open, &user.Fans_count, &user.Follow_count, &user.List_count, &user.User_avatar)

		//fmt.Println("-----------------authuser user_name:", user.User_name)
		if len(user.User_avatar) != 0 {
			user.User_avatar = base_url + user.User_avatar
		}
		Response.UserBase = append(Response.UserBase, user)
	}

	//if ct == 0 {
	////fmt.Println("-------------------------->>>>>>>>>>>>>>>>user_name is null")
	//Response := Base{}
	//Response.From = "user"
	//Response.Name = table
	//Response.Status = http.StatusBadRequest
	//output, _ := json.Marshal(Response)
	//fmt.Fprintln(w, string(output))
	//return
	//}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// fans 查询
	sql = "select * from " + table + " where user_name=(select user_name from app1_follow where follow_user_name='" + authuser.User_name + "');"
	fmt.Println(sql)
	rows, err = database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		rows.Scan(&user.User_name, &user.User_password, &user.User_nick_name, &user.User_birth, &user.User_sex, &user.User_intro, &user.User_open, &user.User_list_open, &user.Fans_count, &user.Follow_count, &user.List_count, &user.User_avatar)
		if len(user.User_avatar) != 0 {
			user.User_avatar = base_url + user.User_avatar
		}
		Response.FansBase = append(Response.FansBase, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 获取 username 对应的 upload 数据
	upload_table := "app1_upload"
	sql = "select * from " + upload_table + " where upload_user_name='" + user_name + "';"
	fmt.Println(sql)
	rows, err = database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		upload := Upload{}
		rows.Scan(&upload.Id, &upload.Upload_user_name, &upload.Upload_music_name, &upload.Upload_open, &upload.Upload_date, &upload.Upload_music_file_url, &upload.From_self)
		if len(upload.Upload_music_file_url) != 0 {
			upload.Upload_music_file_url = base_url + upload.Upload_music_file_url
		}
		Response.UploadBase = append(Response.UploadBase, upload)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 获取 username 对应的歌单表
	music_list_table := "app1_my_list"

	sql = "select * from " + music_list_table + " where my_list_user_name='" + user_name + "';"
	fmt.Println(sql)
	rows, err = database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		music_list := Music_list{}
		rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
		Response.MusicListBase = append(Response.MusicListBase, music_list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//获取歌单中的歌曲
	//fmt.Println(len(Response.MusicListBase))
	for i := 0; i < len(Response.MusicListBase); i++ {
		id := Response.MusicListBase[i].Id

		sql = "select app1_upload.id,app1_upload.upload_user_name,app1_upload.upload_music_name,app1_upload.upload_open,app1_upload.upload_date,app1_upload.upload_music_file_url,app1_upload.from_myself from app1_upload inner join app1_my_list_to_music on app1_upload.id=app1_my_list_to_music.music_id  where   app1_my_list_to_music.my_list_id='" + id + "';"
		fmt.Println(sql)
		rows, err = database.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			upload := Upload{}
			rows.Scan(&upload.Id, &upload.Upload_user_name, &upload.Upload_music_name, &upload.Upload_open, &upload.Upload_date, &upload.Upload_music_file_url, &upload.From_self)
			//rows.Scan(&music.Music_id, &music.Music_url, &music.Music_name, &music.Music_album, &music.Music_mv, &music.Music_artist, &music.Music_artist_id, &music.Id, &music.My_list_id, &music.Song_id)
			Response.MusicListBase[i].Music_info = append(Response.MusicListBase[i].Music_info, upload)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

	}

	// 获取 username 对应的收藏表
	collection_table := "app1_collection"
	sql = "select * from " + collection_table + " where user_name='" + user_name + "';"
	fmt.Println(sql)
	rows, err = database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	coll := Collect{}
	for rows.Next() {
		collection := Collection{}
		rows.Scan(&collection.User_name, &collection.My_list_id)
		coll.Collect = append(coll.Collect, collection)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 收藏的歌单信息
	for i := 0; i < len(coll.Collect); i++ {
		id := coll.Collect[i].My_list_id
		fmt.Println("收藏的歌单信息--------------->", id)
		sql = "select app1_my_list.id,app1_my_list.my_list_name,app1_my_list.my_list_count,app1_my_list.my_list_open,app1_my_list.create_date,app1_my_list.my_list_user_name,app1_my_list.from_myself,app1_my_list.collection from app1_my_list inner join app1_collection  on app1_my_list.id=app1_collection.my_list_id  where my_list_id=" + id + ";"

		fmt.Println("collection:", sql)
		rows, err = database.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			music_list := Music_list{}
			rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
			//	rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
			fmt.Println("music_list:", music_list)
			Response.CollectionBase = append(Response.CollectionBase, music_list)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	//获取收藏歌单中的歌曲
	//fmt.Println(len(Response.MusicListBase))
	for i := 0; i < len(Response.CollectionBase); i++ {
		id := Response.CollectionBase[i].Id
		sql = "select app1_upload.id,app1_upload.upload_user_name,app1_upload.upload_music_name,app1_upload.upload_open,app1_upload.upload_date,app1_upload.upload_music_file_url,app1_upload.from_myself from app1_upload inner join app1_my_list_to_music on app1_upload.id=app1_my_list_to_music.music_id  where   app1_my_list_to_music.my_list_id='" + id + "';"
		fmt.Println(sql)
		rows, err = database.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			upload := Upload{}
			rows.Scan(&upload.Id, &upload.Upload_user_name, &upload.Upload_music_name, &upload.Upload_open, &upload.Upload_date, &upload.Upload_music_file_url, &upload.From_self)
			//rows.Scan(&music.Music_id, &music.Music_url, &music.Music_name, &music.Music_album, &music.Music_mv, &music.Music_artist, &music.Music_artist_id, &music.Id, &music.My_list_id, &music.Song_id)
			Response.CollectionBase[i].Music_info = append(Response.CollectionBase[i].Music_info, upload)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

	}

	// 获取 username 的关注者 follow
	//follow_table := "app1_follow"
	sql = "select * from " + table + " where user_name=(select follow_user_name from app1_follow where user_name='" + user_name + "');"

	fmt.Println("follow:", sql)
	rows, err = database.Query(sql)
	fmt.Println(rows)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		rows.Scan(&user.User_name, &user.User_password, &user.User_nick_name, &user.User_birth, &user.User_sex, &user.User_intro, &user.User_open, &user.User_list_open, &user.Fans_count, &user.Follow_count, &user.List_count, &user.User_avatar)
		if len(user.User_avatar) != 0 {
			user.User_avatar = base_url + user.User_avatar
		}
		Response.FollowBase = append(Response.FollowBase, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 添加 Status
	Response.From = "user"
	Response.Name = table
	Response.Status = http.StatusOK
	output, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}
func GetHotListHandler(w http.ResponseWriter, r *http.Request) {
	// 收藏的歌单信息
	Response := HotList{}
	sql := "select * from app1_my_list where my_list_open=1 order by collection desc limit 10;"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	output := DatabaseFailReturn(err)

	fmt.Println("get hot list:", sql)
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		music_list := Music_list{}
		rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
		fmt.Println("music_list:", music_list)
		Response.MusicListBase = append(Response.MusicListBase, music_list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(Response.MusicListBase); i++ {
		id := Response.MusicListBase[i].Id
		sql = "select app1_upload.id,app1_upload.upload_user_name,app1_upload.upload_music_name,app1_upload.upload_open,app1_upload.upload_date,app1_upload.upload_music_file_url,app1_upload.from_myself from app1_upload inner join app1_my_list_to_music on app1_upload.id=app1_my_list_to_music.music_id  where   app1_my_list_to_music.my_list_id='" + id + "';"
		fmt.Println(sql)
		rows, err = database.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			upload := Upload{}
			rows.Scan(&upload.Id, &upload.Upload_user_name, &upload.Upload_music_name, &upload.Upload_open, &upload.Upload_date, &upload.Upload_music_file_url, &upload.From_self)
			Response.MusicListBase[i].Music_info = append(Response.MusicListBase[i].Music_info, upload)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	// 添加 Status
	Response.From = "user"
	Response.Name = "app1_my_list"
	Response.Status = http.StatusOK
	output, _ = json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}
func GetNewListHandler(w http.ResponseWriter, r *http.Request) {
	// 收藏的歌单信息
	Response := HotList{}
	sql := "select * from app1_my_list where my_list_open=1 order by create_date desc limit 10;"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	output := DatabaseFailReturn(err)

	fmt.Println("get hot list:", sql)
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		music_list := Music_list{}
		rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
		fmt.Println("music_list:", music_list)
		Response.MusicListBase = append(Response.MusicListBase, music_list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(Response.MusicListBase); i++ {
		id := Response.MusicListBase[i].Id
		sql = "select app1_upload.id,app1_upload.upload_user_name,app1_upload.upload_music_name,app1_upload.upload_open,app1_upload.upload_date,app1_upload.upload_music_file_url,app1_upload.from_myself from app1_upload inner join app1_my_list_to_music on app1_upload.id=app1_my_list_to_music.music_id  where   app1_my_list_to_music.my_list_id='" + id + "';"
		fmt.Println(sql)
		rows, err = database.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			upload := Upload{}
			rows.Scan(&upload.Id, &upload.Upload_user_name, &upload.Upload_music_name, &upload.Upload_open, &upload.Upload_date, &upload.Upload_music_file_url, &upload.From_self)
			Response.MusicListBase[i].Music_info = append(Response.MusicListBase[i].Music_info, upload)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	// 添加 Status
	Response.From = "user"
	Response.Name = "app1_my_list"
	Response.Status = http.StatusOK
	output, _ = json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}
func GetPopularUserHandler(w http.ResponseWriter, r *http.Request) {

	base_url := "http://222.24.63.118:8080/upload/"
	sql := "select * from app1_user order by fans_count desc limit 10;"
	fmt.Println("sql:", sql)

	//验证用户
	rows, err := database.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := PopularUser{}
	fmt.Println("-----------------rows:", rows)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.User_name, &user.User_password, &user.User_nick_name, &user.User_birth, &user.User_sex, &user.User_intro, &user.User_open, &user.User_list_open, &user.Fans_count, &user.Follow_count, &user.List_count, &user.User_avatar)

		if len(user.User_avatar) != 0 {
			user.User_avatar = base_url + user.User_avatar
		}
		Response.UserList = append(Response.UserList, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// 添加 Status
	Response.From = "user"
	Response.Name = "app1_user"
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

	//验证用户并获取数据  POST
	//r.HandleFunc("/user/{table}/{user_name}", GetUserHandler).Methods("POST")
	r.HandleFunc("/user", GetUserHandler).Methods("POST")
	// get 用户数据
	r.HandleFunc("/userinfo/{user_name}", GetUserInfoHandler).Methods("GET")
	// get 收藏数前 10 的歌单表
	r.HandleFunc("/hotlist", GetHotListHandler).Methods("GET")

	// get 按时间获取最新歌单表
	r.HandleFunc("/newlist", GetNewListHandler).Methods("GET")

	// 获取关注数最多的用户
	r.HandleFunc("/popularuser", GetPopularUserHandler).Methods("GET")
	server := &http.Server{
		Addr: ":8881",
		//Addr:    ":7771",
		Handler: r,
	}
	log.Printf("Listening at %s ...", server.Addr)
	server.ListenAndServe()
}
