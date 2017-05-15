package main

import (
	"crypto/md5"
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
var id = 300

var timestamp = time.Now().Unix()
var tm = time.Unix(timestamp, 0)

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

type Follow struct {
	User_name        string `json:"user_name"`
	Follow_user_name string `json:"follow_user_name"`
}

//歌单
type Music_list struct {
	Id                string `json:"my_list_id"`
	My_list_name      string `json:"my_list_name"`
	My_list_count     string `json:"my_list_count"`
	My_list_open      string `json:"my_list_open"`
	Create_date       string `json:"create_date"`
	My_list_user_name string `json:"my_list_user_name"`
	From_self         string `json:"from_self"`
	Collection        string `json:"collection"`
}

type Collection struct {
	User_name  string `json:"user_name"`
	My_list_id string `json:"my_list_id"`
}

type ListToMusic struct {
	My_list_id string `json:"my_list_id"`
	Music_id   string `json:"music_id"`
}

//对字符串进行MD5哈希
func Md5hash(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
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

	// 判断传入参数是否正确
	if err != nil || len(create_user.User_name) == 0 || len(create_user.User_password) == 0 {
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
	base_url := "avatar/"
	user := User{}
	//r.ParseForm()
	r.ParseMultipartForm(32 << 20)
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
	//user.User_avatar = r.FormValue("user_avatar")

	//output, err := json.Marshal(user)
	//fmt.Println(string(output))

	// 判断是否有数据为空
	if len(user.User_name) == 0 || len(user.User_password) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	// download    avator
	file, handler, err := r.FormFile("user_avatar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	//handler.Filename = Md5hash(handler.Filename) + ".png"
	handler.Filename = user.User_name + ".png"
	fmt.Println("filename", handler.Filename)
	f, err := os.OpenFile("/usr/demo/upload/avatar/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		//return
	}
	defer f.Close()
	io.Copy(f, file)

	user.User_avatar = base_url + handler.Filename

	sql := "update app1_user set user_name='" + user.User_name + "',user_password='" + user.User_password + "',user_nick_name='" + user.User_nick_name + "',user_birth='" + user.User_birth + "',user_sex='" + user.User_sex + "',user_intro='" + user.User_intro + "',user_open='" + user.User_open + "',user_list_open='" + user.User_list_open + "',fans_count='" + user.Fans_count + "',follow_count='" + user.Follow_count + "',list_count='" + user.List_count + "',user_avatar='" + user.User_avatar + "' where user_name='" + user.User_name + "' and user_password='" + user.User_password + "';"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)

	output := DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))

}

// HTTP POST - /upload  上传歌曲
func PostUploadHandler(w http.ResponseWriter, r *http.Request) {
	table := "app1_upload"
	base_url := "upload/"
	upload := Upload{}
	upload.Upload_date = tm.Format("2006-01-02 03:04:05")
	//upload.Id = strconv.Itoa(id)
	r.ParseMultipartForm(32 << 20)
	upload.Upload_user_name = r.FormValue("upload_user_name")
	upload.Upload_music_name = r.FormValue("upload_music_name")

	// download  music  file
	file, handler, err := r.FormFile("upload_music_file_url")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//fmt.Fprintf(w, "%v", handler.Header)
	handler.Filename = Md5hash(handler.Filename) + ".mp3"
	fmt.Println("filename", handler.Filename)
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
	sql := "insert into " + table + " set upload_user_name='" + upload.Upload_user_name + "',upload_music_name='" + upload.Upload_music_name + "',upload_date='" + upload.Upload_date + "',upload_music_file_url='" + upload.Upload_music_file_url + "';"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)

	Response := UploadBase{}
	if err != nil {
		fmt.Println("insert failed!")
		Response.Base.Status = http.StatusRequestTimeout
	} else {
		fmt.Println("insert success!")
		Response.Base.Status = http.StatusOK
		Response.Upload = upload
	}
	output, _ = json.Marshal(Response)
	fmt.Fprintln(w, string(output))
	id++
}

