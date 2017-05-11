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

// user
type Base struct {
	Status int `json:"status"`
	//From   string `json:"from"`
	//Name   string `json:name`
}
type UserBase struct {
	Base
	CreateUser CreateUser `json:"createuser"`
}

type CreateUser struct {
	User_name     string `json:"user_name"`
	User_password string `json:"user_password"`
}

// HTTP POST - /api/createuser
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	create_user := CreateUser{}
	//解码到note, 使用此方式时接受到的类型为 string
	//err := json.NewDecoder(r.Body).Decode(&new_note)
	r.ParseForm()
	fmt.Println(r)
	create_user.User_name = r.FormValue("user_name")
	create_user.User_password = r.FormValue("user_password")
	output, err := json.Marshal(create_user)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Errorn\n:", err)
	}
	sql := "insert into app1_user set user_name='" + create_user.User_name + "',user_password=" + create_user.User_password + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	if err != nil {
		fmt.Println("Error:\n", err)
	} else {
		fmt.Println("exec insert sql success!")
	}

	Response := Base{}
	if err != nil {
		Response.Status = http.StatusBadRequest

	} else {
		Response.Status = http.StatusOK
	}
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

///HTTP Put - /modifyuser/{id}
func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	r.ParseForm()
	user.User_name = r.FormValue("user_name")
	user.User_password = r.FormValue("user_password")
	user.User_nick_name = r.FormValue("user_nick_name")
	user.User_birth = r.FormValue("user_birth")
	user.User_sex = r.FormValue("user_sex")
	user.User_intro = r.FormValue("user_intro")
	user.User_open = r.FormValue("user_open")
	user.User_list_open = r.FormValue("user_list_open")
	user.Fans_count = r.FormValue("fans_count")
	user.Follow_count = r.FormValue("follow_count")
	user.List_count = r.FormValue("list_count")
	user.User_avatar = r.FormValue("user_avatar")

	output, err := json.Marshal(user)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Errorn\n:", err)
	}
	sql := "update app1_user set user_name='" + user.User_name + "',user_password='" + user.User_password + "',user_nick_name='" + user.User_nick_name + "',user_birth='" + user.User_birth + "',user_sex='" + user.User_sex + "',user_intro='" + user.User_intro + "',user_open='" + user.User_open + "',user_list_open='" + user.User_list_open + "',fans_count='" + user.Fans_count + "',follow_count='" + user.Follow_count + "',list_count='" + user.List_count + "',user_avatar='" + user.User_avatar + "' where user_name=" + user.User_name + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	if err != nil {
		fmt.Println("Error:\n", err)
	} else {
		fmt.Println("exec insert sql success!")
	}
	Response := Base{}
	if err != nil {
		Response.Status = http.StatusBadRequest

	} else {
		Response.Status = http.StatusOK
	}
	output, _ = json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}

// HTTP POST - /createupload  上传歌曲
func CreateUploadHandler(w http.ResponseWriter, r *http.Request) {
	create_user := CreateUser{}
	r.ParseForm()
	fmt.Println(r)
	create_user.User_name = r.FormValue("user_name")
	create_user.User_password = r.FormValue("user_password")
	output, err := json.Marshal(create_user)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Errorn\n:", err)
	}
	sql := "insert into app1_user set user_name='" + create_user.User_name + "',user_password=" + create_user.User_password + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	if err != nil {
		fmt.Println("Error:\n", err)
	} else {
		fmt.Println("exec insert sql success!")
	}
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/demo")
	if err != nil {
		log.Fatal("Error on initializing database connection: %s", err.Error())
	}
	defer db.Close()
	database = db
	r := mux.NewRouter().StrictSlash(false)

	//创建用户
	r.HandleFunc("/createuser", CreateUserHandler).Methods("POST")
	//修改用户信息
	r.HandleFunc("/modifyuser/{id}", PutUserHandler).Methods("PUT")

	// 添加上传歌曲
	r.HandleFunc("/createupload", CreateUploadHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8880",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
