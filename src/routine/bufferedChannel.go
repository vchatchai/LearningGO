package routine

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func BufferedChannel() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}
	results := make(chan HomePageSize)
	for _, url := range urls {
		go func(url string) {
			fmt.Println("Get", url)

			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for _, url := range urls {
		fmt.Println(url)
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL)
}