// HTTP Delete   deleet upload  删除上传歌曲
func DeleteUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 通过 id  删除歌单
	table := "app1_upload"
	r.ParseForm()
	fmt.Println(r)
	upload := Upload{}
	upload.Id = r.FormValue("id")
	//upload.Upload_user_name = r.FormValue("upload_user_name")
	//upload.Upload_music_name = r.FormValue("upload_music_name")
	//upload.Upload_open = r.FormValue("upload_open")
	//upload.Upload_date = r.FormValue("upload_date")
	//upload.Upload_music_file_url = r.FormValue("upload_music_file_url")
	//upload.From_self = r.FormValue("from_self")

	//output, err := json.Marshal(upload)
	//fmt.Println(string(output))
	if len(upload.Id) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}
	sql := "delete from " + table + " where id=" + upload.Id + ";"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	output := DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))

}

// HTTP POST - /create/follow
func CreateFollowHandler(w http.ResponseWriter, r *http.Request) {
	follow := Follow{}
	r.ParseForm()
	fmt.Println(r)
	follow.User_name = r.FormValue("user_name")
	follow.Follow_user_name = r.FormValue("follow_user_name")
	output, err := json.Marshal(follow)
	fmt.Println(string(output))
	if err != nil || len(follow.User_name) == 0 || len(follow.Follow_user_name) == 0 {
		output = FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}
	sql := "insert into app1_follow set user_name='" + follow.User_name + "',follow_user_name='" + follow.Follow_user_name + "';"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	Response := Base{}
	if err != nil {
		fmt.Println("insert failed!")
		Response.Status = http.StatusRequestTimeout
		output, _ := json.Marshal(Response)
		fmt.Fprintln(w, string(output))
		return
	} else {
		fmt.Println("insert success!")
		Response.Status = http.StatusOK
	}
	//TODO ：三次插入操作必须都成功才算成功 ， 未实现

	//更新 关注数
	update_user_follow_count := "update app1_user set follow_count = follow_count +1 where user_name='" + follow.User_name + "';"
	fmt.Println("sql:", update_user_follow_count)
	_, err = database.Exec(update_user_follow_count)
	output = DatabaseFailReturn(err)

	//更新关注人粉丝数
	upload_follow_user_fans_count := "update app1_user set fans_count = fans_count +1 where user_name='" + follow.Follow_user_name + "';"
	fmt.Println("sql:", upload_follow_user_fans_count)
	_, err = database.Exec(upload_follow_user_fans_count)
	output = DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))
}

// HTTP POST - /delete/follow
func DeleteFollowHandler(w http.ResponseWriter, r *http.Request) {
	follow := Follow{}
	r.ParseForm()
	fmt.Println(r)
	follow.User_name = r.FormValue("user_name")
	follow.Follow_user_name = r.FormValue("follow_user_name")
	output, err := json.Marshal(follow)
	fmt.Println(string(output))
	if err != nil || len(follow.User_name) == 0 || len(follow.Follow_user_name) == 0 {
		output = FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}
	sql := "delete from app1_follow where user_name='" + follow.User_name + "' and follow_user_name='" + follow.Follow_user_name + "';"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	Response := Base{}
	if err != nil {
		fmt.Println("insert failed!")
		Response.Status = http.StatusRequestTimeout
		output, _ := json.Marshal(Response)
		fmt.Fprintln(w, string(output))
		return
	} else {
		fmt.Println("insert success!")
		Response.Status = http.StatusOK
	}
	//TODO ：三次操作必须都成功才算成功 ， 未实现

	//更新 关注数
	update_user_follow_count := "update app1_user set follow_count = follow_count -1 where user_name='" + follow.User_name + "';"
	fmt.Println("sql:", update_user_follow_count)
	_, err = database.Exec(update_user_follow_count)
	output = DatabaseFailReturn(err)

	//更新关注人粉丝数
	//upload_follow_user_fans_count := "update app1_user set fans_count = fans_count -1 where user_name='" + follow.Follow_user_name + "';"
	//fmt.Println("sql:", upload_follow_user_fans_count)
	//_, err = database.Exec(upload_follow_user_fans_count)
	//output = DatabaseFailReturn(err)

	fmt.Fprintln(w, string(output))
}

