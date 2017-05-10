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

type Notes struct {
	Notes  []Note `json:"notes"`
	Status int
}
type Note struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// HTTP POST - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	new_note := Note{}
	//解码到note, 使用此方式时接受到的类型为 string
	//err := json.NewDecoder(r.Body).Decode(&new_note)
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
	sql := "insert into notes set id=" + new_note.Id + ",title='" + new_note.Title + "', description='" + new_note.Description + "';"
	fmt.Println("sql:", sql)
	_, err = database.Exec(sql)
	if err != nil {
		fmt.Println("Error:\n", err)
	} else {
		fmt.Println("exec insert sql success!")
	}
}

// HTTP GET - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r)
	// 不缓存数据
	w.Header().Set("Pragma", "no-cache")
	rows, err := database.Query("select * from notes limit 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Response := Notes{}
	for rows.Next() {
		note := Note{}
		rows.Scan(&note.Id, &note.Title, &note.Description)
		Response.Notes = append(Response.Notes, note)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// 添加 status
	Response.Status = http.StatusOK
	output, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}

///HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
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
	db, err := sql.Open("mysql", "root@/web")
	if err != nil {
		log.Fatal("Error on initializing database connection: %s", err.Error())
	}
	defer db.Close()
	database = db
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
