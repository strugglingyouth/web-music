package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05"))
}
