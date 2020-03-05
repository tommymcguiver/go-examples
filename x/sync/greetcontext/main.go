package main

import (
	"context"
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

	ctx, _ := context.WithTimeout(r.Context(), time.Second)

	ch := gt.DoChan("greetTimer", func() (interface{}, error) {
		time.Sleep(3 * time.Second)
		t := time.Now()
		fmt.Println("Simulate results", t)
		return t, nil
	})

	select {
	case <-ctx.Done():
		if err := ctx.Err(); ctx.Err() != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case results = <-ch:
	}

	fmt.Fprintf(w, "Hello World! %s", results.Val)
}

func main() {
	http.HandleFunc("/greetTimer", greetTimer)
	http.ListenAndServe(":8080", nil)
}
