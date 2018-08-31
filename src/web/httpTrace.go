package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"os"
)

func HttpTrace(url string) {

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done", network, addr)
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("Requessting data from Server!")
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, response.Body)
}
