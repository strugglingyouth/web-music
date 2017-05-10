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
	new_note := Note{}
	r.ParseForm()
	fmt.Println(r)
	new_note.Id = r.FormValue("id")
	new_note.Title = r.FormValue("title")
	new_note.Description = r.FormValue("description")
	output, err := json.Marshal(new_note)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Errorn\n:", err)
	}
	sql := "update notes set title='" + new_note.Title + "',description='" + new_note.Description + "' where id=" + new_note.Id + ";"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	if err != nil {
		fmt.Println("Error:\n", err)
	} else {
		fmt.Println("exec insert sql success!")
	}
	//解码
	//err = json.NewDecoder(r.Body).Decode(&noteToUpd)
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
	r.HandleFunc("/modifyuser/{id}", PutUserHandler).Methods("PUT")

	server := &http.Server{
		Addr:    ":8880",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