// HTTP POST - /create/musiclist
func CreateListHandler(w http.ResponseWriter, r *http.Request) {
	music_list := Music_list{}
	r.ParseForm()
	fmt.Println(r)
	music_list.My_list_name = r.FormValue("my_list_name")
	music_list.My_list_user_name = r.FormValue("user_name")

	music_list.Create_date = tm.Format("2006-01-02 03:04:05")
	//music_list.Id = strconv.Itoa(id)

	if len(music_list.My_list_name) == 0 || len(music_list.My_list_user_name) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	sql := "insert into app1_my_list set my_list_user_name='" + music_list.My_list_user_name + "',my_list_name='" + music_list.My_list_name + "',create_date='" + music_list.Create_date + "';"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	Response := Base{}
	if err != nil {
		fmt.Println("insert failed!")
		Response.Status = http.StatusRequestTimeout
		output, _ := json.Marshal(Response)
		fmt.Fprintln(w, string(output))
		return
	} else {
		fmt.Println("insert success!")
		Response.Status = http.StatusOK
	}
	//TODO ：三次插入操作必须都成功才算成功 ， 未实现

	update_user_list_count := "update app1_user set list_count = list_count +1 where user_name='" + music_list.My_list_user_name + "';"
	fmt.Println("sql:", update_user_list_count)
	_, err = database.Exec(update_user_list_count)
	output := DatabaseFailReturn(err)

	fmt.Fprintln(w, string(output))
}

// HTTP POST - /delete/musiclist
func DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	// 通过 id 删除歌单
	music_list := Music_list{}
	r.ParseForm()
	fmt.Println(r)
	music_list.Id = r.FormValue("id")
	//music_list.My_list_name = r.FormValue("my_list_name")
	//music_list.My_list_user_name = r.FormValue("user_name")

	//music_list.Create_date = tm.Format("2006-01-02 03:04:05")
	//music_list.Id = strconv.Itoa(id)

	if len(music_list.Id) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	sql := "delete from app1_my_list where id=" + music_list.Id + ";"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	Response := Base{}
	if err != nil {
		fmt.Println("insert failed!")
		Response.Status = http.StatusRequestTimeout
		output, _ := json.Marshal(Response)
		fmt.Fprintln(w, string(output))
		return
	} else {
		fmt.Println("insert success!")
		Response.Status = http.StatusOK
	}
	//TODO ：二次插入操作必须都成功才算成功 ， 未实现

	//更新 歌单数
	update_user_list_count := "update app1_user set list_count = list_count - 1 where user_name='" + music_list.My_list_user_name + "';"
	fmt.Println("sql:", update_user_list_count)
	_, err = database.Exec(update_user_list_count)
	output := DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))
}

// HTTP POST - /create/collection( 自己上传的歌曲)
func CreateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	collection := Collection{}
	r.ParseForm()
	fmt.Println(r)
	collection.My_list_id = r.FormValue("my_list_id")
	collection.User_name = r.FormValue("user_name")

	if len(collection.My_list_id) == 0 || len(collection.User_name) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	//写入收藏表
	sql := "insert into app1_collection set user_name='" + collection.User_name + "',my_list_id='" + collection.My_list_id + "';"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	output := DatabaseFailReturn(err)
	//Response := Base{}
	//if err != nil {
	//fmt.Println("insert failed!")
	//Response.Status = http.StatusRequestTimeout
	//output, _ := json.Marshal(Response)
	//fmt.Fprintln(w, string(output))
	//return
	//} else {
	//fmt.Println("insert success!")
	//Response.Status = http.StatusOK
	//}

	// 更新歌单中歌曲数
	sql = "update app1_my_list set my_list_count=my_list_count+1 where id=" + collection.My_list_id + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	output = DatabaseFailReturn(err)

	fmt.Fprintln(w, string(output))
}

