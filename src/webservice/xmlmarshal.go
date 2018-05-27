package webservice

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func Marshal() {

	post := Post{
		Id:      "1",
		Content: "Hello World",
		Author: Author{
			Id:   "2",
			Name: "Goher",
		},
	}

	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(xml.Header)
	fmt.Println(string(output))

	err = ioutil.WriteFile("posted_create.xml", output, 0644)
	if err != nil {
		panic(err)
	}

}
