package main

import "testing"

func TestGetPhotos(t *testing.T) {
	slice := []Photo{}
	err := GetPhotos(slice)

	if err != nil {
		t.Error("Error Test Get Photos")
	}

	if len(slice) <= 0 {
		t.Error("expect to get a photo but it empty")
	}
}
