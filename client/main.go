package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Album struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

const (
	URL = "https://jsonplaceholder.typicode.com"
)

type Clienter interface {


}

func main() {
	resp, err := http.Get(URL + "/photos")
	if err != nil {
		// handle error here.
		log.Fatal("get user error", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read all err:", err)
	}
	resp.Body.Close()
	var p []Photo
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	resp2, err := http.Get(URL + "/albums")
	if err != nil {
		log.Fatal("get user error", err)
	}
	b2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatal("read all err:", err)
	}
	resp2.Body.Close()
	var a []Album
	err = json.Unmarshal(b2, &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}


