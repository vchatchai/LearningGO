package web

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Httprouter() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.GET("/body", body)

	server := http.Server{
		Addr:    "127.0.0.1:8011",
		Handler: mux,
	}
	server.ListenAndServe()

}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o := r.Header.Get("Accept-Encoding")
	for k, v := range o {
		fmt.Println(k, v)
	}

	for key, value := range r.Header {
		fmt.Println(key, value)

		for o, v := range value {
			fmt.Println(o, v)
		}

	}

	fmt.Fprintln(w, r.Header)
}

func body(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintf(w, string(body))
}
