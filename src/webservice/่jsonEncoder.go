package webservice

import (
	"fmt"
	"encoding/json"
	"os"
)

func JsonEcoder(){

	post := Post{
		Id: "1",
		Content: "Hello World!",
		Author: Author{
		Id: "2",
		Name: "Sau Sheong",
		},
		Comments: []Comment{
		Comment{
		Id: "3",
		Content: "Have a great day!",
		Author: Author{
			Id:"1",
			Name:"Adam",

		}, 
		},
		Comment{
		Id: "4",
		Content: "How are you today?",
		Author: Author{
			Id: "2",
			Name:"Betty",

		} ,
		},
		},
		}

	jsonFile, err  := os.Create("jsonEncode.json")
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
		
		err = encoder.Encode(&post)
		if err != nil {
			panic(err)
			return
		}


	fmt.Println("done")

}