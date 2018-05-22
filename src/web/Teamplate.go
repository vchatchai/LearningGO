package web

import (
	"html/template"
	"net/http"
)

func processTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	// daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, "hello")
}

func processTemplate2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}

func Template() {
	server := http.Server{
		Addr: "127.0.0.1:8088",
	}
	http.HandleFunc("/process", processTemplate)
	http.HandleFunc("/process2", processTemplate2)
	server.ListenAndServe()
}
