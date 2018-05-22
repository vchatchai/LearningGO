package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Uploading() {
	server := http.Server{
		Addr: "127.0.0.1:8011",
	}

	http.HandleFunc("/process", process)
	http.Handle("/", http.FileServer(http.Dir(".")))

	server.ListenAndServe()

}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)

	fmt.Println(r.MultipartForm)

	fileHader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}
