package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World")
		}))
	if err != nil {
		fmt.Printf("failed to terminate server: %v\n", err)
		os.Exit(1)
	}
}
