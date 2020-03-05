package main

import (
	"net/http"
	"testing"
)

func Test_greetTimer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			greetTimer(tt.args.w, tt.args.r)
		})
	}

	t.Cleanup(func() {
		t.Log("1. cleanup")
	})
	t.Cleanup(func() {
		t.Log("2. cleanup")
	})

	t.Log("completed greet test")
}

func Test_main(t *testing.T) {

	t.Cleanup(func() {
		t.Log("1. cleanup")
	})
	t.Cleanup(func() {
		t.Log("2. cleanup")
	})

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
	t.Log("Completed test main")
}
