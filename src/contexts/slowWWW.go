package contexts

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func random(min, max int) int {

	return rand.Intn(max-min) + min
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	delay := random(0, 15)
	time.Sleep(time.Duration(delay) * time.Second)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Delay: %d\n", delay)

}

func SlowWWWW() {
	fmt.Println("Start  slowWWW")
	seed := time.Now().Unix()
	rand.Seed(seed)
	PORT := ":9000"
	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}
