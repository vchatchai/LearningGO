package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func WebClient(url string) {
	data, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
