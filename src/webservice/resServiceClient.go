package webservice

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RestService() {
	response, err := http.Get("http://api.theysaidso.com/qod")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(contents))

	// decoder := xml.NewDecoder(response.Body)

}
