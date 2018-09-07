package web

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func AdvanceWebClient(urlname string) {
	URL, err := url.Parse(urlname)
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}

	c := http.Client{
		Timeout: 15 * time.Second,
	}

	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		fmt.Println("GET:", err)
		return
	}

	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in DO():", err)
		return
	}
	fmt.Println("Status Code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}
	if httpData.ContentLength == -1 {
		fmt.Println("Contentlength is unknow!")
	} else {
		fmt.Println("Contentlength ", httpData.ContentLength)
	}

	length := 0
	var buffer [1000]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Caculated response data length:", length)
}
