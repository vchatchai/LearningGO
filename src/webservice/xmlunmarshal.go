package webservice

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct { //#A
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr",json:"id"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml,"comments->comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func Unmarshal() {
	fmt.Println("Start XML")
	xmlFile, err := os.Open("post2.xml")
	if err != nil {
		fmt.Println("Err opening XML file:", err)
		return
	}

	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(string(xmlData))
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)

	xml, err := xml.Marshal(post)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(xml))

}
