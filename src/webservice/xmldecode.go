package webservice

import (
	"fmt"
	"encoding/xml"
	"io"
	"os"
)

func XMLDecode() {
	xmlFile, err := os.Open("post2.xml")

	if err != nil {
		panic(err)
		return
	}

	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		// t, err := decoder.Token()
		// if err == io.EOF {
		// 	break
		// }
		// if err != nil {
		// 	panic(err)
		// 	return
		// }

		// switch se := t.(type) {
		// case xml.StartElement:
		// 	if se.Name.Local == "comment" {
		// 		var comment Comment
		// 		decoder.DecodeElement(&comment, &se)
		// 	}
		// }
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}

		if err != nil { 
			panic(err)
			return
		}

		fmt.Println(post)
		
	}
	fmt.Println("done")
}