// HTTP POST - /delete/collection (自己上传的歌曲)
func DeleteCollectionHandler(w http.ResponseWriter, r *http.Request) {
	collection := Collection{}
	r.ParseForm()
	fmt.Println(r)
	collection.My_list_id = r.FormValue("my_list_id")
	collection.User_name = r.FormValue("user_name")

	if len(collection.My_list_id) == 0 || len(collection.User_name) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	sql := "delete from app1_collection where user_name='" + collection.User_name + "' and my_list_id='" + collection.My_list_id + "';"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	Response := Base{}
	if err != nil {
		fmt.Println("insert failed!")
		Response.Status = http.StatusRequestTimeout
		output, _ := json.Marshal(Response)
		fmt.Fprintln(w, string(output))
		return
	} else {
		fmt.Println("insert success!")
		Response.Status = http.StatusOK
	}
	output, _ := json.Marshal(Response)

	// 更新歌单中歌曲数
	sql = "update app1_my_list set my_list_count=my_list_count-1 where id=" + collection.My_list_id + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	output = DatabaseFailReturn(err)

	fmt.Fprintln(w, string(output))
}

// HTTP POST - /create/collection (外链歌曲)
func CreateNotMyselfCollectionHandler(w http.ResponseWriter, r *http.Request) {
	// insert to upload 表和关联表
	table := "app1_upload"
	upload := Upload{}
	listtomusic := ListToMusic{}
	upload.Upload_date = tm.Format("2006-01-02 03:04:05")
	upload.Id = strconv.Itoa(id)
	r.ParseForm()
	//r.ParseMultipartForm(32 << 20)
	upload.Upload_user_name = r.FormValue("upload_user_name")
	upload.Upload_music_name = r.FormValue("upload_music_name")
	upload.Upload_music_file_url = r.FormValue("upload_music_file_url")
	upload.From_self = r.FormValue("from_self")

	listtomusic.My_list_id = r.FormValue("my_list_id")
	listtomusic.Music_id = upload.Id

	if len(upload.Upload_user_name) == 0 || len(upload.Upload_music_name) == 0 || len(upload.Upload_music_file_url) == 0 || len(listtomusic.My_list_id) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	sql := "insert into " + table + " set id=" + upload.Id + ",upload_user_name='" + upload.Upload_user_name + "',upload_music_name='" + upload.Upload_music_name + "',upload_date='" + upload.Upload_date + "',upload_music_file_url='" + upload.Upload_music_file_url + "',from_myself=" + upload.From_self + ";"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	id++
	output := DatabaseFailReturn(err)
	//fmt.Fprintln(w, string(output))

	//写入关联表
	sql = "insert into app1_my_list_to_music set my_list_id=" + listtomusic.My_list_id + ",music_id=" + listtomusic.Music_id + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	output = DatabaseFailReturn(err)

	// 更新歌单中歌曲数
	sql = "update app1_my_list set my_list_count=my_list_count+1 where id=" + listtomusic.My_list_id + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	output = DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))

	//// 通过 歌单id获取 对应的歌单表
	//music_list_table := "app1_my_list"
	//music_list := Music_list{}

	//sql := "select * from " + music_list_table + " where id='" + collection.My_list_id + "';"
	//fmt.Println(sql)
	//rows, err := database.Query(sql)
	//if err != nil {
	//log.Fatal(err)
	//}
	//defer rows.Close()
	//for rows.Next() {
	//rows.Scan(&music_list.Id, &music_list.My_list_name, &music_list.My_list_count, &music_list.My_list_open, &music_list.Create_date, &music_list.My_list_user_name, &music_list.From_self, &music_list.Collection)
	////	Response.MusicListBase = append(Response.MusicListBase, music_list)
	//}
	//err = rows.Err()
	//if err != nil {
	//log.Fatal(err)
	//}

	//// 写入收藏表
	//sql = "insert into app1_collection set user_name='" + collection.User_name + "',my_list_id='" + collection.My_list_id + "';"
	//fmt.Println("sql:", sql)
	//_, err = database.Exec(sql)
	//Response := Base{}
	//if err != nil {
	//fmt.Println("insert failed!")
	//Response.Status = http.StatusRequestTimeout
	//output, _ := json.Marshal(Response)
	//fmt.Fprintln(w, string(output))
	//return
	//} else {
	//fmt.Println("insert success!")
	//Response.Status = http.StatusOK
	//}

	//// 插入歌单

	//if len(music_list.My_list_name) == 0 || len(music_list.My_list_user_name) == 0 {
	//var err error
	//output := FormatErrorReturn(err)
	//fmt.Fprintln(w, string(output))
	//return
	//}

	//sql = "insert into app1_my_list set my_list_user_name='" + music_list.My_list_user_name + "',my_list_name='" + music_list.My_list_name + "',my_list_count='" + music_list.My_list_count + "',collection='" + music_list.Collection + "',create_date='" + music_list.Create_date + "';"
	//fmt.Println("sql:", sql)
	//_, err = database.Exec(sql)
	//Response = Base{}
	//if err != nil {
	//fmt.Println("insert failed!")
	//Response.Status = http.StatusRequestTimeout
	//output, _ := json.Marshal(Response)
	//fmt.Fprintln(w, string(output))
	//return
	//} else {
	//fmt.Println("insert success!")
	//Response.Status = http.StatusOK
	//}
	//output, _ := json.Marshal(Response)
	//fmt.Fprintln(w, string(output))
}

