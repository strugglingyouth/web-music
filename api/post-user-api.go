package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var database *sql.DB

//设置 upload 初始 id
var id = 200

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
}
type UserBase struct {
	Base
	CreateUser CreateUser `json:"createuser"`
}

type UploadBase struct {
	Base
	Upload Upload `json:"upload"`
}

type CreateUser struct {
	User_name     string `json:"user_name"`
	User_password string `json:"user_password"`
}

// 上传歌曲表
type Upload struct {
	Id                    string `json:"upload_id"`
	Upload_user_name      string `json:"upload_user_name"`
	Upload_music_name     string `json:"upload_music_name"`
	Upload_date           string `json:"upload_date"`
	Upload_music_file_url string `json:"upload_music_file_url"`
	From_self             string `json:"from_self"`
	Upload_open           string `json:"upload_open"`
}

type DeleteUpload struct {
	Id               string `json:"upload_id"`
	Upload_user_name string `json:"upload_user_name"`
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
		output = FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}
	sql := "insert into app1_user set user_name='" + create_user.User_name + "',user_password=" + create_user.User_password + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)

	output = DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))

}

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
	//fmt.Println(string(output))

	// 判断是否有数据为空
	if err != nil || len(user.User_name) == 0 || len(user.User_password) == 0 {
		output = FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}
	sql := "update app1_user set user_name='" + user.User_name + "',user_password='" + user.User_password + "',user_nick_name='" + user.User_nick_name + "',user_birth='" + user.User_birth + "',user_sex='" + user.User_sex + "',user_intro='" + user.User_intro + "',user_open='" + user.User_open + "',user_list_open='" + user.User_list_open + "',fans_count='" + user.Fans_count + "',follow_count='" + user.Follow_count + "',list_count='" + user.List_count + "',user_avatar='" + user.User_avatar + "' where user_name=" + user.User_name + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)

	output = DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))

}

// HTTP POST - /upload  上传歌曲
func PostUploadHandler(w http.ResponseWriter, r *http.Request) {
	table := "app1_upload"
	base_url := "http://222.24.63.118:8080/upload/"
	upload := Upload{}
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	upload.Upload_date = tm.Format("2006-01-02 03:04:05")
	upload.Id = strconv.Itoa(id)
	r.ParseMultipartForm(32 << 20)
	upload.Upload_user_name = r.FormValue("upload_user_name")
	upload.Upload_music_name = r.FormValue("upload_music_name")

	// download  file
	file, handler, err := r.FormFile("upload_music_file_url")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("/usr/demo/upload/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	upload.Upload_music_file_url = base_url + handler.Filename
	upload.Upload_open = "0"
	upload.From_self = "1"

	output, err := json.Marshal(upload)
	//fmt.Println(string(output))
	if err != nil || len(upload.Upload_user_name) == 0 {
		output = FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}
	sql := "insert into " + table + " set id=" + upload.Id + ",upload_user_name='" + upload.Upload_user_name + "',upload_music_name='" + upload.Upload_music_name + "',upload_date='" + upload.Upload_date + "',upload_music_file_url='" + upload.Upload_music_file_url + "';"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)

	Response := UploadBase{}
	if err != nil {
		fmt.Println("insert failed!")
		Response.Base.Status = http.StatusRequestTimeout
	} else {
		fmt.Println("insert success!")
		Response.Base.Status = http.StatusOK
		//Response.Upload = append(Response.Upload, upload)
		Response.Upload = upload
	}
	output, _ = json.Marshal(Response)

	//output = DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))
	id++
}

// HTTP Delete   deleet upload  删除上传歌单
func DeleteUploadHandler(w http.ResponseWriter, r *http.Request) {
	table := "app1_upload"
	r.ParseForm()
	fmt.Println(r)
	upload := Upload{}
	upload.Id = r.FormValue("id")
	upload.Upload_user_name = r.FormValue("upload_user_name")
	upload.Upload_music_name = r.FormValue("upload_music_name")
	upload.Upload_open = r.FormValue("upload_open")
	upload.Upload_date = r.FormValue("upload_date")
	upload.Upload_music_file_url = r.FormValue("upload_music_file_url")
	upload.From_self = r.FormValue("from_self")

	output, err := json.Marshal(upload)
	fmt.Println(string(output))
	if err != nil || len(upload.Upload_user_name) == 0 || len(upload.Id) == 0 {
		output = FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}
	sql := "delete from " + table + " where id=" + upload.Id + " and upload_user_name='" + upload.Upload_user_name + "';"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	output = DatabaseFailReturn(err)
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

	//创建用户
	r.HandleFunc("/createuser", CreateUserHandler).Methods("POST")
	//修改用户信息
	r.HandleFunc("/modifyuser/{id}", PutUserHandler).Methods("PUT")

	// 添加上传歌曲
	r.HandleFunc("/upload", PostUploadHandler).Methods("POST")

	//删除上传歌曲
	r.HandleFunc("/delete/upload", DeleteUploadHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8883",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
