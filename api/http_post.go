package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"strings"
	"github.com/widuu/gojson"
)

func httpPost() {
	resp, err := http.Get("http://songsearch.kugou.com/song_search_v2?keyword=故乡&page=1&pagesize=1")
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
	lists := gojson.Json(content).Get("data").Get("lists")
	fmt.Println(len(lists))

	for i := range lists {
		fmt.Println(i)
	}

	//map1 := gojson.Json(content).Get("data").Get("lists").ToArray()

	//// map to json
	//for i := range map1 {
	//str, err := json.Marshal(map1)

	//if err != nil {
	//fmt.Println(err)
	//}
	//fmt.Println("map to json", string(str))
	//}

	//fmt.Println(gojson.Json(list).Get("SongName"))
}

func main() {
	httpPost()
}