// HTTP POST - /delete/myself/collection (myself create)
func DeleteNotMyselfCollectionHandler(w http.ResponseWriter, r *http.Request) {
	// 先删除歌曲再删除关联表
	// insert to upload 表和关联表
	r.ParseForm()
	upload := Upload{}
	listtomusic := ListToMusic{}
	upload.Upload_date = tm.Format("2006-01-02 03:04:05")
	upload.Id = strconv.Itoa(id)
	//r.ParseMultipartForm(32 << 20)
	//upload.Upload_user_name = r.FormValue("upload_user_name")
	//upload.Upload_music_name = r.FormValue("upload_music_name")
	//upload.Upload_music_file_url = r.FormValue("upload_music_file_url")
	//upload.From_self = r.FormValue("from_self")

	listtomusic.My_list_id = r.FormValue("my_list_id")
	listtomusic.Music_id = r.FormValue("music_id")

	if len(listtomusic.My_list_id) == 0 || len(listtomusic.Music_id) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	// 删除歌曲
	sql := "delete from app1_upload where id=" + listtomusic.Music_id + ";"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)

	output := DatabaseFailReturn(err)

	//删除关联表
	sql = "delete from app1_my_list_to_music where my_list_id=" + listtomusic.My_list_id + " and music_id=" + listtomusic.Music_id + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	output = DatabaseFailReturn(err)

	// 更新歌单中歌曲数
	sql = "update app1_my_list set my_list_count=my_list_count-1 where id=" + listtomusic.My_list_id + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	output = DatabaseFailReturn(err)

	fmt.Fprintln(w, string(output))

	//Response := Base{}
	//if err != nil {
	//fmt.Println("insert failed!")
	//Response.Base.Status = http.StatusRequestTimeout
	//} else {
	//fmt.Println("insert success!")
	//Response.Base.Status = http.StatusOK
	//Response.Upload = upload
	//}
	//output, _ = json.Marshal(Response)
	//fmt.Fprintln(w, string(output))

	//collection := Collection{}
	//r.ParseForm()
	//fmt.Println(r)
	//collection.My_list_id = r.FormValue("my_list_id")
	//collection.User_name = r.FormValue("user_name")

	//if len(collection.My_list_id) == 0 || len(collection.User_name) == 0 {
	//var err error
	//output := FormatErrorReturn(err)
	//fmt.Fprintln(w, string(output))
	//return
	//}

	//sql := "delete from app1_collection where user_name='" + collection.User_name + "' and my_list_id='" + collection.My_list_id + "';"
	//fmt.Println("sql:", sql)
	//_, err := database.Exec(sql)
	//Response := Base{}
	//if err != nil {
	//fmt.Println("insert failed!")
	//Response.Status = http.StatusRequestTimeout
	//output, _ := json.Marshal(Response)
	//fmt.Fprintln(w, string(output))
	//return
	//} else {
	//fmt.Println("insert success!")
	//Response.Status = http.StatusOK
	//}
	//output, _ := json.Marshal(Response)
	//fmt.Fprintln(w, string(output))
}

