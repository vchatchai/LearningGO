package webservice 

import (
	"io"
	"fmt"
	"path"
	"storage"
	"strconv"
	"encoding/json"
	"net/http"
)


 


func WebService(){
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}

	http.HandleFunc("/post/", crudHandler)

	server.ListenAndServe();

}

func crudHandler(w http.ResponseWriter, r *http.Request){
	var err error
	switch r.Method {
	case "GET":
		err = handleGET(w,r)
	case "POST":
		err = handlePost(w,r)
 
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
 }

 func handlePost(w http.ResponseWriter,r *http.Request)(err error){
	 
	 
	encoder := json.NewDecoder(r.Body)
	
	var topic storage.Topic

	for{
		err = encoder.Decode(&topic)
		if err == io.EOF {
			break
		}

		if err != nil { 
			panic(err)
			return
		}


	} 
	fmt.Println(topic)
	err =	topic.Create()
	if err != nil {
		panic(err)
		w.WriteHeader(404)
		return
	}
	// topic.Comments[0].Create()
	for _,comment := range topic.Comments {
		err = comment.Create()
		if err != nil {
			panic(err) 
			return
		}	
	}


	w.WriteHeader(200)
	return 

 }

 func handleGET(w http.ResponseWriter,r *http.Request)(err error){
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return 
	}
	topic, err  := storage.GetTopic(id)
	if err != nil {
		panic(err)
		return
	} 
	json, err := json.Marshal(&topic)
	if err  != nil { 
		panic(err)
		return 
	}

	fmt.Println(string(json))
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	fmt.Println("handleGET done.")
	return 



 }