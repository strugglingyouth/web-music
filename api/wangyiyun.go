package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"strings"
	"github.com/widuu/gojson"
	"strings"
)

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

func httpPost() {
	resp, err := http.Post("http://music.163.com/api/search/get/",
		"application/x-www-form-urlencoded",
		strings.NewReader("s=玫瑰色的你&limit=20&type=1&offset=0"))
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
	fmt.Println(content)

	// parse json
	music_id = gojson.Json(content).Get("result").Get("songs").Getkey("id", 1).Tostring()
	fmt.Println(gojson.Json(content).Get("result").Get("songs").Getkey("id", 1).Tostring())

}

func main() {
	httpPost()
}
