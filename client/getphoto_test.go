package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockDoGet struct {
	Doer
}

func (m mockDoGet) Do(url string) (*http.Response, error) {
	return &http.Response{
		Body:             ioutil.NopCloser(bytes.NewReader([]byte(
			`[{
			"title": "Hello"
		  	}]`))),
	},nil
}

func TestGetPhotos(t *testing.T) {
	p := []Photo{}
	tc := getTypicode{
		client:mockDoGet{},
	}
	tc.GetPhotos(&p)

	if p[0].Title != "Hello" {
		t.Error("expect to get a photo but it empty")
	}
}
