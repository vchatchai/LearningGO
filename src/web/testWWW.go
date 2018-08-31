package web

import (
	"fmt"
	"net/http"
)

func CheckStatusOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `Fine!`)
}

func StatusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func MyHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func TestWWW() {
	PORT := ":8001"

	http.HandleFunc("/CheckStatusOK", CheckStatusOK)
	http.HandleFunc("/StatusNotFound", StatusNotFound)
	http.HandleFunc("/", MyHandler2)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
