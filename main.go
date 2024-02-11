package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)

	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "Hello world")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fprintf)
}
