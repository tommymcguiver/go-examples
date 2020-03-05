package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Greet", time.Now())
	defer fmt.Println("Greet Finished", time.Now())

	results, err, _ := group.Do("greet", func() (interface{}, error) {
		time.Sleep(3 * time.Second)
		t := time.Now()
		fmt.Println("Simulate results", t)
		return t, nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello World! %s", results)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
