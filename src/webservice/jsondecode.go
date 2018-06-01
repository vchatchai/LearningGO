package webservice

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func JsonDecode() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	post, err := JsonDecoder(jsonFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(post)

}

func JsonDecoder(jsonFile io.Reader) (post Post, err error) {

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)

	return
}