// HTTP POST - /create/collectionlist  收藏歌单
func CreateCollectionListHandler(w http.ResponseWriter, r *http.Request) {
	// insert to upload 表和关联表
	collection := Collection{}
	r.ParseForm()
	//r.ParseMultipartForm(32 << 20)

	collection.My_list_id = r.FormValue("my_list_id")
	collection.User_name = r.FormValue("user_name")

	if len(collection.My_list_id) == 0 || len(collection.User_name) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	sql := "insert into app1_collection set user_name='" + collection.User_name + "',my_list_id=" + collection.My_list_id + ";"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	id++
	output := DatabaseFailReturn(err)
	fmt.Fprintln(w, string(output))

	////写入关联表
	//sql = "insert into app1_my_list_to_music set my_list_id=" + listtomusic.My_list_id + ",music_id=" + listtomusic.Music_id + ";"
	//fmt.Println("sql:", sql)
	//_, err = database.Exec(sql)
	//output = DatabaseFailReturn(err)

	//// 更新歌单中歌曲数
	//sql = "update app1_my_list set my_list_count=my_list_count+1 where id=" + listtomusic.My_list_id + ";"
	//fmt.Println("sql:", sql)
	//_, err = database.Exec(sql)
	//output = DatabaseFailReturn(err)
	//fmt.Fprintln(w, string(output))
}

// HTTP POST - /create/collectionlist  删除收藏歌单
func DeleteCollectionListHandler(w http.ResponseWriter, r *http.Request) {
	// insert to upload 表和关联表
	collection := Collection{}
	r.ParseForm()
	//r.ParseMultipartForm(32 << 20)

	collection.My_list_id = r.FormValue("my_list_id")
	collection.User_name = r.FormValue("user_name")

	if len(collection.My_list_id) == 0 || len(collection.User_name) == 0 {
		var err error
		output := FormatErrorReturn(err)
		fmt.Fprintln(w, string(output))
		return
	}

	sql := "delete from app1_collection where user_name='" + collection.User_name + "' and my_list_id=" + collection.My_list_id + ";"
	fmt.Println("sql:", sql)
	_, err := database.Exec(sql)
	id++
	output := DatabaseFailReturn(err)
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
	r.HandleFunc("/modifyuser", PutUserHandler).Methods("POST")

	// 添加上传歌曲
	r.HandleFunc("/upload", PostUploadHandler).Methods("POST")

	//删除上传歌曲
	r.HandleFunc("/delete/upload", DeleteUploadHandler).Methods("POST")

	//添加删除关注表
	r.HandleFunc("/create/follow", CreateFollowHandler).Methods("POST")
	r.HandleFunc("/delete/follow", DeleteFollowHandler).Methods("POST")

	//创建歌单
	r.HandleFunc("/create/musiclist", CreateListHandler).Methods("POST")
	// 删除歌单
	r.HandleFunc("/delete/musiclist", DeleteListHandler).Methods("POST")

	//收藏(歌曲not myself)
	r.HandleFunc("/create/collection", CreateCollectionHandler).Methods("POST")
	// 收藏歌曲(not myself)
	r.HandleFunc("/delete/collection", DeleteCollectionHandler).Methods("POST")

	////收藏歌曲( myself)
	r.HandleFunc("/create/notmyself/collection", CreateNotMyselfCollectionHandler).Methods("POST")
	// 收藏歌曲( myself)
	r.HandleFunc("/delete/notmyself/collection", DeleteNotMyselfCollectionHandler).Methods("POST")

	////收藏歌单
	r.HandleFunc("/create/collectionlist", CreateCollectionListHandler).Methods("POST")
	//删除收藏歌单
	r.HandleFunc("/delete/collectionlist", DeleteCollectionListHandler).Methods("POST")
	server := &http.Server{
		Addr: ":8880",
		//Addr:    ":7770",
		Handler: r,
	}
	log.Printf("Listening at %s...", server.Addr)
	server.ListenAndServe()
}
