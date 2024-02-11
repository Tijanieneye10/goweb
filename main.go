package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}

}
