package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/photos",func(w http.ResponseWriter,r *http.Request){
		pht := []byte(`{"title":"Hello, It's me"}`)
		w.WriteHeader(http.StatusOK)
		w.Write(pht)
	})
	fmt.Println("Starting...")
	
	log.Fatal(http.ListenAndServe(":1234",nil))
}
