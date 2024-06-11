package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Too Tasty of a Trade!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the website player")
	})

	http.ListenAndServe(":3333", nil)
}
