package web

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func HttpHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called -", name)
		h(w, r)
	}
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protect")
		h.ServeHTTP(w, r)
	})
}

func ListenAndServeTLS() {
	handler := MyHandler{}
	// hello := HelloHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8011",
		// Handler: &handler,
	}

	http.HandleFunc("/hello", log(HttpHelloHandler))
	http.Handle("/world", protect(&handler))

	server.ListenAndServeTLS("cert.pem", "key.pem")
	// server.ListenAndServe()

}
