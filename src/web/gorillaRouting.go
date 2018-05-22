package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func GorillaRoute() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)

}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// pageId := vars["id"]
	// fileName := "files/" + pageId + ".html"
	// http.ServeFile(w, r, fileName)
	response := "The time is now " + time.Now().String() + " :" + vars["id"]
	fmt.Fprint(w, response)
}
