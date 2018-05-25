package webservice

import (
	"io/ioutil"
	"os"
	"encoding/xml"
	"fmt"
)

type Post struct {  //#A
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author  `xml:"author"`
	Xml string `xml:",innerxml"`
}

type Author struct {
	Id string  `xml:"id,attr"`
	Name string `xml:",chardata"`

}

func XML() {
	fmt.Println("Start XML")
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Err opening XML file:", err)
		return
	}

	defer  xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(xmlData)

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)

	xml, err := xml.Marshal(post)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(xml[:]))


}
