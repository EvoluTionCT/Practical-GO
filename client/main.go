package main

import (
	"encoding/json"
	"fmt"
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

func GetPhotos(p *[]Photo) error {
	resp, err := http.Get(URL + "/photos")
	if err != nil {
		// handle error here.
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(p)
}

func main() {
	var p []Photo
	err := GetPhotos(&p)
	if err != nil {
		//handle here
		log.Fatal(err)
	}

	fmt.Println(p)
}
