package webservice

import (
	"io"
	"encoding/json"
	"fmt"
	"os"
)


func JsonDecode() {
	jsonFile, err := os.Open("post.json")
	if err != nil { 
		fmt.Println("Error opening JSON file:", err)
		return 
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	count := 1
	for {
		var post Post
		err := decoder.Decode(&post)

		if err == io.EOF {
			break
		}

		if err != nil { 
			panic(err)
			return
		}
		fmt.Println("<",count, ">",post)
		count++ 
	}
}