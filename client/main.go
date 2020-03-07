package client

import (
	"fmt"
	"log"

	"github.com/EvoluTionCT/Practical-GO/typicode"
)

func main() {
	var p []typicode.Photo
	t := typicode.New("/photos")
	err := t.Get(&p)

	tc := typicode.New("/albums")
	var a []typicode.Album
	err = tc.Get(&a)

	if err != nil {
		//handle here
		log.Fatal(err)
	}

	fmt.Println(p)
}
