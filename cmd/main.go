package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	data, err := http.Get("https://drive.google.com/file/d/1IwZ3uUCHGpSL2OoQu4mtbw7Ew3ZamcGB/view")
	if err != nil {
		log.Println(err)
	}
	defer data.Body.Close()
	fmt.Println(data)
}
