package web

import (
	"fmt"
	"net/http"
	"time"
)

var Addr string = "127.0.0.1:8011"
var port rune = 10

func setCookies(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:  "second_cookie",
		Value: "Manning Publications Co",
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())

}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Println(r.Header)
	fmt.Fprintln(w, h)
}

func Cookies() {
	t := time.Now()
	fmt.Println(t)
	server := http.Server{
		Addr: Addr,
	}

	http.HandleFunc("/set_cookies", setCookies)
	http.HandleFunc("/get_cookies", getCookie)

	server.ListenAndServe()
}
