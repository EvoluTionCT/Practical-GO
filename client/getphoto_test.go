package main

import (
	"net/http"
	"testing"
)

type mockDoGet struct {

}

func (m mockDoGet) Do(url string) (*http.Response, error) {
	return nil,nil
}

func TestGetPhotos(t *testing.T) {
	slice := []Photo{}
	tc := getTypicode{
		client:mockDoGet{},
	}
	err := tc.GetPhotos(&slice)

	if err != nil {
		t.Error("Error Test Get Photos")
	}

	if len(slice) <= 0 {
		t.Error("expect to get a photo but it empty")
	}
}
