package webservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func JsonMarshal() {

	file, err := os.Open("post.json")

	if err != nil {
		panic(nil)
	}
	defer file.Close()

	jsonFile, err := ioutil.ReadAll(file)

	post := Post{}

	err = json.Unmarshal(jsonFile, &post)

	if err != nil {
		panic(err)
	}

	fmt.Println(post)

}
