package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

var gt singleflight.Group

func greetTimer(w http.ResponseWriter, r *http.Request) {
	var results singleflight.Result
	fmt.Println("Greet", time.Now())
	defer fmt.Println("Greet finished", time.Now())

	ch := gt.DoChan("greetTimer", func() (interface{}, error) {
		time.Sleep(3 * time.Second)
		t := time.Now()
		fmt.Println("Simulate results", t)
		return t, nil
	})

	timeout := time.After(1 * time.Second)

	select {
	case <-timeout:
		fmt.Println("Timeout", time.Now())
		http.Error(w, "timeout", http.StatusInternalServerError)
		return
	case results = <-ch:
	}

	if results.Err != nil {
		http.Error(w, results.Err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello World! %s", results.Val)
}

func main() {
	http.HandleFunc("/greetTimer", greetTimer)
	http.ListenAndServe(":8080", nil)
}
