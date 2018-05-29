package webservice

import (
	"encoding/json" 
	"fmt"
)

func JsonMarshal(){
	fmt.Println("JsonMarshal")
	// jsonFile, err := os.Create("jsonMarshal.json")
	post  := Post {
		Id: "1",
		Content: "Hello World",
		Author: Author{
			Id: "2",
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
				Id: "3",
				Content: "Have a grate day!",
				Author: Author {
					Id:"1",
					Name:"Adam",

				} ,
			},
			Comment{
				Id: "4",
				Content: "How are you today?",
				Author: Author {
					Id:"2",
					Name:"Betty",
				},
			},
		},
	}
	output,err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))


}