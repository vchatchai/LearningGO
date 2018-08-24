package profiling

import (
	"fmt"
	"net/http"
	"net/http/pprof"
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

	r := http.NewServeMux()

	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/", myHandler)
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(server.Addr, r)
	if err != nil {
		panic(err)
	}
}
