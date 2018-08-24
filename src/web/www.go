package web

import (
	"fmt"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "time current is :"
	fmt.Fprintf(w, `<h1 align="center">%s</h1>`, Body)
	fmt.Fprintf(w, `<h2 align="center">%s</h2>`, t)
	fmt.Fprintf(w, `Serving: %s`, r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func WWW() {
	server := http.Server{
		Addr: ":8001",
	}
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
