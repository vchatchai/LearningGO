package  webservice 

import (
	"encoding/xml"
	"os"
)

func XMLEncode(){
	post := Post {
		Id: "1",
		Content: "Ni Hao",
		Author: Author {
			Id: "2",
			Name: "Mao",
		},
		Comments: []Comment{
			Comment{
				Id:"01",
				Content: "Comment",
				Author: Author{
					Id: "3",
					 Name: "Yang",
				},
			},
			
			Comment{
				Id:"02",
				Content: "Discuss",
				Author: Author{
					Id: "3",
					 Name: "Ming",
				},
			},
		},
	}


	xmlFile, err := os.Create("postEncoder.xml")
	if err != nil {
		panic(err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		panic(err)
	}


}