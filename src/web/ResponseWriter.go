package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>
	`
	w.Write([]byte(str))
}

func wirteHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Chatchai Vichai",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)

	w.Write(json)
}

func ResponseWriter() {
	server := http.Server{
		Addr: "127.0.0.1:8011",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", wirteHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()

}
