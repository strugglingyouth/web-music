package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

//对字符串进行MD5哈希
func Md5hash(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	//fmt.Println(data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func main() {

	var data string = "abc"
	fmt.Printf("MD5 : %s\n", Md5hash(data))

}
