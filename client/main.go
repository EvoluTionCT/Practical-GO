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

type DoGet struct {
}

type getTypicode struct {
	client DoGet
}

func (d DoGet) Do (url string) (*http.Response, error) {
	return http.Get(url + "/photos")
}

func (t getTypicode) GetPhotos(p *[]Photo) error {
	resp, err := t.client.Do(URL)
	if err != nil {
		// handle error here.
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(p)
}

func main() {
	var p []Photo
	t := getTypicode{
		client: DoGet{},
	}
	err := t.GetPhotos(&p)
	if err != nil {
		//handle here
		log.Fatal(err)
	}

	fmt.Println(p)
}
