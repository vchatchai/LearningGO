package network

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func AnotherTimeOut(URL string, second int) {
	var timeout = time.Duration(time.Second * time.Duration(second))
	client := http.Client{
		Timeout: timeout,
	}
	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
